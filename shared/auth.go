package shared

import (
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// AuthManager provides middleware and utility functions to provide bearer token authentication
type AuthManager struct {
	secret []byte
}

type authHeader struct {
	Authorization string `header:"Authorization,required"`
}

// Claims that are being required across the services
type Claims struct {
	Username string `json:"username"`
	Shard    string `json:"shard"`
	jwt.RegisteredClaims
}

func NewAuthManager(secret string) *AuthManager {
	return &AuthManager{[]byte(secret)}
}

// AuthMiddleware authenticates the user based on the bearer token and saves user info into the context
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

		c.Set("user", claims.Username)
		c.Set("userShard", claims.Shard)

		c.Next()
	}
}

func (manager *AuthManager) GetToken(username string, shard string) (string, error) {
	expirationTime := time.Now().Add(time.Hour)

	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: username,
		Shard:    shard,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
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
