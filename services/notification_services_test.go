package services

// var timestamp_test = time.Date(2024, time.April, 28, 17, 0, 0, 0, time.UTC)

// func TestNotificationService_GetRules(t *testing.T) {
// 	testCases := []struct {
// 		name            string
// 		expected        map[string]*domain.RateLimitRule
// 		expectedErr     error
// 		mockedContainer *ContainerMock
// 	}{
// 		{
// 			name: "internal error",
// 			mockedContainer: &ContainerMock{
// 				GetRulesFunc: func() (map[string]*domain.RateLimitRule, error) {
// 					return nil, fmt.Errorf("internal error")
// 				},
// 			},
// 			expectedErr: fmt.Errorf("internal error"),
// 		},
// 		{
// 			name: "success",
// 			mockedContainer: &ContainerMock{
// 				GetRulesFunc: func() (map[string]*domain.RateLimitRule, error) {
// 					return map[string]*domain.RateLimitRule{
// 						"rule1": {
// 							NotificationType: "type1",
// 							MaxLimit:         100,
// 							TimeInterval: domain.Duration{
// 								Duration: time.Minute,
// 							},
// 						},
// 						"rule2": {
// 							NotificationType: "type2",
// 							MaxLimit:         500,
// 							TimeInterval: domain.Duration{
// 								Duration: time.Hour,
// 							},
// 						},
// 					}, nil
// 				},
// 			},
// 			expected: map[string]*domain.RateLimitRule{
// 				"rule1": {
// 					NotificationType: "type1",
// 					MaxLimit:         100,
// 					TimeInterval: domain.Duration{
// 						Duration: time.Minute,
// 					},
// 				},
// 				"rule2": {
// 					NotificationType: "type2",
// 					MaxLimit:         500,
// 					TimeInterval: domain.Duration{
// 						Duration: time.Hour,
// 					},
// 				},
// 			},
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			service := &NotificationService{container: tc.mockedContainer}
// 			rules, err := service.GetRules()
// 			assert.Equal(t, tc.expectedErr, err)
// 			assert.Equal(t, tc.expected, rules)
// 		})
// 	}
// }

// func TestNotificationService_GetRulesByType(t *testing.T) {
// 	testCases := []struct {
// 		name            string
// 		expected        *domain.RateLimitRule
// 		expectedErr     error
// 		mockedContainer *ContainerMock
// 		ruleType        string
// 	}{
// 		{
// 			name: "internal error",
// 			mockedContainer: &ContainerMock{
// 				GetRuleByTypeFunc: func(rt string) (*domain.RateLimitRule, error) {
// 					return nil, fmt.Errorf("internal error")
// 				},
// 			},
// 			expectedErr: fmt.Errorf("internal error"),
// 		},
// 		{
// 			name: "success",
// 			mockedContainer: &ContainerMock{
// 				GetRuleByTypeFunc: func(rt string) (*domain.RateLimitRule, error) {
// 					return &domain.RateLimitRule{
// 						NotificationType: "type1",
// 						MaxLimit:         100,
// 						TimeInterval: domain.Duration{
// 							Duration: time.Minute,
// 						},
// 					}, nil
// 				},
// 			},
// 			expected: &domain.RateLimitRule{
// 				NotificationType: "type1",
// 				MaxLimit:         100,
// 				TimeInterval: domain.Duration{
// 					Duration: time.Minute,
// 				},
// 			},
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			service := &NotificationService{container: tc.mockedContainer}
// 			rule, err := service.GetRuleByType(tc.ruleType)
// 			assert.Equal(t, tc.expectedErr, err)
// 			assert.Equal(t, tc.expected, rule)
// 		})
// 	}
// }

