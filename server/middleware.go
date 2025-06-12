package server

import (
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var password = os.Getenv("PASSWORD")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("todo_auth")
		if err == nil && cookie.Value == password {
			next.ServeHTTP(w, r)
			return
		}

		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
