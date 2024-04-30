package controllers

import (
	"net/http"
	"rate-limiter/domain"
	"rate-limiter/errors"
	"strings"

	"github.com/gin-gonic/gin"
)

type RateLimitService interface {
	SendNotification(domain.SendNotificationParams) error
}
type NotificationController struct {
	RateLimitService RateLimitService
}

func (nc NotificationController) Pong(c *gin.Context) {
	c.Set("skip", true)
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Pong from Notifications"})
}

func (nc NotificationController) SendNotification(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, &errors.ApiError{Message: "userID is mandatory", ErrorStr: "invalid_user_id", Status: http.StatusBadRequest})
		return
	}

	notificationType := c.GetString("type")
	if notificationType == "" {
		c.JSON(http.StatusBadRequest, &errors.ApiError{Message: "notification type is mandatory", ErrorStr: "invalid_type", Status: http.StatusBadRequest})
		return
	}

	err := nc.RateLimitService.SendNotification(domain.SendNotificationParams{
		UserID:           userID,
		NotificationType: notificationType,
	})
	if err != nil {
		if errors.IsTooManyRequestsError(err) {
			c.JSON(http.StatusTooManyRequests, &errors.ApiError{Message: "message limit exceeded", ErrorStr: err.Error(), Status: http.StatusTooManyRequests})
		} else {
			c.JSON(http.StatusInternalServerError, &errors.ApiError{Message: "internal server error", ErrorStr: err.Error(), Status: http.StatusInternalServerError})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "notification sent"})
	}
}

func (nc NotificationController) ValidateNotificationType(c *gin.Context) error {
	notificationType := c.Param("type")
	if notificationType == "" {
		return &errors.ApiError{Message: "notification type is mandatory", ErrorStr: "invalid_type", Status: http.StatusBadRequest}
	}
	c.Set("type", strings.ToLower(notificationType))
	return nil
}

func (nc NotificationController) ValidateUserID(c *gin.Context) error {
	userID := c.Param("user_id")
	if userID == "" {
		return &errors.ApiError{Message: "userID is mandatory", ErrorStr: "invalid_user_id", Status: http.StatusBadRequest}
	}
	c.Set("userID", userID)
	return nil
}
