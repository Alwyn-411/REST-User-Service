package main

import (
	"context"
	"go-user-api/internal/database"
	"go-user-api/internal/handlers"
	"go-user-api/internal/routes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main(){

	mongoURI := "mongodb://localhost:27017"

	db, err := database.NewMongoDBConnection(mongoURI)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB Successfully")

	UserHandler := handlers.NewUserHandler(db)

	router := mux.NewRouter()

	routes.SetupRoutes(router, UserHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Println("Server started on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start: ", err)
	}

	defer db.Close(context.Background())

	select {}
}