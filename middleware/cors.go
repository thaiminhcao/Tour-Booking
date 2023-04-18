package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

func Cors() rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
			if r.Method == http.MethodOptions {
				return
			}
			next(w, r)
		}
	}
}
