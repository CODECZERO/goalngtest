package main

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	user_id     uuid.UUID `json:"id"`
	created_at  time.Time `json:"created_at"`
	updated_at  time.Time `json:"updated_at"`
	name        string    `json:"name"`
	phoneNumber string    `json:"phoneNumber"`
	email       string    `json:"email"`
	address     string    `json:"address"`
	password    string    `json:"password"`
}

func databaseUserToUser(dbUser database.User) User { //this function return user as the above define User
	return User{
		user_id:dbUser.user_id,
		created_at:dbUser.created_at,
		updated_at:dbUser.updated_at,
		name:dbUser.name,
		phoneNumber:dbUser.phoneNumber,
		email:dbUser.email,
		address:dbUser.address,
		password:dbUser.password,
	}
}
