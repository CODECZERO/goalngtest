package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (apiCnf *apiConfig) handlerUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {//parameters for the data
		name        string    `json:"name"`
		phoneNumber string    `json:"phoneNumber"`
		email       string    `json:"email"`
		address     string    `json:"address"`
		password    string    `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		responseWithJson(w, 400, fmt.Sprint("error while parsing json", err))
		return
	}

	user, err := apiCnf.DB.CreateUser(r.Context(), database.CreateUserParams{
		user_id:    uuid.New(),
		created_at: time.Now().UTC(),
		updated_at: time.Now().UTC(),
		name:params.name,
		phoneNumber:params.phoneNumber,
		email:params.email,
		address:params.address,
		password:params.password,
		//write db parameters here
	})

	if err != nil {
		responseWithJson(w, 500, fmt.Sprint("error while creating user ", err))
		return
	}

	responseWithJson(w, 200, databaseUserToUser(user)) // here this function will return value in json to client
	//the databaseUserToUser function is used to only send limit data to clinet because the data may consited of many filed

}
