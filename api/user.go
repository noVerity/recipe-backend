package main

import (
	"net/http"

	"adomeit.xyz/recipe/ent"
	"adomeit.xyz/recipe/ent/user"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func UserModelToResponse(user *ent.User) UserResponse {
	return UserResponse{
		Username: user.Username,
		Email:    user.Email,
	}
}

type UserController struct {
	router *gin.Engine
	client *ent.Client
}

// NewUserController takes the gin engine and creates routes for user sign up and login
func NewUserController(r *gin.Engine, client *ent.Client) *UserController {
	controller := UserController{r, client}
	userRoute := r.Group("/user")
	{
		userRoute.POST("", controller.HandleCreateUser)
	}
	r.POST("/login", controller.HandleLogin)
	return &controller
}

// HandleCreateUser creates a route handler to allow registering of a user
func (controller *UserController) HandleCreateUser(c *gin.Context) {
	var user UserRequest
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passwordHash, err := HashPassword(user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := controller.client.User.Create().
		SetUsername(user.Username).
		SetEmail(user.Email).
		SetPassword(passwordHash).
		Save(c)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User or email already taken"})
		return
	}

	c.JSON(http.StatusOK, UserModelToResponse(createdUser))
}

// HandleLogin creates a route handler that checks if a user/password combination is valid
func (controller *UserController) HandleLogin(c *gin.Context) {
	var login UserRequest
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	foundUser := controller.client.User.
		Query().
		Where(user.Username(login.Username)).
		FirstX(c)

	if !CheckPasswordHash(login.Password, foundUser.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username/password",
		})
		return
	}

	c.JSON(http.StatusOK, UserModelToResponse(foundUser))
}
