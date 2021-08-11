package router

import (
	"apigolang/controller"
	"fmt"

	"github.com/gin-gonic/gin"
)

/*type VideoRoutesI interface {
	InitializeRoutes()
	Helloworld()
}*/

// Define struct to declare video controller
type VideoRoutes struct {
	videoController controller.VideoController
}

func (route *VideoRoutes) InitializeRoutes(serverRoute *gin.Engine) {
	// Create group
	video := serverRoute.Group("/video")
	// Create routes
	{
		//video.GET("/", func(c *gin.Context) { c.JSON(200, videoR.videoC.Index()) })
		video.GET("/", route.videoController.Index())
		video.POST("/", route.videoController.Store())
	}
}

func (route *VideoRoutes) Helloworld() {
	fmt.Println("Hola Mundo")
}
