package controller

import (
	"go-gin-api/entity"
	"go-gin-api/service"
	"go-gin-api/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// use context to acess the data comes via ...Http Reuest

// VideoController is ...
type VideoController interface {
	FindAll() []entity.Video
	Save(c *gin.Context) error
	ShowAll(c *gin.Context)
}

// create struct to implement the interface ...

type controller struct {
	service service.VideoService
}

// const fun recieves service return Videocontroller ...

// reference to vali
var validate *validator.Validate

// New is ...
func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-okey", validators.ValidateTitle)

	return &controller{
		service: service,
	}
}

func (cs *controller) FindAll() []entity.Video {
	return cs.service.FindAll()
}

func (cs *controller) Save(c *gin.Context) error {
	var video entity.Video

	// check error while binding ...
	err := c.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	cs.service.Save(video)
	return nil
}

func (cs *controller) ShowAll(c *gin.Context) {
	videos := cs.service.FindAll()
	data := gin.H{
		"titles": "video Page",
		"videos": videos,
	}
	c.HTML(http.StatusOK, "index.html", data)
}
