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

// func TestRateLimitService_SendNotification_Success_IntervalElapsed(t *testing.T) {
// 	// mockRulesContainer := &RulesContainerMock{
// 	// 	GetRuleByTypeFunc: func(s string) ([]*domain.RateLimitRule, error) {
// 	// 		return []*domain.RateLimitRule{
// 	// 			{
// 	// 				NotificationType: "news",
// 	// 				MaxLimit:         2,
// 	// 				TimeInterval:     domain.Duration{Duration: time.Second * 5},
// 	// 			},
// 	// 		}, nil
// 	// 	},
// 	// }

// 	rateLimitService := NewRateLimitService(
// 		notifications.NewInMemoryNotificationsContainer(),
// 		NewRulesService(rules.NewInMemoryRulesContainer()),
// 	)

// 	params := domain.SendNotificationParams{
// 		UserID:           "user1",
// 		NotificationType: "email",
// 	}

// 	rateLimitService.SendNotification(params)
// 	notifications, _ := rateLimitService.notificationsContainer.GetNotificationsByUser(domain.GetNotificationParams{
// 		UserID:           params.UserID,
// 		NotificationType: params.NotificationType,
// 		TimeInterval:     time.Hour * 100,
// 	})
// 	fmt.Println("Notifications from 1st send", utils.SerializeObject(notifications))
// 	rateLimitService.SendNotification(params)
// 	notifications, _ = rateLimitService.notificationsContainer.GetNotificationsByUser(domain.GetNotificationParams{
// 		UserID:           params.UserID,
// 		NotificationType: params.NotificationType,
// 		TimeInterval:     time.Hour * 100,
// 	})
// 	fmt.Println("Notifications from 2nd send", utils.SerializeObject(notifications))
// 	err := rateLimitService.SendNotification(params)
// 	notifications, _ = rateLimitService.notificationsContainer.GetNotificationsByUser(domain.GetNotificationParams{
// 		UserID:           params.UserID,
// 		NotificationType: params.NotificationType,
// 		TimeInterval:     time.Hour * 100,
// 	})
// 	fmt.Println("Notifications from 3rd send", utils.SerializeObject(notifications))

// 	//err := ExcecuteFuncNtimes(rateLimitService.SendNotification, params, 3)
// 	fmt.Println("err frm send", utils.SerializeObject(err))
// 	assert.Equal(t, errors.ErrRateLimitExceeded, err)

// 	fmt.Println("Waiting  10 seconds for test ...")
// 	time.Sleep(time.Second * 10)

// 	err = rateLimitService.SendNotification(domain.SendNotificationParams{
// 		UserID:           "user1",
// 		NotificationType: "email",
// 	})
// 	assert.Equal(t, nil, err)
// }

// func ExcecuteFuncNtimes(fn func(domain.SendNotificationParams) error, params domain.SendNotificationParams, times int) error {
// 	var err error
// 	for i := 0; i < times; i++ {
// 		err = fn(params)
// 	}
// 	return err
// }
