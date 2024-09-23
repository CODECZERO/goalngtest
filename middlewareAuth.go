package main

import (
	"net/http"
	"fmt"

	"github.com/CODECZERO/goalngtest/internal/auth"
	"github.com/CODECZERO/goalngtest/internal/db"
)

type authHeader func(http.ResponseWriter, *http.Request, db.User)

func (apiConfig *apiConfig) middleware(handler authHeader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetApiKey(r.Header)

		if err != nil {
			responseWithJson(w, 403, fmt.Sprint("Auth error:%v", err))
			return
		}
		user, err := apiConfig.DB.GetUser(r.Context(), apikey)

		if err != nil {
			responseWithJson(w, 404, fmt.Sprint("user not found:%v", err))
			return
		}

		handler(w,r,user)
	}
}
