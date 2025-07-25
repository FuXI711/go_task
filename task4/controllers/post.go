package controllers

import (
	"go_gin/models"
	"go_gin/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 创建文章请求体
type CreatePostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// 获取文章列表
func GetPosts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var posts []models.Post

	if err := db.Find(&posts).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "获取文章列表失败", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "获取成功", posts)
}

// 创建文章
func CreatePost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID := c.MustGet("userID").(uint)
	var req CreatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "参数错误", nil)
		return
	}

	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}

	if err := db.Create(&post).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "创建文章失败", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "创建成功", post)
}

// 获取文章详情
func GetPost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "文章ID错误", nil)
		return
	}

	var post models.Post
	if err := db.First(&post, id).Error; err != nil {
		utils.JSONResponse(c, http.StatusNotFound, "文章不存在", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "获取成功", post)
}

// 删除文章
func DeletePost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID := c.MustGet("userID").(uint)
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "文章ID错误", nil)
		return
	}

	var post models.Post
	if err := db.First(&post, id).Error; err != nil {
		utils.JSONResponse(c, http.StatusNotFound, "文章不存在", nil)
		return
	}

	// 检查是否是作者
	if post.UserID != userID {
		utils.JSONResponse(c, http.StatusForbidden, "无权删除此文章", nil)
		return
	}

	if err := db.Delete(&post).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "删除失败", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "删除成功", nil)
}

// 更新文章
func UpdatePost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID := c.MustGet("userID").(uint)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "文章ID错误", nil)
		return
	}

	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "参数错误", nil)
		return
	}

	var post models.Post
	if err := db.First(&post, id).Error; err != nil {
		utils.JSONResponse(c, http.StatusNotFound, "文章不存在", nil)
		return
	}

	if post.UserID != userID {
		utils.JSONResponse(c, http.StatusForbidden, "无权更新此文章", nil)
		return
	}

	if err := db.Model(&post).Updates(req).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "更新失败", nil)
		return
	}

	// 6. 返回成功响应（修正信息）
	utils.JSONResponse(c, http.StatusOK, "更新成功", post)
}
