package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"open-api-backend/config"
)

var r *gin.Engine

var apiGroup *gin.RouterGroup

func init() {
	//
	r = gin.Default()

	// cors

	// routers definition
	apiGroup = r.Group("api")
	addUserRouter()
	//addInterfaceRouter()
}

func StartGinServer() {
	err := r.Run(config.Conf.Gin.GetAddr())
	if err != nil {
		fmt.Println("[Error] r.Run " + err.Error())
		return
	}
}
