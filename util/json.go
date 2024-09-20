package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data,err:=json.Marshal(payload);
	if err==nil{
		log.Fatal("something went wrong while sending json in response")
		w.WriteHeader(500);
		return;
	}

	w.Header().Add("Content-Type","application/json");
	w.WriteHeader(code);
	w.Write(data);
}
