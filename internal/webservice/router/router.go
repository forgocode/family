package router

import (
	"github.com/gin-gonic/gin"

	"github.com/forgocode/family/internal/webservice/controller/category"
	"github.com/forgocode/family/internal/webservice/controller/comment"
	"github.com/forgocode/family/internal/webservice/controller/like"
	"github.com/forgocode/family/internal/webservice/controller/statistic"
	"github.com/forgocode/family/internal/webservice/controller/system"
	"github.com/forgocode/family/internal/webservice/controller/tag"
	"github.com/forgocode/family/internal/webservice/controller/topic"
	"github.com/forgocode/family/internal/webservice/controller/user"
	"github.com/forgocode/family/internal/webservice/controller/web_im"
	"github.com/forgocode/family/internal/webservice/middleware"
)

func Start() {
	engine := gin.New()
	engine.Use(middleware.Logger(), middleware.Recovery())
	//engine.Use(gin.Logger())

	//游客
	{
		engine.POST("/register", system.Register)
		engine.POST("/login", system.Login)
		//查看短评
		engine.GET("/comment", comment.UserGetComment)
		engine.GET("/article")
		// 获取评论
	}
	//普通用户
	{
		normalUserRouter := engine.Group("/normalUser")

		normalUserRouter.Use(middleware.AuthNormal())

		//获取用户的所有个人信息
		normalUserRouter.GET("info")

		//查看所有标签，已启用
		normalUserRouter.GET("/tags", tag.NormalGetAllTag)
		//查看所有分类
		normalUserRouter.GET("/category", category.NormalGetAllCategory)
		//查看所有话题
		normalUserRouter.GET("/topic", topic.NormalGetAllTopic)
		//新建文章
		normalUserRouter.POST("/article")
		//删除文章
		normalUserRouter.DELETE("/article")
		//更新文章，可隐藏
		normalUserRouter.PUT("/article")
		//新建评论
		normalUserRouter.POST("/comment", comment.UserCreateComment)
		// 赞评论
		normalUserRouter.PUT("/comment")
		normalUserRouter.GET("/ws", web_im.ReceiveClientComm)

		//赞
		normalUserRouter.POST("/like", like.GiveLike)
		//踩
		normalUserRouter.POST("/unlike", like.GiveLike)

		//发送好友请求

		//删除好友

	}

	//管理员
	{
		adminRouter := engine.Group("/admin")

		// adminRouter.Use(middleware.AuthAdmin())
		//(解)封禁用户
		adminRouter.PUT("/user/ban")
		// 新增用户
		adminRouter.POST("/user", user.AdminCreateUser)
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

		adminRouter.GET("/topic", topic.AdminGetAllTopic)
		adminRouter.POST("/topic", topic.AdminCreateTopic)
		//更新话题是否启用
		adminRouter.PUT("/topic", topic.AdminUpdateTopic)
		//删除话题
		adminRouter.DELETE("/topic", topic.AdminDeleteTopic)

		//审核所有文章，封禁文章
		adminRouter.PUT("/article")

		//查看系统日志
		adminRouter.GET("/systemLog")
		//查看操作日志
		adminRouter.GET("/operationLog", system.AdminGetOperationLog)

		//获取标签、分类、、文章、短评总数
		adminRouter.GET("/statistic/counts", statistic.Counts)
		adminRouter.GET("/statistic/usertrend", statistic.UserAddTrend)
		adminRouter.GET("/statistic/articletrend", statistic.ArticleAddTrend)
		adminRouter.GET("/statistic/topictop5", statistic.TopicTOP5)
		adminRouter.GET("/statistic/tagtop5", statistic.TagTOP5)
		adminRouter.GET("/statistic/categorytop5", statistic.CategoryTOP5)
		adminRouter.GET("/statistic/scoretop10", statistic.ScoreTop10)
		adminRouter.GET("/statistic/userActive30", statistic.UserActive30)

		adminRouter.GET("/version", system.GetVersion)
		adminRouter.GET("/monitor", system.GetMonitor)
	}
	//超级管理员
	{

		//管理员 管理
	}
	engine.Run(":8800")
}
