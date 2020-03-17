package model

import "time"

type Chat struct {
	Username string
	Content string
	TimeStamp time.Time
	GroupID Group
}