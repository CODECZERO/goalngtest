package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/CODECZERO/goalngtest/internal/db"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerFeedFollowsGet(w http.ResponseWriter, r *http.Request, user db.User) {
	feedsFollwer, err := cfg.DB.GetFeedFollowsForUser(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get feeds follwer")
		return
	}

	responseWithJson(w, http.StatusOK, feedsFollwer)
}

func (cfg *apiConfig) handlerFeedFollowsCreate(w http.ResponseWriter, r *http.Request, user db.User) {
	type parameters struct {
		FeedID uuid.UUID
	}

	decoder := json.NewDecoder(r.Body)
	param := parameters{}
	err := decoder.Decode(&param)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode the parameters")
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), db.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    param.FeedID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't able to follow")
		return
	}

	responseWithJson(w, http.StatusOK, databaseFeedFollowToFeedFollow(feedFollow))

}

func (cfg *apiConfig) handlerDeleteFollwer(w http.ResponseWriter, r *http.Request, user db.User) {
	feedFollwer := chi.URLParam(r, "feedFollowID")
	feedFollwerID, err := uuid.Parse(feedFollwer)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid feed follow ID")
		return
	}

	err = cfg.DB.DeleteFeedFollow(r.Context(), db.DeleteFeedFollowParams{
		UserID: user.ID,
		ID:     feedFollwerID,
	})

	if err!=nil{
		respondWithError(w,http.StatusInternalServerError,"couldn't create feed follow ")
		return
	}

	responseWithJson(w,http.StatusOK,struct{}{})

}
