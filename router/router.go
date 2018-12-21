package router

import (
	"web/apis"
	"web/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//var auth = middleware.Tokenauth()

func InitRouter() *gin.Engine {

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	//config.AllowHeaders = []string{"Token", "Origin", "X-Requested-With", "Content-Type", "Accept"}
	config.AddAllowHeaders("Token")
	config.AddExposeHeaders("Token")

	//允许所有
	//config.AllowAllOrigins = true
	router.Use(cors.New(config))

	//使用默认配置
	//router.Use(cors.Default())

	router.GET("/", apis.IndexApi)
	router.POST("/regist", middleware.Tokenauth(), apis.AddUser)
	router.POST("/login", apis.Login)

	admin := router.Group("/admin")
	{
		admin.GET("/", apis.AdminIndex)
		admin.POST("/addpizza", middleware.Tokenauth(), apis.AddPizza)
		admin.GET("/pizzas", apis.GetPizzas)
		admin.GET("/pizza/:id", apis.GetPizza)
		admin.DELETE("/pizza/:id", apis.DelPizza)
	}
	v1 := router.Group("/v1")
	{
		v1.GET("/", apis.IndexApi)
		v1.GET("/users", apis.GetUsers)
		v1.GET("/user/:id", apis.GetUser)
	}
	return router
}
