package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main(){
	router := mux.NewRouter()

	err := godotenv.Load(".env")
	if err != nil{
		log.Fatalln("Cannot load the .env file")
	}

	fs := http.FileServer(http.Dir("./public/"))

	port := os.Getenv("PORT")
	if port == ""{
		fmt.Println("it needs the godotenv package to access the .env var")
	}

	router.Handle("/", fs)
	panic(http.ListenAndServe(":"+port, router))
}