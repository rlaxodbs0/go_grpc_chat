package model

import "time"

type Status struct {
	Active bool
	LoginTime time.Time
	LogoutTime time.Time
}

