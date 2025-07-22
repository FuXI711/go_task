package lesson01

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100"`
	Age       int
	PostCount int
	Posts     []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	ID       uint   `gorm:"primaryKey"`
	Title    string `gorm:"size:200"`
	Content  string `gorm:"type:text"`
	UserID   uint
	Comments []Comment `gorm:"foreignKey:PostID"`
	Status   string
}

type Comment struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"type:text"`
	PostID  uint
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&User{}).
			Where("id = ?", p.UserID).
			Update("post_count", gorm.Expr("post_count + 1"))

		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return fmt.Errorf("用户不存在")
		}

		return nil
	})
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	if err := tx.Model(&Comment{}).
		Where("post_id = ?", c.PostID).
		Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		if err := tx.Model(&Post{}).
			Where("id = ?", c.PostID).
			Update("status", "无评论").Error; err != nil {
			return err
		}
	}
	return nil
}
func Run(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	user := User{
		Name: "B",
		Age:  21,
		Posts: []Post{
			{
				Title:   "Post b",
				Content: "Content b",
				Status:  "有评论",
				Comments: []Comment{
					{Content: "非常好"},
					{Content: "收藏了"},
				},
			},
		},
	}
	db.Create(&user)

	var post Post
	if err := db.Preload("Comments").First(&post).Error; err != nil {
		fmt.Printf("查询文章失败: %v\n", err)
		return
	}

	var commentsToDelete []Comment
	if err := db.Where("post_id = ?", post.ID).Find(&commentsToDelete).Error; err != nil {
		return
	}

	for _, comment := range commentsToDelete {
		if err := db.Delete(&comment).Error; err != nil {
			fmt.Printf("删除评论ID=%d失败: %v\n", comment.ID, err)
			continue
		}
	}
	var commentCount int64
	db.Model(&Comment{}).Where("post_id = ?", post.ID).Count(&commentCount)
	fmt.Printf("文章状态: %s\n", post.Status)
}
