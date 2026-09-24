package middleware

import (
	"net/http"

	"evermos-backend/internal/helper"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin, exists := c.Get("is_admin")
		if !exists || !isAdmin.(bool) {
			helper.ErrorResponse(c, http.StatusForbidden, "Failed to authorize", "Akses ditolak: Hanya admin yang diizinkan")
			c.Abort()
			return
		}
		c.Next()
	}
}
