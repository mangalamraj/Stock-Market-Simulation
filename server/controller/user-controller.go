package controller

import (
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
