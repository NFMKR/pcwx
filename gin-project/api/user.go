package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"gin-project/config"
	"gin-project/middleware"
	"gin-project/models"
)

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	Username    string              `json:"username" binding:"required"`
	Password    string              `json:"password" binding:"required"`
	Role        models.UserRole     `json:"role" binding:"required"`
	Permissions []models.Permission `json:"permissions"`
}

// Register 超级管理员注册
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查用户是否已存在
	collection := config.DB.Collection("users")
	var existingUser models.User
	err := collection.FindOne(context.Background(), bson.M{"username": req.Username}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	// 对密码进行加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// 创建超级管理员用户
	user := models.User{
		Username:  req.Username,
		Password:  string(hashedPassword),
		Role:      models.RoleSuperAdmin,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    true,
	}

	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}

// Login 用户登录
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := config.DB.Collection("users")
	var user models.User
	err := collection.FindOne(context.Background(), bson.M{"username": req.Username}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// 生成JWT token
	token, err := middleware.GenerateToken(user.ID, string(user.Role))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// 返回token和用户信息
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功",
		"data": gin.H{
			"username": user.Username,
			"role":     user.Role,
		},
		"token": token,
	})
}

// CreateSubUser 创建下级用户
func CreateSubUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 获取当前用户信息
	userID, _ := c.Get("userID")
	role, _ := c.Get("role")

	// 检查权限
	switch role {
	case string(models.RoleSuperAdmin):
		if req.Role != models.RoleSecondAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Super admin can only create second admin"})
			return
		}
	case string(models.RoleSecondAdmin):
		if req.Role != models.RoleThirdAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Second admin can only create third admin"})
			return
		}
	case string(models.RoleThirdAdmin):
		if req.Role != models.RoleUser {
			c.JSON(http.StatusForbidden, gin.H{"error": "Third admin can only create normal users"})
			return
		}
	default:
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		return
	}

	// 检查用户名是否已存在
	collection := config.DB.Collection("users")
	var existingUser models.User
	err := collection.FindOne(context.Background(), bson.M{"username": req.Username}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	// 对密码进行加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// 创建新用户
	user := models.User{
		Username:    req.Username,
		Password:    string(hashedPassword),
		Role:        req.Role,
		Permissions: req.Permissions,
		ParentID:    userID.(primitive.ObjectID),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Status:      true,
	}

	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}

// GetSubUsers 获取下级用户列表
func GetSubUsers(c *gin.Context) {
	userID, _ := c.Get("userID")

	collection := config.DB.Collection("users")
	cursor, err := collection.Find(context.Background(), bson.M{"parent_id": userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}
	defer cursor.Close(context.Background())

	var users []models.User
	if err = cursor.All(context.Background(), &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
