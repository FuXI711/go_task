package controllers

import (
	"go_gin/models"
	"go_gin/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 创建评论请求体
type CreateCommentRequest struct {
	Content string `json:"content" binding:"required"`
}

// 创建评论
func CreateComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID := c.MustGet("userID").(uint)
	postID, err := strconv.Atoi(c.Param("postId"))

	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "文章ID错误", nil)
		return
	}

	var req CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "参数错误", nil)
		return
	}

	// 检查文章是否存在
	var post models.Post
	if err := db.First(&post, postID).Error; err != nil {
		utils.JSONResponse(c, http.StatusNotFound, "文章不存在", nil)
		return
	}

	comment := models.Comment{
		Content: req.Content,
		UserID:  userID,
		PostID:  uint(postID),
	}

	if err := db.Create(&comment).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "评论失败", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "评论成功", comment)
}

// 获取文章评论
func GetComments(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	postID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "文章ID错误", nil)
		return
	}

	var comments []models.Comment
	if err := db.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "获取评论失败", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "获取成功", comments)
}
