package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Comments struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"type:text;not null"`
	PostID  uint   `gorm:"not null"` // 指向 Post 的 ID

	// 如果你希望记录评论者是谁，可以加 UserID 和 User 字段
	// UserID  uint
	// User    User    `gorm:"foreignKey:UserID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Comments) TableName() string {
	return "comments"
}
func (c *Comments) AfterDelete(tx *gorm.DB) (err error) {
	// 目标：评论被删除后，检查对应文章的评论数量
	// 如果评论数为 0，则将文章的 CommentStatus 设置为 "无评论"

	if c.PostID == 0 {
		return nil // 未关联文章，不做处理
	}

	var commentCount int64
	err = tx.Model(&Comments{}).Where("post_id = ?", c.PostID).Count(&commentCount).Error
	if err != nil {
		return err
	}

	commentStatus := "有评论"
	if commentCount == 0 {
		// 评论数为 0，更新文章的 CommentStatus 为 "无评论"
		commentStatus = "无评论"
		// 可选：打印日志
		fmt.Printf("文章 %d 的评论已全部删除，状态已设为 '无评论'\n", c.PostID)
	}

	err = tx.Model(&Posts{}).Where("id = ?", c.PostID).Update("comment_status", commentStatus).Error
	if err != nil {
		return err
	}

	return nil
}
