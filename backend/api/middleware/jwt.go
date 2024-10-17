package middleware

import (
	"backend/domain"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Middleware struct {
	repo domain.UserRepository
}

func NewMiddleware(repo domain.UserRepository) Middleware {
	return Middleware{repo: repo}
}

var mySecretKey = []byte("hekp")

func CreateJWt(id uint) (string, error) {
	Claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)

	tokenString, err := token.SignedString(mySecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(token string) (uint, error) {

	tokenstr, err := jwt.Parse(token,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])

			}
			return mySecretKey, nil
		})
	if err != nil {
		return 0, err
	}

	if claims, ok := tokenstr.Claims.(jwt.MapClaims); ok && tokenstr.Valid {

		id, ok := claims["id"].(float64)

		if !ok {
			return 0, fmt.Errorf("invalid id in token")
		}
		return uint(id), nil
	}
	return 0, err

}

func (n *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")

		id, err := ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		user, err := n.repo.GetUserById(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if user == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			return
		}
		c.Set("user", user)

		c.Next()

	}

}
