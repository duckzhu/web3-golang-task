package controller

import (
	"fmt"
	"net/http"
	"task4/model"

	"github.com/gin-gonic/gin"
)

type PostCon struct{}

func (PostCon) PostCreate(c *gin.Context) {

	post := model.Post{}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请校验传参是否为json格式"})
		return
	}
	userid, exists := c.Get("userid")
	if exists {
		post.UserID = uint(userid.(float64))
	}
	model.DB.Create(&post)

	fmt.Println("文章新增")
	c.JSON(http.StatusOK, gin.H{
		"msg": "文章新增成功",
	})
}

func (PostCon) PostUpdate(c *gin.Context) {

	post := model.Post{}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请校验传参是否为json格式"})
		return
	}
	if post.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id不可为空"})
		return
	}
	oldPost := model.Post{}
	model.DB.Where("id = ?", post.ID).Find(&oldPost)

	useridInterface, _ := c.Get("userid")
	useridFloat, ok := useridInterface.(float64)
	if !ok {
		c.JSON(400, gin.H{"error": "用户ID类型错误,期望是数字(float64)"})
		c.Abort()
		return
	}
	userid := uint(useridFloat)

	if oldPost.UserID != userid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "这个文章只有作者才可以修改"})
		return
	}
	result := model.DB.Model(&post).
		Where("id = ?", post.ID).
		Updates(model.Post{
			Title:   post.Title,
			Content: post.Content,
		})
	if result.Error != nil {
		fmt.Println("更新失败:", result.Error)
	} else {
		fmt.Printf("更新成功，影响了 %d 行\n", result.RowsAffected)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "文章修改成功",
	})
}

func (PostCon) PostQueryList(c *gin.Context) {
	// post := model.Post{}
	// if err := c.ShouldBindJSON(&post); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "请校验传参是否为json格式"})
	// 	return
	// }
	fmt.Println("文章列表查询")

	useridInterface, _ := c.Get("userid")
	useridFloat, ok := useridInterface.(float64)
	if !ok {
		c.JSON(400, gin.H{"error": "用户ID类型错误,期望是数字(float64)"})
		c.Abort()
		return
	}
	userid := uint(useridFloat)

	postList := []model.Post{}
	model.DB.Where("user_id = ?", userid).Find(&postList)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "文章列表查询成功",
		"data": postList,
	})
}

func (PostCon) PostQueryDetail(c *gin.Context) {
	post := model.Post{}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请校验传参是否为json格式"})
		return
	}
	fmt.Println("文章详情查询")
	model.DB.Where("id = ?", post.ID).Scan(&post)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "文章详情查询成功",
		"data": post,
	})
}

func (PostCon) PostDelete(c *gin.Context) {
	post := model.Post{}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请校验传参是否为json格式"})
		return
	}
	fmt.Println("文章删除")
	var oldPost model.Post
	result := model.DB.Where("id = ?", post.ID).First(&oldPost)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "未找到该文章"})
		return
	}

	useridInterface, _ := c.Get("userid")
	useridFloat, ok := useridInterface.(float64)
	if !ok {
		c.JSON(400, gin.H{"error": "用户ID类型错误,期望是数字(float64)"})
		c.Abort()
		return
	}
	userid := uint(useridFloat)
	if oldPost.UserID != userid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "这个文章只有作者才可以删除"})
		return
	}
	model.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{
		"msg": "文章删除成功",
	})

}
