
package main

import (
	"context"
	"fmt"
	"log"
	_"os"
	"time"

	"google.golang.org/grpc"
	pb "go_grpc_chat/pb"
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
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	var username, password string
	fmt.Printf("username: ")
	fmt.Scanln(&username)
	fmt.Printf("password: ")
	fmt.Scanln(&password)
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var username, password string
	fmt.Printf("username: ")
	fmt.Scanln(&username)
	fmt.Printf("password: ")
	fmt.Scanln(&password)
	r, err := c.Login(ctx, &pb.UserInfo{UserName: username, Password: password})
	if err != nil {
		log.Fatalf("%v occured", err)
	}
	if r.Response == pb.ResponseType_NOMATCH {
		log.Printf("There is no match username %s and password\n", username)
	}
	log.Printf("Welcome! %s\n", username)
}

func logout(c pb.ChatTaskClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var username, password string
	fmt.Printf("username: ")
	fmt.Scanln(&username)
	fmt.Printf("password: ")
	fmt.Scanln(&password)
	_, err := c.Login(ctx, &pb.UserInfo{UserName: username, Password: password})
	if err != nil {
		log.Fatalf("%v occured", err)
	}
	log.Printf("Goodbye! %s\n", username)
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