// func TestNotificationService_GetNotifications(t *testing.T) {
// 	testCases := []struct {
// 		name            string
// 		expected        map[string]map[string]*domain.Notification
// 		expectedErr     error
// 		mockedContainer *ContainerMock
// 	}{
// 		{
// 			name: "internal error",
// 			mockedContainer: &ContainerMock{
// 				GetNotificationsFunc: func() (map[string]map[string]*domain.Notification, error) {
// 					return nil, fmt.Errorf("internal error")
// 				},
// 			},
// 			expectedErr: fmt.Errorf("internal error"),
// 		},
// 		{
// 			name: "success",
// 			mockedContainer: &ContainerMock{
// 				GetNotificationsFunc: func() (map[string]map[string]*domain.Notification, error) {
// 					return map[string]map[string]*domain.Notification{
// 						"status": {
// 							"user_test_1": {
// 								Timestamp: timestamp_test,
// 								Count:     1,
// 							},
// 							"user_test_2": {
// 								Timestamp: timestamp_test,
// 								Count:     3,
// 							},
// 						},
// 						"news": {
// 							"user_test_3": {
// 								Timestamp: timestamp_test,
// 								Count:     1,
// 							},
// 						},
// 					}, nil
// 				},
// 			},
// 			expected: map[string]map[string]*domain.Notification{
// 				"status": {
// 					"user_test_1": {
// 						Timestamp: timestamp_test,
// 						Count:     1,
// 					},
// 					"user_test_2": {
// 						Timestamp: timestamp_test,
// 						Count:     3,
// 					},
// 				},
// 				"news": {
// 					"user_test_3": {
// 						Timestamp: timestamp_test,
// 						Count:     1,
// 					},
// 				},
// 			},
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			service := &NotificationService{container: tc.mockedContainer}
// 			notifications, err := service.GetNotifications()
// 			assert.Equal(t, tc.expectedErr, err)
// 			assert.Equal(t, tc.expected, notifications)
// 		})
// 	}
// }

