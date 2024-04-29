package notifications

import (
	"rate-limiter/domain"

	"github.com/redis/go-redis/v9"
)

type RedisContainer struct {
	Client *redis.Client
	DB     string
}

func NewRedisContainer() *RedisContainer {
	return nil
}

func (rc *RedisContainer) GetRules() (map[string]*domain.RateLimitRule, error) {
	return nil, nil
}

func (rc *RedisContainer) GetRuleByType(notificationType string) (*domain.RateLimitRule, error) {
	return nil, nil
}

func (rc *RedisContainer) GetNotifications() (map[string]map[string]*domain.Notification, error) {
	return nil, nil
}

func (rc *RedisContainer) GetNotificationsByType(notificationType string) (map[string]*domain.Notification, error) {
	return nil, nil
}

func (rc *RedisContainer) GetNotificationByTypeAndUser(notificationType, userID string) (*domain.Notification, error) {
	return nil, nil
}

func (rc *RedisContainer) IncrementNotificationCount(userID, notificationType string) error {
	return nil
}

func (rc *RedisContainer) ResetNotificationCount(notificationType, userID string) error {
	return nil
}

func (rc *RedisContainer) SetInitialRules() (map[string]*domain.RateLimitRule, error) {
	return nil, nil
}
