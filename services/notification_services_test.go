package services

import (
	"fmt"
	"rate-limiter/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var timestamp_test = time.Date(2024, time.April, 28, 17, 0, 0, 0, time.UTC)

func TestNotificationService_GetRules(t *testing.T) {
	testCases := []struct {
		name            string
		expected        map[string]*domain.RateLimitRule
		expectedErr     error
		mockedContainer *ContainerMock
	}{
		{
			name: "internal error",
			mockedContainer: &ContainerMock{
				GetRulesFunc: func() (map[string]*domain.RateLimitRule, error) {
					return nil, fmt.Errorf("internal error")
				},
			},
			expectedErr: fmt.Errorf("internal error"),
		},
		{
			name: "success",
			mockedContainer: &ContainerMock{
				GetRulesFunc: func() (map[string]*domain.RateLimitRule, error) {
					return map[string]*domain.RateLimitRule{
						"rule1": {
							NotificationType: "type1",
							MaxLimit:         100,
							TimeInterval: domain.Duration{
								Duration: time.Minute,
							},
						},
						"rule2": {
							NotificationType: "type2",
							MaxLimit:         500,
							TimeInterval: domain.Duration{
								Duration: time.Hour,
							},
						},
					}, nil
				},
			},
			expected: map[string]*domain.RateLimitRule{
				"rule1": {
					NotificationType: "type1",
					MaxLimit:         100,
					TimeInterval: domain.Duration{
						Duration: time.Minute,
					},
				},
				"rule2": {
					NotificationType: "type2",
					MaxLimit:         500,
					TimeInterval: domain.Duration{
						Duration: time.Hour,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := &NotificationService{container: tc.mockedContainer}
			rules, err := service.GetRules()
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expected, rules)
		})
	}
}

func TestNotificationService_GetRulesByType(t *testing.T) {
	testCases := []struct {
		name            string
		expected        *domain.RateLimitRule
		expectedErr     error
		mockedContainer *ContainerMock
		ruleType        string
	}{
		{
			name: "internal error",
			mockedContainer: &ContainerMock{
				GetRuleByTypeFunc: func(rt string) (*domain.RateLimitRule, error) {
					return nil, fmt.Errorf("internal error")
				},
			},
			expectedErr: fmt.Errorf("internal error"),
		},
		{
			name: "success",
			mockedContainer: &ContainerMock{
				GetRuleByTypeFunc: func(rt string) (*domain.RateLimitRule, error) {
					return &domain.RateLimitRule{
						NotificationType: "type1",
						MaxLimit:         100,
						TimeInterval: domain.Duration{
							Duration: time.Minute,
						},
					}, nil
				},
			},
			expected: &domain.RateLimitRule{
				NotificationType: "type1",
				MaxLimit:         100,
				TimeInterval: domain.Duration{
					Duration: time.Minute,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := &NotificationService{container: tc.mockedContainer}
			rule, err := service.GetRuleByType(tc.ruleType)
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expected, rule)
		})
	}
}

func TestNotificationService_GetNotifications(t *testing.T) {
	testCases := []struct {
		name            string
		expected        map[string]map[string]*domain.Notification
		expectedErr     error
		mockedContainer *ContainerMock
	}{
		{
			name: "internal error",
			mockedContainer: &ContainerMock{
				GetNotificationsFunc: func() (map[string]map[string]*domain.Notification, error) {
					return nil, fmt.Errorf("internal error")
				},
			},
			expectedErr: fmt.Errorf("internal error"),
		},
		{
			name: "success",
			mockedContainer: &ContainerMock{
				GetNotificationsFunc: func() (map[string]map[string]*domain.Notification, error) {
					return map[string]map[string]*domain.Notification{
						"status": {
							"user_test_1": {
								Timestamp: timestamp_test,
								Count:     1,
							},
							"user_test_2": {
								Timestamp: timestamp_test,
								Count:     3,
							},
						},
						"news": {
							"user_test_3": {
								Timestamp: timestamp_test,
								Count:     1,
							},
						},
					}, nil
				},
			},
			expected: map[string]map[string]*domain.Notification{
				"status": {
					"user_test_1": {
						Timestamp: timestamp_test,
						Count:     1,
					},
					"user_test_2": {
						Timestamp: timestamp_test,
						Count:     3,
					},
				},
				"news": {
					"user_test_3": {
						Timestamp: timestamp_test,
						Count:     1,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := &NotificationService{container: tc.mockedContainer}
			notifications, err := service.GetNotifications()
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expected, notifications)
		})
	}
}