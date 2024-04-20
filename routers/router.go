package routers

import (
	"gindemo/controllers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	//加载views目录下的html文件
	router.LoadHTMLGlob("views/*")

	// 设置session，初始化
	store := cookie.NewStore([]byte("loginuser"))
	router.Use(sessions.Sessions("mysession", store))

	{
		//注册
		router.GET("/register", controllers.RegisterGet)
		router.POST("/register", controllers.RegisterPost)

		// 登录
		router.GET("/login", controllers.LoginGet)
		router.POST("/login", controllers.LoginPost)

		// 首页
		router.GET("/", controllers.HomeGet)

		// 退出
		router.GET("/exit", controllers.ExitGet)

		v1 := router.Group("/question")
		{
			v1.GET("/add", controllers.AddQuestionGet)
			v1.POST("/add", controllers.AddQuestionPost)
		}

	}

	return router
}
