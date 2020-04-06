package model

import "time"

type Group struct {
	ID uint64
	Name string
	UsernameList []string
	Chat *[]Chat
}

type Chat struct {
	Username string
	Content string
	TimeStamp time.Time
}