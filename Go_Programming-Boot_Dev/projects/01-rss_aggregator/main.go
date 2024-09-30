package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/go-chi/cors"
)

func main() {
	fmt.Println("Hiho! :)")
	godotenv.Load(".env") //load all params into local environment
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found in env.")
	}
	fmt.Println("Port:", portString)
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*","http://*"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"}
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("Server starnig on Port %v.", portString)
	// ListenAndServe runs forever, if nothing gets in-between -> if something happens: report
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
