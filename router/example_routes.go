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
type ExampleRoutes struct {
	exampleController controller.ExampleController
}

func (route *ExampleRoutes) InitializeRoutes(serverRoute *gin.Engine) {
	// Create group
	video := serverRoute.Group("/video")
	// Create routes
	{
		//video.GET("/", func(c *gin.Context) { c.JSON(200, videoR.videoC.Index()) })
		video.GET("/", route.exampleController.Index())
		video.POST("/", route.exampleController.Store())
	}
}

func (route *ExampleRoutes) Helloworld() {
	fmt.Println("Hola Mundo")
}
