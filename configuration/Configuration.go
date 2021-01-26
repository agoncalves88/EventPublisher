package configuration

import (
	"fmt"

	"EventPublisher/docs"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Port string
}

func GetConfig(environment string) Configuration {
	configuration := Configuration{}
	env := "local"
	if environment != "" {
		env = environment
	}
	fileName := fmt.Sprintf("./settings.%s.json", env)
	gonfig.GetConf(fileName, &configuration)
	return configuration
}

func SetSwaggerInfo() {
	docs.SwaggerInfo.Title = "Event Publisher API"
	docs.SwaggerInfo.Description = "This is a concept api to publish events."
	docs.SwaggerInfo.Version = "0.1"
	docs.SwaggerInfo.Host = ""
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
