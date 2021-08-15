package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AuthManager struct {
	secret []byte
}

type authHeader struct {
	Authorization string `header:"Authorization,required"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewAuthManager(secret string) *AuthManager {
	return &AuthManager{[]byte(secret)}
}

func (manager *AuthManager) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := authHeader{}

		if err := c.BindHeader(&header); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		authParts := strings.Split(header.Authorization, " ")

		if len(authParts) < 2 || authParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
			return
		}

		// Initialize a new instance of `Claims`
		claims := &Claims{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		tkn, err := jwt.ParseWithClaims(authParts[1], claims, func(token *jwt.Token) (interface{}, error) {
			return manager.secret, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid signature"})
				return
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}
		if !tkn.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Next()
	}
}

func (manager *AuthManager) GetToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(manager.secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
