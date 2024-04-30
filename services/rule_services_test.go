package services

import (
	"fmt"
	"rate-limiter/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNotificationService_GetRules(t *testing.T) {
	testCases := []struct {
		name                 string
		expected             map[string][]*domain.RateLimitRule
		expectedErr          error
		mockedRulesContainer *RulesContainerMock
	}{
		{
			name: "internal error",
			mockedRulesContainer: &RulesContainerMock{
				GetRulesFunc: func() (map[string][]*domain.RateLimitRule, error) {
					return nil, fmt.Errorf("internal error")
				},
			},
			expectedErr: fmt.Errorf("internal error"),
		},
		{
			name: "success",
			mockedRulesContainer: &RulesContainerMock{
				GetRulesFunc: func() (map[string][]*domain.RateLimitRule, error) {
					return map[string][]*domain.RateLimitRule{
						"rule1": {
							{
								NotificationType: "type1",
								MaxLimit:         100,
								TimeInterval: domain.Duration{
									Duration: time.Minute,
								},
							},
						},
						"rule2": {
							{
								NotificationType: "type2",
								MaxLimit:         500,
								TimeInterval: domain.Duration{
									Duration: time.Hour,
								},
							},
						},
					}, nil
				},
			},
			expected: map[string][]*domain.RateLimitRule{
				"rule1": {
					{
						NotificationType: "type1",
						MaxLimit:         100,
						TimeInterval: domain.Duration{
							Duration: time.Minute,
						},
					},
				},
				"rule2": {
					{
						NotificationType: "type2",
						MaxLimit:         500,
						TimeInterval: domain.Duration{
							Duration: time.Hour,
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := &RulesService{rulesContainer: tc.mockedRulesContainer}
			rules, err := service.GetRules()
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expected, rules)
		})
	}
}

func TestNotificationService_GetRulesByType(t *testing.T) {
	testCases := []struct {
		name                 string
		expected             []*domain.RateLimitRule
		expectedErr          error
		mockedRulesContainer *RulesContainerMock
		ruleType             string
	}{
		{
			name: "internal error",
			mockedRulesContainer: &RulesContainerMock{
				GetRuleByTypeFunc: func(rt string) ([]*domain.RateLimitRule, error) {
					return nil, fmt.Errorf("internal error")
				},
			},
			expectedErr: fmt.Errorf("internal error"),
		},
		{
			name: "success",
			mockedRulesContainer: &RulesContainerMock{
				GetRuleByTypeFunc: func(rt string) ([]*domain.RateLimitRule, error) {
					return []*domain.RateLimitRule{
						{
							NotificationType: "type1",
							MaxLimit:         100,
							TimeInterval: domain.Duration{
								Duration: time.Minute,
							},
						},
					}, nil
				},
			},
			expected: []*domain.RateLimitRule{
				{
					NotificationType: "type1",
					MaxLimit:         100,
					TimeInterval: domain.Duration{
						Duration: time.Minute,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := &RulesService{rulesContainer: tc.mockedRulesContainer}
			rule, err := service.GetRuleByType(tc.ruleType)
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expected, rule)
		})
	}
}
