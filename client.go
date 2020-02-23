
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
cmdStringMap := map[string]interface{} {
	"login": login,
	"logout": logout,
	"signup": signup,
}

func session(c pb.ChatTaskClient) {
	for {
		fmt.scanf("GRPCHAT >>>  %s")
		if cmdString == "login" {
			fmt.Print("UserName: ")
			fmt.Scanln(&cmdString)
			log.Printf("Greeting: %s", r.Response)
		} else {
			fmt.Println("error occur! bye")
			break
		}
	}
}

func signup(c pb.ChatTaskClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var username, password string
	fmt.Scanf("username: %s\n", &username)
	fmt.Scanf("password: %s\n", &password)
	r, err := c.Signup(ctx, &pb.UserInfo{UserName: username, Password: password})
	if err != nil {
		log.Fatalf("%v occured", err)
	}
	if r.Response == pb.ResponseType_ALREADYEXISTS {
		log.Printf("%s already exists\n", username)
	}
	log.Printf("Thank you for signup! %s\n", username)
	return err
}


func login(c pb.ChatTaskClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var username, password string
	fmt.Scanf("username: %s\n", &username)
	fmt.Scanf("password: %s\n", &password)
	r, err := c.Login(ctx, &pb.UserInfo{UserName: username, Password: password})
	if err != nil {
		log.Fatalf("%v occured", err)
	}
	if r.Response == pb.ResponseType_NOMATCH {
		log.Printf("There is no match username %s and password\n", username)
	}
	log.Printf("Welcome! %s\n", username)
	return err
}

func logout(c pb.ChatTaskClient) error{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var username, password string
	fmt.Scanf("username: %s\n", &username)
	fmt.Scanf("password: %s\n", &password)
	_, err := c.Login(ctx, &pb.UserInfo{UserName: username, Password: password})
	if err != nil {
		log.Fatalf("%v occured", err)
	}
	log.Printf("Goodbye! %s\n", username)
	return err
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