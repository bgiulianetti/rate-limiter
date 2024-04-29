package errors

import "errors"

type ApiError struct {
	Message  string `json:"message"`
	ErrorStr string `json:"error"`
	Status   int    `json:"status"`
}

func (e ApiError) Error() string {
	return e.Message
}

var ErrRateLimitExceeded = errors.New("rate limit exceeded")
var ErrGetRateLimitRule = errors.New("error getting rate limit rule for notification type")

func IsTooManyRequestsError(err error) bool {
	return errors.Is(err, ErrRateLimitExceeded)
}
