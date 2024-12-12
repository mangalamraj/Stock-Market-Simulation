package types

import "stock-market-simulation/models"

type SignupRequest struct {
	models.User
	ConfirmPassword string `json:"confirmPassword"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Stock struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}
