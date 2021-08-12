package main

import (
	"net/http"

	"adomeit.xyz/recipe/ent"
	"adomeit.xyz/recipe/ent/user"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

// SetupUserRoutes takes the gin engine and creates routes for user sign up and login
func SetupUserRoutes(r *gin.Engine, client *ent.Client) {
	userRoute := r.Group("/user")
	{
		userRoute.POST("/", HandleCreateUser(client))
	}
	r.POST("/login", HandleLogin(client))
}

// HandleCreateUser creates a route handler to allow registering of a user
func HandleCreateUser(client *ent.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		passwordHash, err := HashPassword(user.Password)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		createdUser, err := client.User.Create().
			SetUsername(user.Username).
			SetEmail(user.Email).
			SetPassword(passwordHash).
			Save(c)

		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "User or email already taken"})
			return
		}

		createdUser.ID = 0
		createdUser.Password = ""

		c.JSON(200, createdUser)
	}
}

// HandleLogin creates a route handler that checks if a user/password combination is valid
func HandleLogin(client *ent.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var login User
		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		foundUser := client.User.
			Query().
			Where(user.Username(login.Username)).
			FirstX(c)

		if !CheckPasswordHash(login.Password, foundUser.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid username/password",
			})
			return
		}

		foundUser.ID = 0
		foundUser.Password = ""

		c.JSON(200, foundUser)
	}
}
