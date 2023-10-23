package router

import (
	"github.com/gin-gonic/gin"

	"github.com/forgocode/family/internal/webservice/controller/system"
	"github.com/forgocode/family/internal/webservice/middleware"
)

func Start() {
	engine := gin.New()
	engine.Use(middleware.Logger(), middleware.Recovery())
	//engine.Use(gin.Logger())

	engine.POST("/register", system.Register)
	engine.POST("/login", system.Login)
	//游客
	{
		commentRouter := engine.Group("/comment")
		commentRouter.Use()
		commentRouter.GET("/comment")
		commentRouter.GET("/article")
	}
	//普通用户
	{
		normalUserRouter := engine.Group("/normalUser")
		normalUserRouter.Use()
		//查看所有标签
		normalUserRouter.GET("/tags")
		//查看所有分类
		normalUserRouter.GET("/category")
		//新建文章
		normalUserRouter.POST("/article")
		//删除文章
		normalUserRouter.DELETE("/article")
		//更新文章，可隐藏
		normalUserRouter.PUT("/article")
	}

	//管理员
	{
		adminRouter := engine.Group("/admin")
		adminRouter.Use()
		//新建用户（管理员？）
		adminRouter.POST("/user/add")
		//(解)封禁用户
		adminRouter.PUT("/user/ban")

		//新建标签
		adminRouter.POST("/tags")
		//删除标签
		adminRouter.DELETE("/tags")

		//新建分类
		adminRouter.POST("/category")
		//删除分类
		adminRouter.DELETE("/category")

		//审核所有文章，封禁文章
		adminRouter.PUT("/article")

		//查看系统日志
		adminRouter.GET("/systemLog")
		//查看操作日志
		adminRouter.GET("/operationLog")
	}
	engine.Run(":8800")
}
