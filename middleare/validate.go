package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	initializers "github.com/edwinnambaje/initizializers"
	"github.com/edwinnambaje/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("SECRET"))

func handleValidationError(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	c.Abort()
}

func ValidateToken(c *gin.Context) {
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		handleValidationError(c)
		return
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("UNEXPECTED SIGNING METHOD: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			handleValidationError(c)
			return
		}
		var user models.User
		result := initializers.DB.Where("id = ?", claims["sub"]).First(&user)
		if result.Error != nil {
			handleValidationError(c)
			return
		}
		c.Set("user", user)
		c.Next()
	} else {
		handleValidationError(c)
		return
	}
}
