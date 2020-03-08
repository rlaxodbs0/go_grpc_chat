package application

import (
	"fmt"
	"go_grpc_chat/pb"
	"go_grpc_chat/server/chat/domain/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"sync"
)

type ChatApplictaion struct{
	broadcast chan *pb.Message
	inviteMessage chan *pb.Message
	message chan *pb.Message
	userSession sync.Map
	repository repository.ChatRepositoryImpl
}

func InitChatApplictaion(repo repository.ChatRepositoryImpl) *ChatApplictaion{
	log.Print("init chat app")
	return &ChatApplictaion{
		broadcast: make(chan *pb.Message, 1000),
		inviteMessage: make(chan *pb.Message, 1000),
		message: make(chan *pb.Message, 1000),
		userSession:sync.Map{},
		repository: repo}
}
/*
func (s *ChatApplictaion) Search(in *pb.UserInfo) *pb.LogoutResponse{
	return s.repository.Logout(&model.User{Username:in.Username, Password:in.Password})
}
*/
func (s *ChatApplictaion) Chat(stream pb.ChatTask_ChatServer) error {
	req, _:= stream.Recv()
	_, ok := s.userSession.Load(req.Username); if !ok{
		s.userSession.Store(req.Username, stream)
	}
	defer s.closeStream(req.Username)

	go s.Broadcast(stream)
	for{
		req, err := stream.Recv()
		if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
			return nil
		} else if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		switch evt := req.Event.(type) {
		case *pb.Message_BroadcastMessage_:
			log.Printf("Broadcast: %v", req.Username)
			s.broadcast <- req
		case *pb.Message_ChatMessage:
			log.Printf("%v", req.Username)
			s.message <- req
		case *pb.Message_InviteMessage_:
			log.Printf("%v", req.Username)
			s.inviteMessage <- req
		default:
			log.Printf("%v", evt)
		}
	}
	<-stream.Context().Done()
	return stream.Context().Err()
}

func (s *ChatApplictaion) Broadcast(stream pb.ChatTask_ChatServer){
	for {
		select {
		case <-stream.Context().Done():
			return
		case res := <- s.broadcast:
			if r, ok := status.FromError(stream.Send(res)); ok {
				switch r.Code() {
				case codes.OK:
					s.userSession.Range(func(k, session interface{}) bool {
						session.(pb.ChatTask_ChatServer).Send(res)
						return true
					})
				case codes.Unavailable, codes.Canceled, codes.DeadlineExceeded:
					log.Printf("client (%s) terminated connection", res.Username)
					return
				default:
					fmt.Println("BROADCAST3")
					return
				}
			}
		}
	}
}


func (s *ChatApplictaion) closeStream(username string) {
	_, ok := s.userSession.Load(username); if ok{
		s.userSession.Delete(username)
	}
}