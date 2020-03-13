package model

import "time"

type ChatType int32

const (
	Chat_Broadcast ChatType = 0
	Chat_Message ChatType = 1
	Chat_Invite ChatType = 2
)

type Chat struct {
	Username string
	Content string
	TimeStamp time.Time
	GroupID Group
}