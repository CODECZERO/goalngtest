package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/CODECZERO/goalngtest/util/handler"
)

func main(){
	godotenv.Load();

	port:=os.Getenv("PORT");
	if port==""{
		log.Fatal("port number is missing");
	}
	fmt.Println(port)

	router:=chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:[]string{"http://*","https://*"} ,
		AllowedMethods:[]string{"GET","PUT","POST","DELETE"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))
	v1router:=chi.NewRouter()
	v1router.HandleFunc("/healt",handler)
	
	router.Mount("/v1",v1router);

	server:=&http.Server{
		Handler: router,
		Addr: ":"+port,
	}

	server.s
	
}