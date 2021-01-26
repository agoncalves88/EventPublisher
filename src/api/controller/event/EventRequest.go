package controller

type EventRequest struct {
	CorrelationId string `json:"correlationId"`
	EventId       string `json:"eventId"`
	EventName     string `json:"eventName"`
	Payload       string `json:"payload"`
}
