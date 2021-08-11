package controller

import (
	"apigolang/entity"
	"apigolang/service"

	"github.com/gin-gonic/gin"
)

type VideoInterface interface {
	Index() gin.HandlerFunc
	Store() gin.HandlerFunc
}

type VideoController struct {
	service service.VideoService
}

func (c *VideoController) Index() gin.HandlerFunc {
	//return make([]entity.Video, 2)
	videos := c.service.FindAll()
	// Return handler function from controller
	return func(c *gin.Context) {
		c.JSON(200, videos)
	}
}

func (c *VideoController) Store() gin.HandlerFunc {
	var video entity.Video
	/*This way I try to bind the info comming from context request
	and set it in a variable called video*/
	// Return handler function from controller
	return func(ctx *gin.Context) {
		ctx.BindJSON(&video)
		c.service.Save(video)
		ctx.JSON(200, video)
	}
}
