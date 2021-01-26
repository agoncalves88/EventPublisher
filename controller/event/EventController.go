package controller

import (
	"EventPublisher/configuration"
	"net/http"

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
		c.JSON(http.StatusOK, EventResponse{
			Code:        "200",
			Description: "",
		})
	}
}
