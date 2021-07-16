package router

import (
	"apigolang/controller"

	"github.com/gin-gonic/gin"
)

// Define struct to declare video controller
type VideoRoutes struct {
	videoC controller.VideoController
}

func (videoR *VideoRoutes) InitializeRoutes(serverRoute *gin.Engine) {
	// Create group
	video := serverRoute.Group("/video")
	// Create routes
	{
		video.GET("/", func(c *gin.Context) { c.JSON(200, videoR.videoC.Index()) })
		video.POST("/", func(c *gin.Context) { c.JSON(200, videoR.videoC.Store(c)) })
	}
}
