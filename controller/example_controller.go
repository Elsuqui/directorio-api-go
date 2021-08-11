package controller

import (
	"github.com/gin-gonic/gin"
)

type ExampleInterface interface {
	Index() gin.HandlerFunc
	Store() gin.HandlerFunc
}

type ExampleController struct {
	//service service.VideoService
}

func (c *ExampleController) Index() gin.HandlerFunc {
	//return make([]entity.Video, 2)
	//videos := c.service.FindAll()
	// Return handler function from controller
	return func(c *gin.Context) {
		c.JSON(200, "")
	}
}

func (c *ExampleController) Store() gin.HandlerFunc {
	//var video entity.Video
	/*This way I try to bind the info comming from context request
	and set it in a variable called video*/
	// Return handler function from controller
	return func(ctx *gin.Context) {
		//ctx.BindJSON(&video)
		//c.service.Save(video)
		ctx.JSON(200, "")
	}
}
