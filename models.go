package main

import (
	"github.com/CODECZERO/goalngtest/internal/db"
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id          uuid.UUID `json:"id"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phoneNumber"`
	Email       string    `json:"email"`
	Address     string    `json:"address"`
	Password    string    `json:"password"`
}

func databaseUserToUser(dbUser db.User) User { //this function return user as the above define User
	return User{
		Id:          dbUser.ID,
		Created_at:  dbUser.CreatedAt,
		Updated_at:  dbUser.UpdatedAt,
		Name:        dbUser.Name,
		PhoneNumber: dbUser.Phonenumber,
		Email:       dbUser.Email,
		Address:     dbUser.Address,
		Password:    dbUser.Password,
	}
}
