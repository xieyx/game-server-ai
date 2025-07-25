package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xieyx/game-server-ai/internal/models"
	"github.com/xieyx/game-server-ai/internal/services"
)

// MockDB is a mock database for testing
type MockDB struct {
	mock.Mock
}

func (m *MockDB) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockDB) GetUserByID(id uint) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockDB) GetUserByUsername(username string) (*models.User, error) {
	args := m.Called(username)
	return args.Get(0).(*models.User), args.Error(1)
}

func TestUserService_CreateUser(t *testing.T) {
	// 测试数据
	input := &models.UserCreateInput{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	// 创建mock数据库
	mockDB := new(MockDB)
	mockDB.On("GetUserByUsername", input.Username).Return(&models.User{}, nil).Once()
	mockDB.On("GetUserByEmail", input.Email).Return(&models.User{}, nil).Once()
	mockDB.On("CreateUser", mock.AnythingOfType("*models.User")).Return(nil).Once()

	// 创建用户服务
	userService := services.NewUserService()

	// 执行测试
	user, err := userService.CreateUser(input)

	// 验证结果
	assert.NoError(t, err)
	assert.Equal(t, input.Username, user.Username)
	assert.Equal(t, input.Email, user.Email)
	assert.NotZero(t, user.ID)

	// 验证mock调用
	mockDB.AssertExpectations(t)
}

func TestUserService_CreateUser_UsernameExists(t *testing.T) {
	// 测试数据
	input := &models.UserCreateInput{
		Username: "existinguser",
		Email:    "new@example.com",
		Password: "password123",
	}

	// 创建mock数据库
	mockDB := new(MockDB)
	mockDB.On("GetUserByUsername", input.Username).Return(&models.User{ID: 1, Username: "existinguser"}, nil).Once()

	// 创建用户服务
	userService := services.NewUserService()

	// 执行测试
	user, err := userService.CreateUser(input)

	// 验证结果
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "username already exists", err.Error())

	// 验证mock调用
	mockDB.AssertExpectations(t)
}

func TestUserService_GetUserByID(t *testing.T) {
	// 测试数据
	expectedUser := &models.User{
		ID:       1,
		Username: "testuser",
		Email:    "test@example.com",
	}

	// 创建mock数据库
	mockDB := new(MockDB)
	mockDB.On("GetUserByID", uint(1)).Return(expectedUser, nil).Once()

	// 创建用户服务
	userService := services.NewUserService()

	// 执行测试
	user, err := userService.GetUserByID(1)

	// 验证结果
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)

	// 验证mock调用
	mockDB.AssertExpectations(t)
}

func TestUserService_AuthenticateUser(t *testing.T) {
	// 测试数据
	username := "testuser"
	password := "password123"
	hashedPassword := "$2a$10$N.zmdr9k7uOCQb0bta/OauRxaOKSr.QhqyD2R5FKvMQjmHoLkm5Sy" // bcrypt hash of "password123"

	user := &models.User{
		ID:       1,
		Username: username,
		Email:    "test@example.com",
		Password: hashedPassword,
	}

	// 创建mock数据库
	mockDB := new(MockDB)
	mockDB.On("GetUserByUsername", username).Return(user, nil).Once()

	// 创建用户服务
	userService := services.NewUserService()

	// 执行测试
	authUser, err := userService.AuthenticateUser(username, password)

	// 验证结果
	assert.NoError(t, err)
	assert.Equal(t, user, authUser)

	// 验证mock调用
	mockDB.AssertExpectations(t)
}

func TestUserService_AuthenticateUser_InvalidPassword(t *testing.T) {
	// 测试数据
	username := "testuser"
	password := "wrongpassword"
	hashedPassword := "$2a$10$N.zmdr9k7uOCQb0bta/OauRxaOKSr.QhqyD2R5FKvMQjmHoLkm5Sy" // bcrypt hash of "password123"

	user := &models.User{
		ID:       1,
		Username: username,
		Email:    "test@example.com",
		Password: hashedPassword,
	}

	// 创建mock数据库
	mockDB := new(MockDB)
	mockDB.On("GetUserByUsername", username).Return(user, nil).Once()

	// 创建用户服务
	userService := services.NewUserService()

	// 执行测试
	authUser, err := userService.AuthenticateUser(username, password)

	// 验证结果
	assert.Error(t, err)
	assert.Nil(t, authUser)
	assert.Equal(t, "invalid username or password", err.Error())

	// 验证mock调用
	mockDB.AssertExpectations(t)
}
