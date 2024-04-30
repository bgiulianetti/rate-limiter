package notifications

import (
	"context"
	"encoding/json"
	"rate-limiter/domain"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisContainer struct {
	Client *redis.Client
	//DB     string
}

func NewRedisContainer() *RedisContainer {
	addr, err := redis.ParseURL("some_uri")
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(addr)
	return &RedisContainer{
		Client: rdb,
	}
}

func (rc *RedisContainer) GetNotifications() (map[string][]*domain.Notification, error) {
	notificationsJSON, err := rc.Client.Get(context.Background(), "notifications").Result()
	if err != nil {
		return nil, err
	}

	var notifications map[string][]*domain.Notification
	if err := json.Unmarshal([]byte(notificationsJSON), &notifications); err != nil {
		return nil, err
	}

	return notifications, nil
}

func (rc *RedisContainer) GetNotificationsByUser(userID string) ([]*domain.Notification, error) {
	notifications, err := rc.GetNotifications()
	if err != nil {
		return nil, err
	}

	return notifications[userID], nil
}

func (rc *RedisContainer) GetNotificationsByUserAndTypeAndInterval(params domain.GetNotificationParams) ([]*domain.Notification, error) {
	notifications, err := rc.GetNotifications()
	if err != nil {
		return nil, err
	}

	notificationsToReturn := []*domain.Notification{}
	timestamp := time.Now().Add(-params.TimeInterval)
	for _, notification := range notifications[params.UserID] {
		if notification.Timestamp.After(timestamp) && notification.Type == params.NotificationType {
			notificationsToReturn = append(notificationsToReturn, notification)
		}
	}
	return notificationsToReturn, nil
}

func (rc *RedisContainer) AddNotification(userID, notificationType string) error {
	notifications, err := rc.GetNotifications()
	if err != nil {
		return err
	}

	notifications[userID] = append(notifications[userID], &domain.Notification{
		Timestamp: time.Now(),
		UserID:    userID,
		Type:      notificationType,
	})

	notificationsJSON, err := json.Marshal(notifications)
	if err != nil {
		return err
	}

	err = rc.Client.Set(context.Background(), "notifications", notificationsJSON, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
