package notifications

import (
	"context"
	"encoding/json"
	"fmt"
	"rate-limiter/domain"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisContainer struct {
	Client *redis.Client
}

func NewRedisContainer() *RedisContainer {

	ctx := context.Background()

	// Credentials harcoded. It is just a sandbox
	client := redis.NewClient(&redis.Options{
		Addr:     "redis-15763.c84.us-east-1-2.ec2.redns.redis-cloud.com:15763",
		Password: "sJyxJIhJBZJK0l9q0MxIbOEwl2CTEExa",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
	}
	fmt.Println("Connected to Redis:", pong)
	return &RedisContainer{
		Client: client,
	}
}

func (rc *RedisContainer) GetNotifications() (map[string][]*domain.Notification, error) {
	notificationsJSON, err := rc.Client.Get(context.Background(), "notifications").Result()
	if err != nil {
		return map[string][]*domain.Notification{}, nil
	}

	var notifications map[string][]*domain.Notification
	if err := json.Unmarshal([]byte(notificationsJSON), &notifications); err != nil {
		return nil, err
	}
	return notifications, nil
}

func (rc *RedisContainer) GetNotificationsByUser(params domain.GetNotificationParams) ([]*domain.Notification, error) {
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
