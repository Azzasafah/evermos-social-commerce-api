package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"evermos-backend/internal/helper"
	"evermos-backend/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("token")
		if tokenStr == "" {
			authHeader := c.GetHeader("Authorization")
			if authHeader != "" {
				if strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
					tokenStr = strings.TrimSpace(authHeader[7:])
				} else {
					tokenStr = strings.TrimSpace(authHeader)
				}
			}
		}

		if tokenStr == "" {
			helper.ErrorResponse(c, http.StatusUnauthorized, "Failed to authenticate", "Token is required")
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(tokenStr)
		if err != nil {
			helper.ErrorResponse(c, http.StatusUnauthorized, "Failed to authenticate", err.Error())
			c.Abort()
			return
		}

		userIDUint64, err := strconv.ParseUint(claims.ID, 10, 32)
		if err != nil {
			helper.ErrorResponse(c, http.StatusUnauthorized, "Failed to authenticate", "Invalid user ID in token")
			c.Abort()
			return
		}
		userID := uint(userIDUint64)

		// Fetch current user from database to ensure fresh state (e.g. isAdmin updated in DB)
		var user model.User
		if err := db.First(&user, userID).Error; err != nil {
			helper.ErrorResponse(c, http.StatusUnauthorized, "Failed to authenticate", "User not found")
			c.Abort()
			return
		}

		c.Set("user_id", user.ID)
		c.Set("user_email", user.Email)
		c.Set("is_admin", user.IsAdmin)
		c.Set("user", user)

		c.Next()
	}
}
