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
	ApiKey      string    `json:"ApiKey"`
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
		ApiKey:      dbUser.Apikey,
	}
}

type Feeds struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`

}

func databaseFeedToFeed(dbFeed db.Feed) Feeds {
	return Feeds{
		ID: dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name: dbFeed.Name,
		Url: dbFeed.Url,
		UserID: dbFeed.UserID,
	}
}
