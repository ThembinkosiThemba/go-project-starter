package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetTokensAsCookies(c *gin.Context, token string, refreshToken string) {
	// Set the access token as a cookie
	c.SetCookie("token", token, 3600*2, "/", "", false, true)

	// Set the refresh token as a cookie
	c.SetCookie("refresh_token", refreshToken, 3600*24*7, "/", "", false, true)
}

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the cookie
		token, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No access token provided"})
			c.Abort()
			return
		}

		// Validate the token
		claims, msg := ValidateToken(token)
		if msg != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
			c.Abort()
			return
		}

		// Set the user information in the context
		c.Set("email", claims.Email)
		c.Set("firstname", claims.Firstname)
		c.Set("lastname", claims.Lastname)

		// Continue to the next handler
		c.Next()
	}
}
