package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"stock-market-simulation/db"
	"stock-market-simulation/models"
	"stock-market-simulation/types"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	collection := db.GetCollection("stock-market-simulation", "users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	
	
	var signupData types.SignupRequest
	err := json.NewDecoder(r.Body).Decode(&signupData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if signupData.Password != signupData.ConfirmPassword {
		http.Error(w, "Passwords do not match", http.StatusBadRequest)
		return
	}

	_, err = collection.InsertOne(ctx, signupData.User)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(signupData.User)
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	collection := db.GetCollection("stock-market-simulation", "users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var loginData types.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginData)

	if err != nil || loginData.Email == "" || loginData.Password == "" {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

	filter := bson.M{"email": loginData.Email}
	var user models.User
	err = collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if user.Password != loginData.Password {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(user)

}

