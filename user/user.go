package main

import (
	"adomeit.xyz/shared"
	"math/rand"
	"net/http"

	"adomeit.xyz/user/ent"
	"adomeit.xyz/user/ent/user"

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
	Jwt      string `json:"jwt"`
}

type UserController struct {
	router   *gin.Engine
	client   *ent.Client
	auth     *shared.AuthManager
	shardMap *ShardMap
}

// NewUserController takes the gin engine and creates routes for user sign up and login
func NewUserController(r *gin.Engine, client *ent.Client, auth *shared.AuthManager, shardMap *ShardMap) *UserController {
	controller := UserController{r, client, auth, shardMap}
	userRoute := r.Group("/user")
	{
		userRoute.POST("", controller.HandleCreateUser)
		userRoute.GET("", auth.AuthMiddleware(), controller.HandleGetUser)
	}
	r.POST("/login", controller.HandleLogin)
	return &controller
}

func (controller *UserController) UserModelToResponse(user *ent.User) (UserResponse, error) {
	tokenString, err := controller.auth.GetToken(user.Username, user.RecipeShard)

	if err != nil {
		return UserResponse{}, err
	}

	return UserResponse{
		Username: user.Username,
		Email:    user.Email,
		Jwt:      tokenString,
	}, nil
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

	// Let's just assume 2 shards for the moment
	// Ideally the number would be assigned based on load balancing instead of random
	// otherwise if we add another shard it woldn't take a lot of load from the existing ones
	shard := rand.Intn(len(controller.shardMap.Map))

	createdUser, err := controller.client.User.Create().
		SetUsername(user.Username).
		SetEmail(user.Email).
		SetPassword(passwordHash).
		SetRecipeShard(controller.shardMap.Map[shard].Name).
		Save(c)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User or email already taken"})
		return
	}

	userResponse, err := controller.UserModelToResponse(createdUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userResponse)
}

// HandleLogin creates a route handler that checks if a user/password combination is valid
func (controller *UserController) HandleLogin(c *gin.Context) {
	var login UserRequest
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	foundUser, err := controller.client.User.
		Query().
		Where(user.Username(login.Username)).
		First(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if !CheckPasswordHash(login.Password, foundUser.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username/password",
		})
		return
	}

	userResponse, err := controller.UserModelToResponse(foundUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userResponse)
}

// HandleGetUser reads the user from context/jwt and returns if found
func (controller *UserController) HandleGetUser(c *gin.Context) {
	username := c.GetString("user")

	if username == "" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorised"})
		return
	}

	foundUser, err := controller.client.User.
		Query().
		Where(user.Username(username)).
		First(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	userResponse, err := controller.UserModelToResponse(foundUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userResponse)
}
