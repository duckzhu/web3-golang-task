package controller

import (
	"fmt"
	"net/http"
	"task4/model"

	"github.com/gin-gonic/gin"
)

type CommentCon struct{}

func (CommentCon) CommentCreate(c *gin.Context) {

	//评论创建
	comment := model.Comment{}
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请校验传参是否为json格式"})
		return
	}
	userid, exists := c.Get("userid")
	if exists {
		comment.UserID = uint(userid.(float64))
	}
	model.DB.Create(&comment)

	fmt.Println("评论创建")
	c.JSON(http.StatusOK, gin.H{
		"msg": "评论创建成功",
	})
}

func (CommentCon) CommentQuery(c *gin.Context) {
	//评论读取
	commentList := []model.Comment{}
	// if err := c.ShouldBindJSON(&comment); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "请校验传参是否为json格式"})
	// 	return
	// }
	userid, _ := c.Get("userid")
	fmt.Println("评论查询", userid)
	model.DB.Where("user_id=?", userid).Find(&commentList)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "评论查询成功",
		"data": commentList,
	})
}
