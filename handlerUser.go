package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/CODECZERO/goalngtest/internal/auth"
	"github.com/CODECZERO/goalngtest/internal/db"
	"github.com/google/uuid"
)

func (apiCnf *apiConfig) HandlerUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct { //parameters for the data
		Name        string `json:"name"`
		PhoneNumber string `json:"phoneNumber"`
		Email       string `json:"email"`
		Address     string `json:"address"`
		Password    string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		responseWithJson(w, 400, fmt.Sprint("error while parsing json", err))
		return
	}

	user, err := apiCnf.DB.CreateUser(r.Context(), db.CreateUserParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Name:        params.Name,
		Phonenumber: params.PhoneNumber,
		Email:       params.Email,
		Address:     params.Address,
		Password:    params.Password,
		//write db parameters here
	})

	if err != nil {
		responseWithJson(w, 500, fmt.Sprint("error while creating user ", err))
		return
	}

	responseWithJson(w, 201, databaseUserToUser(user)) // here this function will return value in json to client
	//the databaseUserToUser function is used to only send limit data to clinet because the data may consited of many filed
}

func (apiCfg *apiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request) {
	apikey, err := auth.GetApiKey(r.Header)

	if err != nil {
		responseWithJson(w, 403, fmt.Sprint("Auth error:%v", err))
		return
	}
	user,err:=apiCfg.DB.GetUser(r.Context(),apikey)
	
	if err!=nil{
		responseWithJson(w, 404, fmt.Sprint("user not found:%v", err))
		return
	}

	responseWithJson(w,200,databaseUserToUser(user));

}
