package controller

import (
	"github.com/QiaoLin22/Golang-Gin/entity"
	"github.com/QiaoLin22/Golang-Gin/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type videoController struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &videoController{
		service: service,
	}
}

func (c *videoController) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}

func (c *videoController) FindAll() []entity.Video {
	return c.service.FindAll()
}
