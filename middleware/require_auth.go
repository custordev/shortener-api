package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/custordev/shortener-api/dtos"
	"github.com/custordev/shortener-api/initializers"
	"github.com/custordev/shortener-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuthWithToken(c *gin.Context) {
	// Get token from Authorization header
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, dtos.ErrorResponse{
			Success: false,
			Error:   "Unauthorized - No token provided",
		})
		c.Abort()
		return
	}

	// Extract token from "Bearer <token>" format
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	if tokenString == authHeader {
		c.JSON(http.StatusUnauthorized, dtos.ErrorResponse{
			Success: false,
			Error:   "Unauthorized - Invalid token format. Use 'Bearer <token>'",
		})
		c.Abort()
		return
	}

	// Decode/Validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, dtos.ErrorResponse{
			Success: false,
			Error:   "Unauthorized - Invalid token",
		})
		c.Abort()
		return
	}

	// Extract claims and validate
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, dtos.ErrorResponse{
				Success: false,
				Error:   "Unauthorized - Token expired",
			})
			c.Abort()
			return
		}

		// Find user with token sub
		var user models.User
		initializers.Db.First(&user, claims["sub"])

		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, dtos.ErrorResponse{
				Success: false,
				Error:   "Unauthorized - User not found",
			})
			c.Abort()
			return
		}

		// Attach user to context
		c.Set("user", user)

		// Continue to next handler
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, dtos.ErrorResponse{
			Success: false,
			Error:   "Unauthorized - Invalid token claims",
		})
		c.Abort()
		return
	}
}
