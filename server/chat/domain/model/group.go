package model

type Group struct {
	ID uint64
	Name string
	UsernameList []string
	Chat *[]Chat
}