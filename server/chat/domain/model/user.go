package model

import (
	"time"
)

type User struct {
	Username string `gorm:"primary_key"`
	Password string	`gorm:"not null"`
	Friends []*User	`gorm:"many2many:friendships;association_jointable_foreignkey:friend_id"`
	Status Status
}

type Status struct {
	Active bool
	LastLoginTime time.Time
	LastLogoutTime time.Time
}

func (usr *User) Login()  {
	usr.Status.Active = true
	usr.Status.LastLoginTime = time.Now()
}

func (usr *User) Logout()  {
	usr.Status.Active = false
	usr.Status.LastLogoutTime = time.Now()
}