package routes

import (
	v1 "tsweblist/api/v1"
	"tsweblist/middleware"
	"tsweblist/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.GINMODE)
	r := gin.Default()
	r.Use(middleware.Cors())

	router := r.Group("api/v1")
	{
		router.POST("addServerInfo", v1.AddServerInfo)
		router.POST("createChannel", v1.CreateChannel)
		router.GET("getServerChannel", v1.GetServerChannel)
		router.POST("changeChannelPassword", v1.ChangeChannelPassword)
	}
	r.Run(utils.GINPORT)
}
