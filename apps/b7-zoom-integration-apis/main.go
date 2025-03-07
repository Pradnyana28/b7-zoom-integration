package main

import (
	"b7-zoom-integration/apps/b7-zoom-integration-apis/internal/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Hello(name string) string {
	result := "Hello " + name
	return result
}

func main() {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	engine.Use(cors.New(corsConfig))

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	controllers.AddUsersRoutes(engine)
	controllers.AddMeetingsRoutes(engine)

	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}
