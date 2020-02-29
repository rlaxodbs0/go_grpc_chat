package domain

import (
	"go_grpc_chat/pb"
)

type Client struct {
	pb.ChatTaskClient
	HostIP string
	UserName string
	Password string
	ChatGroupList map[string]*ChatGroup
}

func (c *Client) GetUserName() string {
	return c.UserName
}

func (c *Client) GetHostIP() string {
	return c.HostIP
}

func (c *Client) GetChatGroupList() map[string]*ChatGroup {
	return c.ChatGroupList
}