package services

import (
	"fmt"

	"github.com/xieyx/game-server-ai/internal/database"
	"github.com/xieyx/game-server-ai/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// UserServiceInterface defines the interface for user service
type UserServiceInterface interface {
	CreateUser(input *models.UserCreateInput) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	AuthenticateUser(username, password string) (*models.User, error)
}

// UserService implements UserServiceInterface
type UserService struct{}

// NewUserService creates a new UserService
func NewUserService() UserServiceInterface {
	return &UserService{}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(input *models.UserCreateInput) (*models.User, error) {
	// Check if username already exists
	var existingUser models.User
	result := database.DB.Where("username = ?", input.Username).First(&existingUser)
	if result.Error == nil {
		return nil, fmt.Errorf("username already exists")
	}

	// Check if email already exists
	result = database.DB.Where("email = ?", input.Email).First(&existingUser)
	if result.Error == nil {
		return nil, fmt.Errorf("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user
	user := &models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	result = database.DB.Create(user)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create user: %w", result.Error)
	}

	return user, nil
}

// GetUserByID gets a user by ID
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return nil, fmt.Errorf("user not found: %w", result.Error)
	}

	return &user, nil
}

// GetUserByUsername gets a user by username
func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("user not found: %w", result.Error)
	}

	return &user, nil
}

// AuthenticateUser authenticates a user
func (s *UserService) AuthenticateUser(username, password string) (*models.User, error) {
	user, err := s.GetUserByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid username or password")
	}

	return user, nil
}
