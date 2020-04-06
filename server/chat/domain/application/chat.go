package application

import (
	"fmt"
	"go_grpc_chat/pb"
	"go_grpc_chat/server/chat/domain/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"sync"
)

type ChatApplication struct {
	broadcast chan *pb.Message
	inviteMessage chan *pb.Message
	message chan *pb.Message
	userSession sync.Map
	repository repository.UserRepositoryImpl
}

func InitChatApplication(repo repository.UserRepositoryImpl) *ChatApplication {
	return &ChatApplication{
		broadcast: make(chan *pb.Message, 1000),
		inviteMessage: make(chan *pb.Message, 1000),
		message: make(chan *pb.Message, 1000),
		userSession:sync.Map{},
		repository: repo}
}

func (s *ChatApplication) Search(in *pb.UserInfo) map[string]string{
	userMap := map[string]string{}
	activeUserPointerSlice := s.repository.GetActiveUserPointerSlice(in.Username)
	fmt.Println(activeUserPointerSlice)
	for _, v := range activeUserPointerSlice {
		userMap[v.Username] = fmt.Sprintf("Login at %s", v.Status.LoginTime)
	}
	return userMap
}


func (s *ChatApplication) Chat(stream pb.ChatTask_ChatServer) error {
	req, _:= stream.Recv()
	_, ok := s.userSession.Load(req.Username); if !ok{
		s.userSession.Store(req.Username, stream)
	}
	defer s.closeStream(req.Username)

	go s.Broadcast()
	for{
		req, err := stream.Recv()
		if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
			return nil
		} else if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		switch req.Event.(type) {
		case *pb.Message_BroadcastMessage_:
			s.broadcast <- req
		case *pb.Message_ChatMessage:
			s.message <- req
		case *pb.Message_InviteMessage_:
			s.inviteMessage <- req
		default:
		}
	}
	<-stream.Context().Done()
	return stream.Context().Err()
}

func (s *ChatApplication) Broadcast(){
	for {
		select {
		case res := <- s.broadcast:
			s.userSession.Range(func(k, session interface{}) bool {
				if k != res.Username {
					if r, ok := status.FromError(session.(pb.ChatTask_ChatServer).Send(res)); ok {
						switch r.Code() {
						case codes.OK:
							return true
						case codes.Unavailable, codes.Canceled, codes.DeadlineExceeded:
							s.userSession.Delete(k)
							return false
						}
					}
				}
				return true
			})
		}
	}
}

func (s *ChatApplication) closeStream(username string) {
	_, ok := s.userSession.Load(username); if ok{
		s.userSession.Delete(username)
	}
}