package router

import (
	"github.com/gin-gonic/gin"
)

type PruebaRoute struct {
	nombre string
}

func (route *PruebaRoute) InitializeRoutes(serverRoute *gin.Engine) {
	// Create group
	video := serverRoute.Group("/prueba")
	// Create routes
	{
		//video.GET("/", func(c *gin.Context) { c.JSON(200, videoR.videoC.Index()) })
		video.GET("/", func(c *gin.Context) {})
	}
}

func (p *PruebaRoute) visualizarNombre() string {
	return p.nombre
}
