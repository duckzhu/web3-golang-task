package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Posts struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"size:200;not null"`
	Content   string `gorm:"type:text"`           // 文章内容，文本类型
	AuthorID  uint   `gorm:"not null"`            // 外键，指向 User 表的 ID
	Author    User   `gorm:"foreignKey:AuthorID"` // 属于哪个用户（可选，用于预加载）
	CreatedAt time.Time
	UpdatedAt time.Time

	CommentStatus string `gorm:"size:20;default:'无评论'"` // 新增：评论状态字段

	// 一对多：一篇文章有多个评论
	Comments []Comments `gorm:"foreignKey:PostID"` // Comments 表中的 PostID
}

func (Posts) TableName() string {
	return "posts"
}

func (p *Posts) AfterCreate(tx *gorm.DB) (err error) {
	// 目标：找到该文章的作者，并将其 PostCount +1

	if p.AuthorID == 0 {
		return nil // 未关联作者，不做处理
	}

	// 更新用户的 PostCount 字段
	err = tx.Model(&User{}).Where("id = ?", p.AuthorID).Update("post_count", gorm.Expr("post_count + 1")).Error
	if err != nil {
		// 可以记录日志，但不中断程序
		return err
	}
	fmt.Printf("文章创建成功，用户 %d 的文章数已更新\n", p.AuthorID)
	return nil
}
