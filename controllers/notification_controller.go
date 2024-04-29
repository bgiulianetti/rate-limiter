package controllers

import (
	"fmt"
	"net/http"
	"rate-limiter/errors"
	"rate-limiter/services"
	"strings"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	NotificationService services.Service
}

func (nc NotificationController) Pong(c *gin.Context) {
	c.Set("skip", true)
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Pong from Notifications"})
}

func (nc NotificationController) GetRules(c *gin.Context) {
	rules, err := nc.NotificationService.GetRules()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &errors.ApiError{Message: "error getting rules", ErrorStr: err.Error(), Status: http.StatusInternalServerError})
	}
	c.JSON(http.StatusOK, rules)
}

func (nc NotificationController) GetRuleByType(c *gin.Context) {
	notificationType := c.GetString("type")
	if notificationType == "" {
		c.JSON(http.StatusBadRequest, &errors.ApiError{Message: "notification type is mandatory", ErrorStr: "invalid_rule", Status: http.StatusBadRequest})
	}
	rule, err := nc.NotificationService.GetRuleByType(notificationType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &errors.ApiError{Message: "error getting rule by type", ErrorStr: err.Error(), Status: http.StatusInternalServerError})
	}
	if rule == nil {
		c.JSON(http.StatusNotFound, &errors.ApiError{Message: "notification type not found", ErrorStr: "rule_not_found", Status: http.StatusNotFound})
	}
	c.JSON(http.StatusOK, rule)
}

func (nc NotificationController) GetNotifications(c *gin.Context) {
	notifications, err := nc.NotificationService.GetNotifications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &errors.ApiError{Message: "error getting rule by type", ErrorStr: err.Error(), Status: http.StatusInternalServerError})
	}
	c.JSON(http.StatusOK, notifications)
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

	err := nc.NotificationService.SendNotification(userID, notificationType)
	fmt.Println("error in controller:", err)
	if err != nil {
		fmt.Println("enter err != nil")
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
