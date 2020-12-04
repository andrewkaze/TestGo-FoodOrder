package Middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"foodorder/Models"
	"foodorder/Service"
	"strings"
)

func parseMap(aMap map[string]interface{}) {
	for key, val := range aMap {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println(key)
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			fmt.Println(key)
			parseArray(val.([]interface{}))
		default:
			fmt.Println(key, ":", concreteVal)

		}
	}
}
func parseArray(anArray []interface{}) {
	for i, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println("Index:", i)
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			fmt.Println("Index:", i)
			parseArray(val.([]interface{}))
		default:
			fmt.Println("Index", i, ":", concreteVal)

		}
	}
}


func AuthorizeUser() gin.HandlerFunc{
	return func(context *gin.Context) {
		//set when token from user validated
		var user,_ = context.Get("user")


		if user != ""{
			method := context.Request.URL.Path
			allowed := false
			var modelUser Models.User

			Models.GetUserByEmail(&modelUser, fmt.Sprintf("%v",user))

			for _, auth := range Service.StaticAuthService(){
				if modelUser.Role == auth.Role{
					for _, perm := range auth.Permission{
						if strings.Contains(method, perm){
							allowed = true
						}
					}
				}
			}

			if !allowed{
				context.AbortWithStatus(http.StatusUnauthorized)
			}
		}
		return
	}
}
