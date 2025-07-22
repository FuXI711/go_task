package lesson01

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:100"`
	Age   int
	Posts []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	ID       uint   `gorm:"primaryKey"`
	Title    string `gorm:"size:200"`
	Content  string `gorm:"type:text"`
	UserID   uint
	Comments []Comment `gorm:"foreignKey:PostID"`
}

type Comment struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"type:text"`
	PostID  uint
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	// //插入数据
	// user := User{
	// 	Name: "A",
	// 	Age:  25,
	// 	Posts: []Post{
	// 		{
	// 			Title:   "Post 1",
	// 			Content: "Content 1",
	// 			Comments: []Comment{
	// 				{Content: "真棒"},
	// 				{Content: "点赞"},
	// 			},
	// 		},
	// 		{
	// 			Title:   "Post 2",
	// 			Content: "Content 2",
	// 		},
	// 	},
	// }
	// db.Create(&user)

	//查询A的评论，文章，用户信息
	var result User
	err := db.Preload("Posts.Comments").
		Where("name = ?", "A").
		Find(&result).
		Error

	if err != nil {
		log.Fatalf("查询失败: %v", err)
	}

	fmt.Printf("用户 %s (ID: %d) 的文章列表:\n", result.Name, result.ID)
	for _, post := range result.Posts {
		fmt.Printf("\n文章标题: %s\n内容: %s\n", post.Title, post.Content)
		fmt.Println("评论:")
		for _, comment := range post.Comments {
			fmt.Printf("  - %s\n", comment.Content)
		}
	}

	//查询评论数最高的文章信息
	var post Post
	err = db.Debug().Select("posts.*, COUNT(comments.id) as comment_count").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count DESC").
		Preload("Comments").
		First(&post).
		Error

	if err != nil {
		log.Fatalf("查询失败: %v", err)
	}

	fmt.Printf("评论数最多的文章:\n")
	fmt.Printf("标题: %s\n内容: %s\n", post.Title, post.Content)
	fmt.Println("评论:")
	for _, comment := range post.Comments {
		fmt.Printf("  - %s\n", comment.Content)
	}
}
