package utils

import "github.com/gin-gonic/gin"

// 统一响应格式
func JSONResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}
