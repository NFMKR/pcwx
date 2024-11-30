package routes

import (
	"github.com/gin-gonic/gin"

	"gin-project/api"
	"gin-project/middleware"
	"gin-project/models"
)

// SetupRouter 配置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 公开路由
	public := r.Group("/api")
	{
		public.POST("/register", api.Register)
		public.POST("/login", api.Login)
	}

	// 需要认证的路由
	authorized := r.Group("/api")
	authorized.Use(middleware.JWTAuth())
	{
		// 超级管理员路由
		superAdmin := authorized.Group("/super")
		superAdmin.Use(middleware.CheckRole(string(models.RoleSuperAdmin)))
		{
			superAdmin.POST("/users", api.CreateSubUser) // 创建二级管理员
			superAdmin.GET("/users", api.GetSubUsers)    // 获取二级管理员列表
		}

		// 二级管理员路由
		secondAdmin := authorized.Group("/second")
		secondAdmin.Use(middleware.CheckRole(string(models.RoleSecondAdmin)))
		{
			secondAdmin.POST("/users", api.CreateSubUser) // 创建三级管理员
			secondAdmin.GET("/users", api.GetSubUsers)    // 获取三级管理员列表
		}

		// 三级管理员路由
		thirdAdmin := authorized.Group("/third")
		thirdAdmin.Use(middleware.CheckRole(string(models.RoleThirdAdmin)))
		{
			thirdAdmin.POST("/users", api.CreateSubUser) // 创建普通用户
			thirdAdmin.GET("/users", api.GetSubUsers)    // 获取普通用户列表
		}
	}

	return r
}
