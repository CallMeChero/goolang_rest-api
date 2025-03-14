package routes

import (
	"example/rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse ID"})
		return
	}
	event, err := models.GetEvent(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	userId := context.GetInt64("userId")

	event.UserID = userId
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse ID"})
		return
	}

	event, err := models.GetEvent(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to find ID"})
		return
	}

	userId := context.GetInt64("userId")

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Only user that created event can edit it"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Updated event successfuly"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse ID"})
		return
	}

	event, err := models.GetEvent(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to find ID"})
		return
	}

	userId := context.GetInt64("userId")

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Only user that created event can delete it"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
