package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "go_grpc_chat/pb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)


type server struct {
	pb.UnimplementedChatTaskServer
	UserInfo map[string]string
	UserStatus map[string]string
	UserInviteSession map[string]pb.ChatTask_GetInviteNotifyServer
	UserChatSession map[string]pb.ChatTask_ChatMessageServer
}

func (s *server) SignUp(ctx context.Context, info *pb.UserInfo) (*pb.SignupResponse, error){
	_, exists := s.UserInfo[info.UserName]
	if exists {
		return &pb.SignupResponse{Response: pb.ResponseType_ALREADYEXISTS}, nil
	}
	s.UserInfo[info.UserName] = info.Password
	log.Printf("User Signup: %v",info.UserName)
	return &pb.SignupResponse{Response:pb.ResponseType_SUCCESS}, nil
}

func (s *server) Login(ctx context.Context, info *pb.UserInfo) (*pb.LoginResponse, error) {
	_, exists := s.UserInfo[info.UserName]
	if !exists {
		return &pb.LoginResponse{Response: pb.ResponseType_NOMATCH}, nil
	}
	log.Printf("User Login: %v",info.UserName)
	s.UserInfo[info.UserName] = "Active"
	return &pb.LoginResponse{Response: pb.ResponseType_SUCCESS}, nil
}

func (s *server) Logout(ctx context.Context, info *pb.UserInfo) (*pb.LogoutResponse, error) {
	log.Printf("User Logout: %v",info.UserName)
	userStatus[info.UserName] = "Inactive"
	return &pb.LogoutResponse{Response: pb.ResponseType_SUCCESS}, nil
}

func (s * server) Search(ctx context.Context, info *pb.UserInfo) (*pb.UserList, error) {
	log.Printf("User Search: %v",info.UserName)
	return &pb.UserList{UserNameActiveMap: userStatus}, nil
}

func (s * server) GetInviteNotify(info *pb.UserInfo, stream pb.ChatTask_GetInviteNotifyServer) error {
	userInviteSession[info.UserName] = stream
	fmt.Print(userInviteSession)
	for{}
}

func (s * server) Invite(ctx context.Context, info *pb.InviteInfo) (*pb.InviteResponse, error) {
	log.Printf("%v invites %v", info.Sender, info.Receiver)
	fmt.Print(userInviteSession)
	userInviteSession[info.Receiver].Send(&pb.UserInfo{UserName:info.Sender})
	return &pb.InviteResponse{Response: pb.ResponseType_SUCCESS}, nil
}

func (s * server) ChatMessage(stream pb.ChatTask_ChatMessageServer) error {
	for {
		msg, _ := stream.Recv()
		_, exists := userChatSession[msg.Receiver]
		if !exists{
			userChatSession[msg.Receiver] = stream
		}
		userChatSession[msg.Receiver].Send(&pb.Message{Text: msg.Text})
	}
}

func main() {
	userInfo = map[string]string{"1":"1", "2":"2"}
	userStatus = make(map[string]string)
	userInviteSession = make(map[string]pb.ChatTask_GetInviteNotifyServer)
	userChatSession = make(map[string]pb.ChatTask_ChatMessageServer)
	//userStatus = make(map[string]pb.ChatTask_GetInviteNotifyServer)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatTaskServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}