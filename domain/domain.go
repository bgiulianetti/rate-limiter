package domain

import (
	"encoding/json"
	"rate-limiter/utils"
	"time"
)

type RateLimitRule struct {
	NotificationType string   `json:"notificationType"`
	MaxLimit         int      `json:"maxLimit"`
	TimeInterval     Duration `json:"timeInterval"`
}

type Notification struct {
	Timestamp time.Time `json:"timeStamp"`
	Count     int       `json:"count"`
}

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var durationStr string
	if err := json.Unmarshal(b, &durationStr); err != nil {
		return err
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return err
	}

	d.Duration = duration
	return nil
}

func (d Duration) MarshalJSON() ([]byte, error) {
	durationStr := utils.FormatDuration(d.Duration)
	return json.Marshal(durationStr)
}
