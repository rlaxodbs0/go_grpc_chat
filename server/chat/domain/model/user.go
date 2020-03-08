package model

import (
	"go_grpc_chat/pb"
)

type User struct {
	Session pb.ChatTask_ChatServer
	HostIP string
	Username string
	Password string
	Sex string
	ChatGroupList []*ChatGroup
	Status
}

func (c *User) GetUsername() string {
	return c.Username
}

func (c *User) GetHostIP() string {
	return c.HostIP
}

func (c *User) GetChatGroupList() []*ChatGroup {
	return c.ChatGroupList
}

func (c *User) GetChatSession() pb.ChatTask_ChatServer {
	return c.Session
}

