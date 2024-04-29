package services

import (
	"fmt"
	"rate-limiter/domain"
	"rate-limiter/errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var timestamp_test = time.Date(2024, time.April, 28, 17, 0, 0, 0, time.UTC)
var userIDTest = "userID_test"
var notificationTypeTest = "type_test"

func TestNotificationService_GetNotifications(t *testing.T) {
	testCases := []struct {
		name                         string
		expected                     []*domain.Notification
		expectedErr                  error
		mockedNotificationsContainer *NotificationsContainerMock
		userID                       string
	}{
		{
			name: "internal error",
			mockedNotificationsContainer: &NotificationsContainerMock{
				GetNotificationsByUserFunc: func(userID string) ([]*domain.Notification, error) {
					return nil, fmt.Errorf("internal error")
				},
			},
			expectedErr: fmt.Errorf("internal error"),
		},
		{
			name: "success",
			mockedNotificationsContainer: &NotificationsContainerMock{
				GetNotificationsByUserFunc: func(userID string) ([]*domain.Notification, error) {
					return []*domain.Notification{
						{
							Timestamp: timestamp_test,
							UserID:    userIDTest,
							Type:      notificationTypeTest,
						},
					}, nil
				},
			},
			expected: []*domain.Notification{
				{
					Timestamp: timestamp_test,
					UserID:    userIDTest,
					Type:      notificationTypeTest,
				},
			},
			userID: userIDTest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := &NotificationService{notificationsContainer: tc.mockedNotificationsContainer}
			notifications, err := service.GetNotificationsByUser(tc.userID)
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expected, notifications)
		})
	}
}

func TestNotificationService_SendNotification(t *testing.T) {
	testCases := []struct {
		name                         string
		expected                     map[string]*domain.RateLimitRule
		expectedErr                  error
		mockedNotificationsContainer *NotificationsContainerMock
		mockedRulesContainer         *RulesContainerMock
	}{
		{
			name: "internal error",
			mockedRulesContainer: &RulesContainerMock{
				GetRuleByTypeFunc: func(ns string) (*domain.RateLimitRule, error) {
					return nil, fmt.Errorf("error getting rate limit rule for notification type")
				},
			},
			expectedErr: fmt.Errorf("error getting rate limit rule for notification type"),
		},
		{
			name: "success without rule limit",
			mockedRulesContainer: &RulesContainerMock{
				GetRuleByTypeFunc: func(ns string) (*domain.RateLimitRule, error) {
					return nil, nil
				},
			},
		},
		{
			name: "error getting notification by user, type and interval",
			mockedRulesContainer: &RulesContainerMock{
				GetRuleByTypeFunc: func(ns string) (*domain.RateLimitRule, error) {
					return &domain.RateLimitRule{
						NotificationType: "example",
						MaxLimit:         10,
						TimeInterval:     domain.Duration{Duration: time.Second * 60},
					}, nil
				},
			},
			mockedNotificationsContainer: &NotificationsContainerMock{
				GetNotificationsByUserAndTypeAndIntervalFunc: func(params domain.GetNotificationParams) ([]*domain.Notification, error) {
					return nil, fmt.Errorf("some error")
				},
			},
			expectedErr: fmt.Errorf("some error"),
		},
		{
			name: "error - max exceeded",
			mockedRulesContainer: &RulesContainerMock{
				GetRuleByTypeFunc: func(ns string) (*domain.RateLimitRule, error) {
					return &domain.RateLimitRule{
						NotificationType: "example",
						MaxLimit:         2,
						TimeInterval:     domain.Duration{Duration: time.Second * 60},
					}, nil
				},
			},
			mockedNotificationsContainer: &NotificationsContainerMock{
				GetNotificationsByUserAndTypeAndIntervalFunc: func(params domain.GetNotificationParams) ([]*domain.Notification, error) {
					return []*domain.Notification{
						{
							Timestamp: timestamp_test,
							UserID:    userIDTest,
							Type:      notificationTypeTest,
						},
						{
							Timestamp: timestamp_test.Add(2 * time.Minute),
							UserID:    userIDTest,
							Type:      notificationTypeTest,
						},
					}, nil
				},
			},
			expectedErr: errors.ErrRateLimitExceeded,
		},
		{
			name: "error registering notification",
			mockedRulesContainer: &RulesContainerMock{
				GetRuleByTypeFunc: func(ns string) (*domain.RateLimitRule, error) {
					return &domain.RateLimitRule{
						NotificationType: "example",
						MaxLimit:         2,
						TimeInterval:     domain.Duration{Duration: time.Second * 60},
					}, nil
				},
			},
			mockedNotificationsContainer: &NotificationsContainerMock{
				GetNotificationsByUserAndTypeAndIntervalFunc: func(params domain.GetNotificationParams) ([]*domain.Notification, error) {
					return []*domain.Notification{
						{
							Timestamp: timestamp_test,
							UserID:    userIDTest,
							Type:      notificationTypeTest,
						},
					}, nil
				},
				AddNotificationFunc: func(userID string, notificationType string) error {
					return fmt.Errorf("some error")
				},
			},
			expectedErr: fmt.Errorf("some error"),
		},
		{
			name: "success",
			mockedRulesContainer: &RulesContainerMock{
				GetRuleByTypeFunc: func(ns string) (*domain.RateLimitRule, error) {
					return &domain.RateLimitRule{
						NotificationType: "example",
						MaxLimit:         2,
						TimeInterval:     domain.Duration{Duration: time.Second * 60},
					}, nil
				},
			},
			mockedNotificationsContainer: &NotificationsContainerMock{
				GetNotificationsByUserAndTypeAndIntervalFunc: func(params domain.GetNotificationParams) ([]*domain.Notification, error) {
					return []*domain.Notification{
						{
							Timestamp: timestamp_test,
							UserID:    userIDTest,
							Type:      notificationTypeTest,
						},
					}, nil
				},
				AddNotificationFunc: func(userID string, notificationType string) error {
					return nil
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := NewNotificationService(tc.mockedNotificationsContainer, tc.mockedRulesContainer)
			err := service.SendNotification(domain.SendNotificationParams{
				UserID:           userIDTest,
				NotificationType: notificationTypeTest,
			})
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