// func TestNotificationService_SendNotification(t *testing.T) {
// 	testCases := []struct {
// 		name            string
// 		expected        map[string]*domain.RateLimitRule
// 		expectedErr     error
// 		mockedContainer *ContainerMock
// 	}{
// 		{
// 			name: "internal error",
// 			mockedContainer: &ContainerMock{
// 				GetRuleByTypeFunc: func(ns string) (*domain.RateLimitRule, error) {
// 					return nil, fmt.Errorf("error getting rate limit rule for notification type")
// 				},
// 			},
// 			expectedErr: fmt.Errorf("error getting rate limit rule for notification type"),
// 		},
// 		{
// 			name: "success without rule limit",
// 			mockedContainer: &ContainerMock{
// 				GetRuleByTypeFunc: func(ns string) (*domain.RateLimitRule, error) {
// 					return nil, nil
// 				},
// 			},
// 		},
// 		{
// 			name: "error getting notification by type and user",
// 			mockedContainer: &ContainerMock{
// 				GetRuleByTypeFunc: func(ns string) (*domain.RateLimitRule, error) {
// 					return &domain.RateLimitRule{
// 						NotificationType: "example",
// 						MaxLimit:         10,
// 						TimeInterval:     domain.Duration{Duration: time.Second * 60},
// 					}, nil
// 				},
// 				GetNotificationByTypeAndUserFunc: func(s1 string, s2 string) (*domain.Notification, error) {
// 					return nil, fmt.Errorf("some error")
// 				},
// 			},
// 			expectedErr: fmt.Errorf("some error"),
// 		},
// 		{
// 			name: "error with first notification",
// 			mockedContainer: &ContainerMock{
// 				GetRuleByTypeFunc: func(ns string) (*domain.RateLimitRule, error) {
// 					return &domain.RateLimitRule{
// 						NotificationType: "example",
// 						MaxLimit:         10,
// 						TimeInterval:     domain.Duration{Duration: time.Second * 60},
// 					}, nil
// 				},
// 				GetNotificationByTypeAndUserFunc: func(s1 string, s2 string) (*domain.Notification, error) {
// 					return nil, nil
// 				},
// 				IncrementNotificationCountFunc: func(s1 string, s2 string) error {
// 					return fmt.Errorf("some error")
// 				},
// 			},
// 			expectedErr: fmt.Errorf("some error"),
// 		},
// 		{
// 			name: "success first notification",
// 			mockedContainer: &ContainerMock{
// 				GetRuleByTypeFunc: func(ns string) (*domain.RateLimitRule, error) {
// 					return &domain.RateLimitRule{
// 						NotificationType: "example",
// 						MaxLimit:         10,
// 						TimeInterval:     domain.Duration{Duration: time.Second * 60},
// 					}, nil
// 				},
// 				GetNotificationByTypeAndUserFunc: func(s1 string, s2 string) (*domain.Notification, error) {
// 					return nil, nil
// 				},
// 				IncrementNotificationCountFunc: func(s1 string, s2 string) error {
// 					return nil
// 				},
// 			},
// 		},
// 		{
// 			name: "error restarting counter - interval elapsed",
// 			mockedContainer: &ContainerMock{
// 				GetRuleByTypeFunc: func(ns string) (*domain.RateLimitRule, error) {
// 					return &domain.RateLimitRule{
// 						NotificationType: "example",
// 						MaxLimit:         3,
// 						TimeInterval:     domain.Duration{Duration: time.Second * 1},
// 					}, nil
// 				},
// 				GetNotificationByTypeAndUserFunc: func(s1 string, s2 string) (*domain.Notification, error) {
// 					return &domain.Notification{
// 						Timestamp: timestamp_test,
// 						Count:     3,
// 					}, nil
// 				},
// 				ResetNotificationCountFunc: func(s1 string, s2 string) error {
// 					return fmt.Errorf("some error")
// 				},
// 			},
// 			expectedErr: fmt.Errorf("some error"),
// 		},
// 		{
// 			name: "success restarting counter - interval elapsed",
// 			mockedContainer: &ContainerMock{
// 				GetRuleByTypeFunc: func(ns string) (*domain.RateLimitRule, error) {
// 					return &domain.RateLimitRule{
// 						NotificationType: "example",
// 						MaxLimit:         3,
// 						TimeInterval:     domain.Duration{Duration: time.Second * 1},
// 					}, nil
// 				},
// 				GetNotificationByTypeAndUserFunc: func(s1 string, s2 string) (*domain.Notification, error) {
// 					return &domain.Notification{
// 						Timestamp: timestamp_test,
// 						Count:     3,
// 					}, nil
// 				},
// 				ResetNotificationCountFunc: func(s1 string, s2 string) error {
// 					return nil
// 				},
// 			},
// 		},
// 		{
// 			name: "Within interval - max exceeded",
// 			mockedContainer: &ContainerMock{
// 				GetRuleByTypeFunc: func(ns string) (*domain.RateLimitRule, error) {
// 					return &domain.RateLimitRule{
// 						NotificationType: "example",
// 						MaxLimit:         3,
// 						TimeInterval:     domain.Duration{Duration: time.Hour * 24},
// 					}, nil
// 				},
// 				GetNotificationByTypeAndUserFunc: func(s1 string, s2 string) (*domain.Notification, error) {
// 					return &domain.Notification{
// 						Timestamp: timestamp_test,
// 						Count:     3,
// 					}, nil
// 				},
// 			},
// 			expectedErr: errors.ErrRateLimitExceeded,
// 		},
// 		{
// 			name: "Error within interval - increment counter",
// 			mockedContainer: &ContainerMock{
// 				GetRuleByTypeFunc: func(ns string) (*domain.RateLimitRule, error) {
// 					return &domain.RateLimitRule{
// 						NotificationType: "example",
// 						MaxLimit:         3,
// 						TimeInterval:     domain.Duration{Duration: time.Hour * 24},
// 					}, nil
// 				},
// 				GetNotificationByTypeAndUserFunc: func(s1 string, s2 string) (*domain.Notification, error) {
// 					return &domain.Notification{
// 						Timestamp: timestamp_test,
// 						Count:     1,
// 					}, nil
// 				},
// 				IncrementNotificationCountFunc: func(s1 string, s2 string) error {
// 					return fmt.Errorf("some error")
// 				},
// 			},
// 			expectedErr: fmt.Errorf("some error"),
// 		},
// 		{
// 			name: "success within interval - increment counter",
// 			mockedContainer: &ContainerMock{
// 				GetRuleByTypeFunc: func(ns string) (*domain.RateLimitRule, error) {
// 					return &domain.RateLimitRule{
// 						NotificationType: "example",
// 						MaxLimit:         3,
// 						TimeInterval:     domain.Duration{Duration: time.Hour * 24},
// 					}, nil
// 				},
// 				GetNotificationByTypeAndUserFunc: func(s1 string, s2 string) (*domain.Notification, error) {
// 					return &domain.Notification{
// 						Timestamp: timestamp_test,
// 						Count:     1,
// 					}, nil
// 				},
// 				IncrementNotificationCountFunc: func(s1 string, s2 string) error {
// 					return nil
// 				},
// 			},
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			service := &NotificationService{container: tc.mockedContainer}
// 			err := service.SendNotification("recipient", "type")
// 			assert.Equal(t, tc.expectedErr, err)
// 		})
// 	}
// }
