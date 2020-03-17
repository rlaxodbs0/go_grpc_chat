package model

import (
	"time"
)

type User struct {
	ID uint64
	Username string
	Password string
	GroupIDSlice []uint64
	FriendIDSlice []uint64
	Status Status
}

type Status struct {
	Active bool
	LoginTime time.Time
	LogoutTime time.Time
}

