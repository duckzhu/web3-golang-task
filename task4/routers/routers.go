package routers

import (
	"task4/controller"
	"task4/middlewares"

	"github.com/gin-gonic/gin"
)

// type Routers struct{}

func SetUserApi(r *gin.Engine) {

	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", controller.UserCon{}.Register)
		userRouter.POST("/login", controller.UserCon{}.Login)
		userRouter.GET("/test", controller.UserCon{}.Test)
	}

}

func SetPostApi(r *gin.Engine) {

	postRouter := r.Group("/post")
	postRouter.Use(middlewares.JWTAuthMiddleware()) // 路由分组使用中间件
	{
		postRouter.POST("/create", controller.PostCon{}.PostCreate)
		postRouter.POST("/update", controller.PostCon{}.PostUpdate)
		postRouter.POST("/list", controller.PostCon{}.PostQueryList)
		postRouter.POST("/detail", controller.PostCon{}.PostQueryDetail)
		postRouter.POST("/delete", controller.PostCon{}.PostDelete)
	}

}

func SetCommentApi(r *gin.Engine) {

	commentRouter := r.Group("/comment")
	commentRouter.Use(middlewares.JWTAuthMiddleware()) // 路由分组使用中间件
	{
		commentRouter.POST("/create", controller.CommentCon{}.CommentCreate)
		commentRouter.POST("/list", controller.CommentCon{}.CommentQuery)
	}

}
