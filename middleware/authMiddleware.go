package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

// create function auth middleware
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("masuk ke middleware AUTH")
		next.ServeHTTP(w, r)
		logrus.Info("keluar dari middleware AUTH")
	})
}
