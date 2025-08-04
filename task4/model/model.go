package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string    `gorm:"unique;not null" json:"username"`
	Password  string    `gorm:"not null" json:"password"`
	Email     string    `gorm:"unique;not null" json:"email"`
	CreatedAt time.Time `json:"createTime"`
	UpdatedAt time.Time `json:"updateTime"`
}

type Post struct {
	gorm.Model
	Title     string `gorm:"not null" json:"title"`
	Content   string `gorm:"not null" json:"content"`
	UserID    uint   `gorm:"column:user_id" json:"userId"`
	User      User
	CreatedAt time.Time `json:"createTime"`
	UpdatedAt time.Time `json:"updateTime"`
}

type Comment struct {
	gorm.Model
	Content   string `gorm:"not null" json:"content"`
	UserID    uint   `gorm:"column:user_id" json:"userId"`
	User      User
	PostID    uint `gorm:"column:post_id" json:"postId"`
	Post      Post
	CreatedAt time.Time `json:"createTime"`
	UpdatedAt time.Time `json:"updateTime"`
}

func init() {

	// 自动迁移模型
	DB.AutoMigrate(&User{}, &Post{}, &Comment{})
}
