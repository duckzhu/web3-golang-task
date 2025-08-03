package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`                    // 用户ID，主键
	Name      string    `gorm:"size:100;not null"`             // 用户名
	Email     string    `gorm:"size:100;uniqueIndex;not null"` // 邮箱，唯一
	PostCount int       `gorm:"default:0"`                     // 统计该用户发布的文章数
	CreatedAt time.Time // GORM 默认会包含 CreatedAt, UpdatedAt, DeletedAt（如果你启用）
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // 可选：软删除

	// 一对多：一个用户有多篇文章
	Posts []Posts `gorm:"foreignKey:AuthorID"` // Post 表中的 AuthorID
}

func (User) TableName() string {
	return "user"
}
