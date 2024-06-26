package controllers

import (
	"net/http"
	"net/http/httptest"
	"rate-limiter/domain"
	"rate-limiter/errors"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNotificationController_Pong(t *testing.T) {
	testCases := []struct {
		name             string
		expectedResponse string
	}{
		{
			name:             "success",
			expectedResponse: `{"message":"Pong from Notifications","status":"success"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)
			controller := NotificationController{}

			controller.Pong(context)
			recorder.Flush()

			assert.Equal(t, http.StatusOK, recorder.Code)

			responseBody := recorder.Body.String()
			assert.Equal(t, tc.expectedResponse, responseBody)
		})
	}
}

func TestNotificationController_SendNotification(t *testing.T) {
	testCases := []struct {
		name                       string
		userID                     string
		notificationType           string
		expectedCode               int
		expectedResponse           string
		rateLimitServiceMockConfig func(*RateLimitServiceMock)
	}{
		{
			name:             "missing userID",
			userID:           "",
			notificationType: "testType",
			expectedCode:     http.StatusBadRequest,
			expectedResponse: `{"message":"userID is mandatory","error":"invalid_user_id","status":400}`,
			rateLimitServiceMockConfig: func(mock *RateLimitServiceMock) {
				mock.SendNotificationFunc = func(domain.SendNotificationParams) error {
					return nil
				}
			},
		},
		{
			name:             "missing notification type",
			userID:           "testUserID",
			notificationType: "",
			expectedCode:     http.StatusBadRequest,
			expectedResponse: `{"message":"notification type is mandatory","error":"invalid_type","status":400}`,
			rateLimitServiceMockConfig: func(mock *RateLimitServiceMock) {
				mock.SendNotificationFunc = func(domain.SendNotificationParams) error {
					return nil
				}
			},
		},
		{
			name:             "success",
			userID:           "testUserID",
			notificationType: "testType",
			expectedCode:     http.StatusOK,
			expectedResponse: `{"message":"notification sent","status":"success"}`,
			rateLimitServiceMockConfig: func(mock *RateLimitServiceMock) {
				mock.SendNotificationFunc = func(domain.SendNotificationParams) error {
					return nil
				}
			},
		},
		{
			name:             "limit exceeded",
			userID:           "testUserID",
			notificationType: "testType",
			expectedCode:     http.StatusTooManyRequests,
			expectedResponse: `{"message":"message limit exceeded","error":"rate limit exceeded","status":429}`,
			rateLimitServiceMockConfig: func(mock *RateLimitServiceMock) {
				mock.SendNotificationFunc = func(domain.SendNotificationParams) error {
					return errors.ErrRateLimitExceeded
				}
			},
		},
		{
			name:             "error getting rule limit",
			userID:           "testUserID",
			notificationType: "testType",
			expectedCode:     http.StatusInternalServerError,
			expectedResponse: `{"message":"internal server error","error":"error getting rate limit rule for notification type","status":500}`,
			rateLimitServiceMockConfig: func(mock *RateLimitServiceMock) {
				mock.SendNotificationFunc = func(domain.SendNotificationParams) error {
					return errors.ErrGetRateLimitRule
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)

			serviceMock := &RateLimitServiceMock{}
			tc.rateLimitServiceMockConfig(serviceMock)

			controller := NotificationController{
				RateLimitService: serviceMock,
			}

			context.Set("userID", tc.userID)
			context.Set("type", tc.notificationType)

			controller.SendNotification(context)
			assert.Equal(t, tc.expectedCode, recorder.Code)
			assert.Equal(t, tc.expectedResponse, recorder.Body.String())
		})
	}
}
