
package main

import (
	"context"
	"fmt"
	"log"
	_ "os"
	"time"

	pb "go_grpc_chat/pb"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
)

var cmdString string

func session(c pb.ChatTaskClient) {
	cmdStringMap := map[string]interface{} {
		"login": login,
		"logout": logout,
		"signup": signup,
		"search": search,
		"invite": invite,
	}
	for {
		fmt.Printf("GRPCHAT >>> ")
		fmt.Scanln(&cmdString)
		v, exists := cmdStringMap[cmdString]
		if !exists {
			log.Printf("%v is invalid command", cmdString)
			break
		}
		v.(func(pb.ChatTaskClient))(c)
	}
}

func signup(c pb.ChatTaskClient) {
	var username, password string
	fmt.Printf("username: ")
	fmt.Scanln(&username)
	fmt.Printf("password: ")
	fmt.Scanln(&password)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Signup(ctx, &pb.UserInfo{UserName: username, Password: password})
	if err != nil {
		log.Fatalf("%v occured", err)
	}
	if r.Response == pb.ResponseType_ALREADYEXISTS {
		log.Printf("%s already exists\n", username)
	}
	log.Printf("Thank you for signup! %s\n", username)
}


func login(c pb.ChatTaskClient)  {
	var username, password string
	fmt.Printf("username: ")
	fmt.Scanln(&username)
	fmt.Printf("password: ")
	fmt.Scanln(&password)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Login(ctx, &pb.UserInfo{UserName: username, Password: password})
	if err != nil {
		log.Fatalf("%v occured", err)
	}
	if r.Response == pb.ResponseType_NOMATCH {
		log.Printf("There is no match username %s and password\n", username)
	} else if r.Response == pb.ResponseType_SUCCESS {
		log.Printf("Welcome! %s\n", username)
		go getInviteNotify(c, username)
	}
}

func logout(c pb.ChatTaskClient) {
	var username, password string
	fmt.Printf("username: ")
	fmt.Scanln(&username)
	fmt.Printf("password: ")
	fmt.Scanln(&password)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Logout(ctx, &pb.UserInfo{UserName: username, Password: password})
	if err != nil {
		log.Fatalf("%v occured", err)
	}
	if r.Response == pb.ResponseType_SUCCESS {
		log.Printf("Goodbye! %s\n", username)
	}
}

func search(c pb.ChatTaskClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Search(ctx, &pb.UserInfo{UserName: "temp", Password: "temp"})
	if err != nil {
		log.Fatalf("%v occured", err)
	}
	log.Printf("Goodbye! %s\n", r.UserNameActiveMap)

}

func invite(c pb.ChatTaskClient) {
	var username, password string
	fmt.Printf("username: ")
	fmt.Scanln(&username)
	fmt.Printf("invite username: ")
	fmt.Scanln(&password)
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	_, err := c.Invite(ctx, &pb.InviteInfo{Sender: username, Receiver: password})
	if err != nil {
		log.Fatalf("%v occured", err)
	}

}

func getInviteNotify(c pb.ChatTaskClient, username string){
	ctx, cancel := context.WithTimeout(context.Background(), 1000 * time.Second)
	defer cancel()
	stream, err := c.GetInviteNotify(ctx, &pb.UserInfo{UserName: username})
	if err != nil {
		log.Fatalf("%v occured", err)
	}
	for {
		resp, _	 := stream.Recv()
		log.Print(resp.UserName, "invites you y/n")
	}

}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChatTaskClient(conn)
	session(c)
}