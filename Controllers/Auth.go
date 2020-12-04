package Controllers

import (
	"github.com/gin-gonic/gin"
	"foodorder/Dto"
	"foodorder/Models"
	"foodorder/Service"
)

type LoginController interface {
	Login(c *gin.Context) string
}

type loginController struct {
	loginService Service.LoginService
	jwtService Service.JWTService
}

func LoginHandler(loginService Service.LoginService, jwtservice Service.JWTService)  LoginController{
	return &loginController{
		loginService: loginService,
		jwtService:   jwtservice,
	}
}

func (controller *loginController)Login(c *gin.Context)  string{
	var credential Dto.LoginCredentials
	err := c.ShouldBind(&credential)
	var user Models.User
	if err != nil{
		return "No Data Found"
	}

	query := Models.GetUserByEmailPassword(&user, credential.Email,credential.Password)
	if query == nil{
		return controller.jwtService.GenerateToken(credential.Email,user.Role, true)
	}else{
		return ""
	}

	//isUserAuthenticated := controller.loginService.LoginUser(credential.Email, credential.Password)
	//if isUserAuthenticated{

	//}


}