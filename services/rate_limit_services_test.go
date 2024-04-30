package services

import (
	"fmt"
	"rate-limiter/domain"
	"rate-limiter/errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var userIDTest = "userID_test"
var notificationTypeTest = "type_test"

func TestRateLimitService_SendNotification_ErrorGetRules(t *testing.T) {
	mockRulesContainer := &RulesContainerMock{
		GetRuleByTypeFunc: func(s string) ([]*domain.RateLimitRule, error) {
			return nil, fmt.Errorf("some error")
		},
	}

	rateLimitService := NewRateLimitService(&NotificationsContainerMock{}, NewRulesService(mockRulesContainer))
	err := rateLimitService.SendNotification(domain.SendNotificationParams{
		UserID:           "user1",
		NotificationType: "email",
	})

	assert.Equal(t, errors.ErrGetRateLimitRule, err)
}

func TestRateLimitService_SendNotification_Success_RuleNotExists(t *testing.T) {
	mockNotificationsContainer := &NotificationsContainerMock{
		AddNotificationFunc: func(userID string, notificationType string) error {
			return nil
		},
		GetNotificationsByUserFunc: func(params domain.GetNotificationParams) ([]*domain.Notification, error) {
			return []*domain.Notification{}, nil
		},
	}
	mockRulesContainer := &RulesContainerMock{
		GetRuleByTypeFunc: func(s string) ([]*domain.RateLimitRule, error) {
			return nil, nil
		},
	}

	rateLimitService := NewRateLimitService(mockNotificationsContainer, NewRulesService(mockRulesContainer))
	err := rateLimitService.SendNotification(domain.SendNotificationParams{
		UserID:           "user1",
		NotificationType: "email",
	})

	assert.NoError(t, err)
}

func TestRateLimitService_SendNotification_ErrorGetNotifications(t *testing.T) {
	mockNotificationsContainer := &NotificationsContainerMock{
		GetNotificationsByUserFunc: func(params domain.GetNotificationParams) ([]*domain.Notification, error) {
			return nil, fmt.Errorf("error getting rate limit rule for notification type")
		},
	}
	mockRulesContainer := &RulesContainerMock{
		GetRuleByTypeFunc: func(s string) ([]*domain.RateLimitRule, error) {
			return []*domain.RateLimitRule{
				{
					NotificationType: "news",
					MaxLimit:         3,
					TimeInterval:     domain.Duration{Duration: time.Second * 60},
				},
			}, nil
		},
	}

	rateLimitService := NewRateLimitService(mockNotificationsContainer, NewRulesService(mockRulesContainer))
	err := rateLimitService.SendNotification(domain.SendNotificationParams{
		UserID:           "user1",
		NotificationType: "email",
	})

	expectedError := fmt.Errorf("error getting rate limit rule for notification type")
	assert.Equal(t, expectedError, err)
}

func TestRateLimitService_SendNotification_LimitExceeded(t *testing.T) {
	mockNotificationsContainer := &NotificationsContainerMock{
		GetNotificationsByUserFunc: func(params domain.GetNotificationParams) ([]*domain.Notification, error) {
			return []*domain.Notification{
				{
					Timestamp: time.Now().Add(-time.Second * 10),
					UserID:    userIDTest,
					Type:      notificationTypeTest,
				},
				{
					Timestamp: time.Now().Add(-time.Second * 5),
					UserID:    userIDTest,
					Type:      notificationTypeTest,
				},
			}, nil
		},
	}
	mockRulesContainer := &RulesContainerMock{
		GetRuleByTypeFunc: func(s string) ([]*domain.RateLimitRule, error) {
			return []*domain.RateLimitRule{
				{
					NotificationType: "news",
					MaxLimit:         2,
					TimeInterval:     domain.Duration{Duration: time.Second * 60},
				},
			}, nil
		},
	}
	rateLimitService := NewRateLimitService(mockNotificationsContainer, NewRulesService(mockRulesContainer))
	err := rateLimitService.SendNotification(domain.SendNotificationParams{
		UserID:           "user1",
		NotificationType: "email",
	})

	assert.Equal(t, errors.ErrRateLimitExceeded, err)
}

func TestRateLimitService_SendNotification_ErrorAddNotification(t *testing.T) {
	mockNotificationsContainer := &NotificationsContainerMock{
		AddNotificationFunc: func(userID string, notificationType string) error {
			return fmt.Errorf("some error")
		},
		GetNotificationsByUserFunc: func(params domain.GetNotificationParams) ([]*domain.Notification, error) {
			return []*domain.Notification{
				{
					Timestamp: time.Now().Add(-time.Second * 10),
					UserID:    userIDTest,
					Type:      notificationTypeTest,
				},
				{
					Timestamp: time.Now().Add(-time.Second * 5),
					UserID:    userIDTest,
					Type:      notificationTypeTest,
				},
			}, nil
		},
	}
	mockRulesContainer := &RulesContainerMock{
		GetRuleByTypeFunc: func(s string) ([]*domain.RateLimitRule, error) {
			return []*domain.RateLimitRule{
				{
					NotificationType: "news",
					MaxLimit:         3,
					TimeInterval:     domain.Duration{Duration: time.Second * 60},
				},
			}, nil
		},
	}
	rateLimitService := NewRateLimitService(mockNotificationsContainer, NewRulesService(mockRulesContainer))
	err := rateLimitService.SendNotification(domain.SendNotificationParams{
		UserID:           "user1",
		NotificationType: "email",
	})

	assert.NoError(t, err)
}

func TestRateLimitService_SendNotification_Success_WithinInterval_LimitNotExceeded(t *testing.T) {
	mockNotificationsContainer := &NotificationsContainerMock{
		AddNotificationFunc: func(userID string, notificationType string) error {
			return nil
		},
		GetNotificationsByUserFunc: func(params domain.GetNotificationParams) ([]*domain.Notification, error) {
			return []*domain.Notification{
				{
					Timestamp: time.Now().Add(-time.Second * 10),
					UserID:    userIDTest,
					Type:      notificationTypeTest,
				},
				{
					Timestamp: time.Now().Add(-time.Second * 5),
					UserID:    userIDTest,
					Type:      notificationTypeTest,
				},
			}, nil
		},
	}
	mockRulesContainer := &RulesContainerMock{
		GetRuleByTypeFunc: func(s string) ([]*domain.RateLimitRule, error) {
			return []*domain.RateLimitRule{
				{
					NotificationType: "news",
					MaxLimit:         3,
					TimeInterval:     domain.Duration{Duration: time.Second * 60},
				},
			}, nil
		},
	}

	rateLimitService := NewRateLimitService(mockNotificationsContainer, NewRulesService(mockRulesContainer))
	err := rateLimitService.SendNotification(domain.SendNotificationParams{
		UserID:           "user1",
		NotificationType: "email",
	})

	assert.NoError(t, err)
}
