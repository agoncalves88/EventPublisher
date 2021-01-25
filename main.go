package main

import (
	"EventPublisher/docs"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	env := flag.String("env", "", "")
	flag.Parse()
	fmt.Printf("env: \"%v\"\n", string(*env))

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Event Publisher API"
	docs.SwaggerInfo.Description = "This is a concept api to publish events."
	docs.SwaggerInfo.Version = "0.1"
	docs.SwaggerInfo.Host = ""
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.New()

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":5555")
}
