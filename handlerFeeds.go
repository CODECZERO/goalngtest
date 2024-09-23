package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/CODECZERO/goalngtest/internal/db"
)

func (cfg *apiConfig)handlerFeed(w http.ResponseWriter,r *http.Request,user db.User){
	type parameters struct{
		Name string `json:"name"`
		Url string `json:"url"`
	}

	data:=json.NewDecoder(r.Body);
	param:=parameters{}
	err:=data.Decode(&param)
	if err!=nil{
		respondWithError(w,http.StatusInternalServerError,"couldn't decode parameters")
		return
	}
	feed,err:=cfg.DB.CreateFeeds(r.Context(),db.CreateFeedsParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		Name: param.Name,
		Url: param.Url,
	})
	if err!=nil{
		respondWithError(w,http.StatusInternalServerError,"couldn't created feed")
		return
	}

	responseWithJson(w,http.StatusOK,databaseFeedToFeed(feed))
}