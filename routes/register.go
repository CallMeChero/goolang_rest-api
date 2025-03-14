package routes

import (
	"example/rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
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

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event registered successfully"})
}

func cancelRegitration(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse ID"})
		return
	}

	userId := context.GetInt64("userId")

	event, err := models.GetRegisteredEvent(userId, eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to find ID"})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Only user that created registered event can delete it"})
		return
	}

	err = event.DeleteRegisteredEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete registered event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
