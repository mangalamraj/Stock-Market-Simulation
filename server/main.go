package main

import (
	"log"
	"net/http"
	"stock-market-simulation/db"
	"stock-market-simulation/middleware"
	"stock-market-simulation/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize Redis
	db.InitRedis()

	// Connect to MongoDB
	if err := db.ConnectToMongo("mongodb+srv://mango26june:mango123@cluster0.ga9pq.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Create a unified router
	mainRouter := mux.NewRouter()

	// Register user routes
	userRouter := routes.UserRouter()
	mainRouter.PathPrefix("/user").Handler(userRouter)

	// Register stock routes
	stockRouter := routes.StockRouter(db.RedisClient)
	mainRouter.PathPrefix("/stock").Handler(stockRouter)

	// Apply CORS middleware
	handler := middleware.CorsMiddleware(mainRouter)

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
