package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ErrorSentry(c *gin.Context){
	j,_ := strconv.ParseInt(c.Param("j"), 10, 64)
	i := 10/j
	fmt.Println(i)
}
