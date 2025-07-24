package main

import (
	"fmt"
	"go_gin/controllers"
	"go_gin/database"
	"go_gin/models"
	"go_gin/utils"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// 认证中间件
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Header获取token
		tokenString := c.GetHeader("Authorization")
		fmt.Println(tokenString)
		if tokenString == "" {
			//c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "未提供认证token"})
			utils.JSONResponse(c, http.StatusInternalServerError, "未提供认证token", nil)
			return
		}

		// 解析token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("xY8$kLp2@qWr7!sDf9#gHj3%kLz4^mNv6*pBn8&rTm0(fGh5)jKl1"), nil
		})

		if err != nil || !token.Valid {
			//c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效的token"})
			utils.JSONResponse(c, http.StatusInternalServerError, "无效的token", nil)
			return
		}

		// 获取用户ID
		claims := token.Claims.(jwt.MapClaims)
		userID := uint(claims["user_id"].(float64))

		// 保存到context
		c.Set("userID", userID)
		c.Next()
	}
}

// 数据库中间件
func dbMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func main() {
	// 初始化日志
	log.SetOutput(os.Stdout)
	log.Println("启动应用...")

	// 加载配置
	database.Connect()
	// 自动迁移模型
	database.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

	// 初始化Gin
	r := gin.Default()

	// 添加数据库中间件
	r.Use(dbMiddleware(database.DB))

	// 路由定义
	// 用户认证
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// 评论路由
	r.GET("/posts/:id/comments", controllers.GetComments)
	// 文章路由
	r.GET("/posts/:id", controllers.GetPost)
	r.GET("/posts", controllers.GetPosts)

	// 需要认证的路由
	authGroup := r.Group("/")
	authGroup.Use(authMiddleware())
	{
		authGroup.POST("/posts/:postId/comments", controllers.CreateComment)
		authGroup.DELETE("/posts/:id", controllers.DeletePost)
		authGroup.POST("/posts", controllers.CreatePost)
		authGroup.POST("/posts/update/:id", controllers.UpdatePost)
	}

	// 启动服务
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("服务启动，监听端口:", port)
	log.Fatal(r.Run(":" + port))
}
