package models

import "time"

var (
	UserSlice []User
)


type User struct {
	Id string
	Time time.Time
}

func (user *User) Add(){
	UserSlice = append(UserSlice, *user)
}
