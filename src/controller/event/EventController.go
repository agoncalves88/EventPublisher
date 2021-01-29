package controller

import (
	"context"
	"encoding/json"
	"net/http"

	kafka "EventPublisher.Api/Infrastructure"
	"EventPublisher.Api/configuration"
	"github.com/gin-gonic/gin"
)

// Controller example
type EventController struct {
	Configuration configuration.Configuration
}

func (e EventController) PostEvent(c *gin.Context) {

	request := EventRequest{}
	if c.Bind(&request) != nil {
		c.JSON(http.StatusBadRequest, EventResponse{
			Code:        "400",
			Description: "Invalid object data",
		})
	} else {

		ctx := context.Background()
		message, _ := json.Marshal(request)
		go kafka.Produce(ctx, message)

		c.JSON(http.StatusOK, EventResponse{
			Code:        "200",
			Description: "",
		})
	}
}
