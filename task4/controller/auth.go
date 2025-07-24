package controllers

import (
	"go_gin/models"
	"go_gin/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// 注册请求体
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

// 登录请求体
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 用户注册
func Register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var req RegisterRequest

	// 参数校验
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "参数错误", nil)
		return
	}

	// 检查用户名是否已存在
	var user models.User
	if db.Where("username = ?", req.Username).First(&user).Error == nil {
		utils.JSONResponse(c, http.StatusBadRequest, "用户名已存在", nil)
		return
	}

	// 创建用户
	newUser := models.User{
		Username: req.Username,
		Password: utils.HashPassword(req.Password),
		Email:    req.Email,
	}

	if err := db.Create(&newUser).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "注册失败", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "注册成功", nil)
}

// 用户登录
func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "参数错误", nil)
		return
	}

	// 验证用户
	var user models.User
	if db.Where("username = ?", req.Username).First(&user).Error != nil {
		utils.JSONResponse(c, http.StatusUnauthorized, "用户名或密码错误", nil)
		return
	}

	// 验证密码
	if user.Password != utils.HashPassword(req.Password) {
		utils.JSONResponse(c, http.StatusUnauthorized, "用户名或密码错误", nil)
		return
	}

	// 生成JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 48).Unix(),
	})

	tokenString, err := token.SignedString([]byte("xY8$kLp2@qWr7!sDf9#gHj3%kLz4^mNv6*pBn8&rTm0(fGh5)jKl1"))
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "生成token失败", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "登录成功", gin.H{"token": tokenString})
}
