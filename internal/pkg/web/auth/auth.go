package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/COMP4050/square-team-5/api/internal/pkg/db/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func getUserFromJWT(tokenString, secret string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}
	if token == nil {
		return "", errors.New("token is nil")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["sub"].(string), nil
	} else {
		return "", errors.New("token invalid")
	}
}

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt, err := c.Cookie("token")
		if err != nil || jwt == "" {
			return
		}

		username, err := getUserFromJWT(jwt, "catjam")
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		user := models.User{Email: username, Role: models.UserRoleAdmin}

		ctx := context.WithValue(c, userCtxKey, user)

		c.Request = c.Request.WithContext(ctx)
	}
}

func ExtractUser(ctx context.Context) *models.User {
	user, ok := ctx.Value(userCtxKey).(models.User)
	if !ok {
		return nil
	}

	if user.Email == "" {
		return nil
	} else {
		return &user
	}
}
