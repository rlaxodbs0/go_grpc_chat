package model

import (
	"go_grpc_chat/pb"
	"time"
)

type User struct {
	ID uint64
	Username string
	Password string
	GroupIDSlice []uint64
	FriendIDSlice []uint64
	*Status
	Session pb.ChatTask_ChatServer
}

type Status struct {
	Active bool
	LoginTime time.Time
	LogoutTime time.Time
}

