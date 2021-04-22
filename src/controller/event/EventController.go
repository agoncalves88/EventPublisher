package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"EventPublisher.Api/configuration"
	"EventPublisher.Api/infrastructure"
	"EventPublisher.Api/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var _eventCollection = "Event"

type EventController struct {
	Configuration configuration.Configuration
	MongoDb       *mongo.Database
	Kafka         *infrastructure.KafkaBase
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

		err := e.Kafka.Produce(ctx, message, request.EventName)
		var collection = e.MongoDb.Collection(_eventCollection)

		defer ctx.Done()
		if err == nil {
			eventModel := model.EventModel{
				ID:        primitive.NewObjectID(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Payload:   request.Payload,
				Completed: false,
			}

			collection.InsertOne(ctx, eventModel)
			c.Status(http.StatusCreated)
		} else {
			c.JSON(http.StatusUnprocessableEntity, EventResponse{
				Code:        "422",
				Description: err.Error(),
			})
		}

	}
}
