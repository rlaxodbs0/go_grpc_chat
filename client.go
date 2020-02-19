/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
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
	defaultName = "world"
)


func session(c pb.ChatTaskClient) {
	var cmdString string
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for {
		fmt.Print("GRPCHAT >> ")
		fmt.Scanln(&cmdString)
		if cmdString == "exit" {
			break
		}
		if cmdString == "login" {
			fmt.Print("UserName: ")
			fmt.Scanln(&cmdString)
			r, err := c.Login(ctx, &pb.UserInfo{UserName: cmdString, Password: "ppap"})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			log.Printf("Greeting: %s", r.Response)
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
	session(c)

}