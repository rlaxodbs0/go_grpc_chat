package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"os"
	_ "os"
	"time"

	pb "go_grpc_chat/pb"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:4317"
)

var cmdString string

type Client struct {
	username string
	password string
	pb.ChatTaskClient
}

func (cl *Client) session() {
	cmdStringMap := map[string]interface{} {
		"login": cl.login,
		"logout": cl.logout,
		"signup": cl.signup,
	}
	for {
		fmt.Printf("GRPCHAT >>> ")
		fmt.Scanln(&cmdString)
		v, exists := cmdStringMap[cmdString]
		if !exists {
			log.Printf("%v is invalid command", cmdString)
			break
		}
		v.(func())()
	}
}

func (cl *Client)signup() {
	var username, password string
	fmt.Printf("username: ")
	fmt.Scanln(&username)
	fmt.Printf("password: ")
	fmt.Scanln(&password)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := cl.ChatTaskClient.SignUp(ctx, &pb.UserInfo{Username: username, Password: password})
	if err != nil {
		log.Fatalf("%v occured", err)
	}
	if r.Response == pb.ResponseType_ALREADYEXISTS {
		log.Printf("%s already exists\n", username)
	}
	log.Printf("Thank you for signup! %s\n", username)
}


func (cl *Client) login()  {
	var username, password string
	fmt.Printf("username: ")
	fmt.Scanln(&username)
	fmt.Printf("password: ")
	fmt.Scanln(&password)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := cl.ChatTaskClient.Login(ctx, &pb.UserInfo{Username: username, Password: password})
	if err != nil {
		log.Fatalf("%v occured", err)
	}
	if r.Response == pb.ResponseType_NOMATCH {
		log.Printf("There is no match username %s and password\n", username)
	} else if r.Response == pb.ResponseType_SUCCESS {
		log.Printf("Welcome! %s\n", username)
		cl.username = username
		cl.password = password
	}
	cl.chatSession(ctx)
}

func (cl *Client) logout() {
	var username, password string
	fmt.Printf("username: ")
	fmt.Scanln(&username)
	fmt.Printf("password: ")
	fmt.Scanln(&password)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	r, err := cl.ChatTaskClient.Logout(ctx, &pb.UserInfo{Username: username, Password: password})
	if err != nil {
		log.Fatalf("%v occured", err)
	}
	if r.Response == pb.ResponseType_SUCCESS {
		log.Printf("Goodbye! %s\n", username)
	}
}

func (cl *Client)chatSession(ctx context.Context) error{
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	stream, err := cl.ChatTaskClient.Chat(ctx)
	if err != nil {
		log.Fatalf("%v occured", err)
	}
	defer stream.CloseSend()
	go cl.send(stream)
	return cl.receive(stream)
}

func (cl *Client)send (stream pb.ChatTask_ChatClient) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)
	stream.Send(&pb.Message{Username:cl.username})
	for {
		fmt.Printf("CHAT >>>")
		select {
		case <- stream.Context().Done():
			fmt.Println("client send loop disconnected")
		default:
			if sc.Scan() {
				if err := stream.Send(&pb.Message{Username: cl.username,
					Event: &pb.Message_BroadcastMessage_{
						BroadcastMessage: &pb.Message_BroadcastMessage{Message:sc.Text()}}}); err != nil {
					return
				}
			}
		}
	}
}


func (cl *Client) receive(stream pb.ChatTask_ChatClient) error {
	for {
		res, err := stream.Recv()
		if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
			return nil
		} else if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		switch evt := res.Event.(type) {
		case *pb.Message_BroadcastMessage_:
			log.Printf("BROADCAST %v %v", res.Username, evt.BroadcastMessage.Message)
		default:
			log.Println("hi")
		}
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChatTaskClient(conn)
	client := Client{ChatTaskClient: c}
	client.session()
}