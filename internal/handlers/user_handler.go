package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xieyx/game-server-ai/internal/models"
	"github.com/xieyx/game-server-ai/internal/services"
)

// UserHandlerInterface defines the interface for user handler
type UserHandlerInterface interface {
	CreateUser(c *gin.Context)
	Login(c *gin.Context)
	GetUser(c *gin.Context)
}

// UserHandler implements UserHandlerInterface
type UserHandler struct {
	userService services.UserServiceInterface
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(userService services.UserServiceInterface) UserHandlerInterface {
	return &UserHandler{
		userService: userService,
	}
}

// CreateUser creates a new user
func (h *UserHandler) CreateUser(c *gin.Context) {
	var input models.UserCreateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.CreateUser(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return user response without password
	userResponse := &models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusCreated, userResponse)
}

// Login handles user login
func (h *UserHandler) Login(c *gin.Context) {
	var input models.UserLoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.AuthenticateUser(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	// TODO: Generate JWT token

	// Return user response without password
	userResponse := &models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusOK, userResponse)
}

// GetUser gets a user by ID
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	// In a real implementation, we would parse the ID and fetch the user
	// For now, we'll just return a mock response
	user.ID = 1
	user.Username = "testuser"
	user.Email = "test@example.com"

	// Return user response without password
	userResponse := &models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusOK, userResponse)
}
