package application

import (
	"go_grpc_chat/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"sync"
)

type chatApplication struct {
	broadcast chan *pb.Message
	message chan *pb.Message
	userSession sync.Map
}

var ChatApplicationInstance *chatApplication

func init (){
	ChatApplicationInstance = &chatApplication{
		broadcast:   make(chan *pb.Message, 1000),
		message:     make(chan *pb.Message, 1000),
		userSession: sync.Map{},
	}
	go ChatApplicationInstance.Broadcast()
	go ChatApplicationInstance.Message()
}


func ChatApplication() *chatApplication {
	return ChatApplicationInstance
}

func (s *chatApplication) Chat(stream pb.ChatService_ChatServer) error {
	req, _:= stream.Recv()
	_, ok := s.userSession.Load(req.Username); if !ok{
		s.userSession.Store(req.Username, stream)
	}
	defer s.closeStream(req.Username)
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
		case *pb.Message_Message_:
			s.message <- req
		default:
		}
	}
	<-stream.Context().Done()
	return stream.Context().Err()
}

func (s *chatApplication) Broadcast(){
	for {
		select {
		case res := <- s.broadcast:
			s.userSession.Range(func(k, session interface{}) bool {
				if k != res.Username {
					if r, ok := status.FromError(session.(pb.ChatService_ChatServer).Send(res)); ok {
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


func (s *chatApplication) Message(){
	for {
		select {
		case res := <-s.message:
			if v, ok := s.userSession.Load(res.Event.(*pb.Message_Message_).Message.Target); ok {
				v.(pb.ChatService_ChatServer).Send(res)
			}
		}
	}
}

func (s *chatApplication) closeStream(username string) {
	_, ok := s.userSession.Load(username); if ok {
		s.userSession.Delete(username)
	}
}