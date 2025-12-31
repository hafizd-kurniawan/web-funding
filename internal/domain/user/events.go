package user

import "time"

type UserRegistered struct {
	Email string
	Name  string
	At    time.Time
}
