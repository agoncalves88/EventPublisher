package main

import (
	"EventPublisher/configuration"
	controller "EventPublisher/controller/event"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	env := flag.String("env", "", "")
	flag.Parse()

	configuration.SetSwaggerInfo()
	config := configuration.GetConfig(string(*env))

	r := gin.Default()
	//r.Use(gzip.Gzip(gzip.BestSpeed))

	v1 := r.Group("/v1")
	{
		v1.POST("/event", controller.EventController{Configuration: config}.PostEvent)
	}

	fmt.Println(config.Port)
	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Run(":" + config.Port)
}

//
