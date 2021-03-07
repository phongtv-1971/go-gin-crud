package middleware

import "github.com/gin-gonic/gin"

// GoMiddleware ...
type GoMiddleware struct {
}

// CORS will handle the CORS middleware
func (m *GoMiddleware) CORS() gin.HandlerFunc {
	return func(g *gin.Context) {
		g.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		g.Next()
	}
}

// InitMiddleware initialize the middleware
func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}
