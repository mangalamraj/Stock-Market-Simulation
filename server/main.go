package main

import (
	"log"
	"net/http"
	"stock-market-simulation/db"
	"stock-market-simulation/middleware"
	"stock-market-simulation/routes"
)

func main() {
	router := routes.UserRouter()
	handler := middleware.CorsMiddleware(router)
	
	db.InitRedis()
	
	if err := db.ConnectToMongo("mongodb+srv://mango26june:mango123@cluster0.ga9pq.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
