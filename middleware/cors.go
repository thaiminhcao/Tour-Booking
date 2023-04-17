package middleware

import (
	"github.com/tal-tech/go-zero/core/middleware"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func CorsMiddleware() middleware.Middleware {
	return func(next httpx.HandlerFunc) httpx.HandlerFunc {
		return func(ctx *httpx.Context) {
			ctx.Header("Access-Control-Allow-Origin", "*")
			ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			ctx.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
			ctx.Header("Access-Control-Expose-Headers", "Content-Length")
			ctx.Header("Access-Control-Allow-Credentials", "true")
			ctx.Header("Access-Control-Max-Age", "43200")
			if ctx.Request.Method == "OPTIONS" {
				ctx.NoContent()
				return
			}
			next(ctx)
		}
	}
}
