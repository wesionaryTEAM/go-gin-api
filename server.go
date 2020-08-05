package main

import (
	"go-gin-api/controller"
	"go-gin-api/middlewares"
	"go-gin-api/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

//var x int = z;

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

// Logging to a file.
// Use the following code if you need to write the logs to file and console at the same time.
func setUpLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	setUpLogOutput()
	r := gin.New()

	// server static files...
	r.Static("/css", "./templates/css")

	//load html templates...
	r.LoadHTMLGlob("templates/*.html")

	// Login Endpoint : Authentication + Token Creation ...
	r.POST("/login", func(c *gin.Context) {
		token := loginController.Login(c)
		if token != "" {
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			c.JSON(http.StatusUnauthorized, nil)
		}
	})

	r.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	// Group Api endpoints...
	apiGroupRoutes := r.Group("/api", middlewares.AuthorizeJWT())
	{

		apiGroupRoutes.GET("/videos", func(c *gin.Context) {
			c.JSON(200, videoController.FindAll())
		})

		apiGroupRoutes.POST("/videos", func(c *gin.Context) {
			err := videoController.Save(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "given input is valid",
				})
			}
		})
	}

	viewRoutes := r.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
