package Routes

import (
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"net/http"
	"foodorder/Controllers"
	"foodorder/Middlewares"
	"foodorder/Service"
)

func SetupRouter() *gin.Engine  {
	r := gin.Default()
	r.Use(Middlewares.SignRequest)
	r.Use(sentrygin.New(sentrygin.Options{}))

	var loginService Service.LoginService = Service.StaticLoginService()
	var jwtService Service.JWTService = Service.JwtAuthService()
	var loginController Controllers.LoginController = Controllers.LoginHandler(loginService, jwtService)

	r.POST("/login", func(context *gin.Context) {
		token := loginController.Login(context)
		if token != ""{
			context.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		}else{
			context.JSON(http.StatusUnauthorized, nil)
		}
	})

	usergroup := r.Group("/user")
	{
		usergroup.GET("menu", Controllers.GetMenu)
		usergroup.POST("order", Controllers.CreateOrder)
		usergroup.GET("order/:id", Controllers.GetAllOrderByIDUser)

	}

	waitergroup := r.Group("/waiter",Middlewares.AuthorizeJWT(), Middlewares.AuthorizeUser())
	{
		waitergroup.GET("user", Controllers.GetUsers)
		waitergroup.PUT("order/:id", Controllers.UpdateOrder)

	}

	admingroup := r.Group("/admin",  Middlewares.AuthorizeJWT(), Middlewares.AuthorizeUser())
	{
		admingroup.GET("menu", Controllers.GetMenu)
		admingroup.POST("addmenu", Controllers.CreateMenu)
		admingroup.DELETE("menu/:id", Controllers.DeleteMenu)
		admingroup.PUT("menu/:id", Controllers.UpdateMenu)

		admingroup.GET("user/:id", Controllers.GetUserByID)
		admingroup.POST("user", Controllers.CreateUser)
		admingroup.PUT("user/:id", Controllers.UpdateUser)
		admingroup.DELETE("user/:id", Controllers.DeleteUser)
	}
	return r
}