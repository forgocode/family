package router

import (
	"github.com/gin-gonic/gin"

	"github.com/forgocode/family/internal/webservice/controller/category"
	"github.com/forgocode/family/internal/webservice/controller/comment"
	"github.com/forgocode/family/internal/webservice/controller/system"
	"github.com/forgocode/family/internal/webservice/controller/tag"
	"github.com/forgocode/family/internal/webservice/controller/user"
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

		// normalUserRouter.Use(middleware.AuthNormal())

		//查看所有标签，已启用
		normalUserRouter.GET("/tags", tag.NormalGetAllTag)
		//查看所有分类
		normalUserRouter.GET("/category", category.NormalGetAllCategory)
		//新建文章
		normalUserRouter.POST("/article")
		//删除文章
		normalUserRouter.DELETE("/article")
		//更新文章，可隐藏
		normalUserRouter.PUT("/article")
		//新建评论
		normalUserRouter.POST("/comment", comment.UserCreateComment)
		// 获取评论
		normalUserRouter.GET("/comment", comment.UserGetComment)

	}

	//管理员
	{
		adminRouter := engine.Group("/admin")

		// adminRouter.Use(middleware.AuthAdmin())
		//(解)封禁用户
		adminRouter.PUT("/user/ban")
		//获取所有用户
		adminRouter.GET("/user", user.NormalGetAllUser)

		//新建标签
		//获取所有标签，包括不启用的
		adminRouter.GET("/tags", tag.AdminGetAllTag)
		adminRouter.POST("/tags", tag.AdminCreateTag)
		//更新标签是否启用
		adminRouter.PUT("/tags", tag.AdminUpdateTag)
		//删除标签
		adminRouter.DELETE("/tags", tag.AdminDeleteTag)

		//获取所有标签，包括不启用的
		adminRouter.GET("category", category.AdminGetAllCategory)
		//新建分类
		adminRouter.POST("/category", category.AdminCreateCategory)
		//更新分类
		adminRouter.PUT("/category", category.AdminUpdateCategory)
		//删除分类
		adminRouter.DELETE("/category", category.AdminDeleteCategory)

		//审核所有文章，封禁文章
		adminRouter.PUT("/article")

		//查看系统日志
		adminRouter.GET("/systemLog")
		//查看操作日志
		adminRouter.GET("/operationLog", system.AdminGetOperationLog)
	}
	//超级管理员
	{

		//管理员 管理
	}
	engine.Run(":8800")
}
