package controller

import (
	"apigolang/entity"
	"apigolang/service"

	"github.com/gin-gonic/gin"
)

type VideoInterface interface {
	Index() []entity.Video
	Store(*gin.Context) entity.Video
}

type VideoController struct {
	service service.VideoService
}

func (c *VideoController) Index() []entity.Video {
	//return make([]entity.Video, 2)
	return c.service.FindAll()
}

func (c *VideoController) Store(ctx *gin.Context) entity.Video {
	var video entity.Video
	/*This way I try to bind the info comming from context request
	and set it in a variable called video*/
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
