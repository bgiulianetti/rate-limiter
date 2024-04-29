package dao

import (
	"fmt"
	"os"
	"rate-limiter/domain"
)

type Container interface {
	GetRules() (map[string]*domain.RateLimitRule, error)
	GetRuleByType(string) (*domain.RateLimitRule, error)
	GetNotifications() (map[string]map[string]*domain.Notification, error)
	GetNotificationsByType(string) (map[string]*domain.Notification, error)
	GetNotificationByTypeAndUser(string, string) (*domain.Notification, error)
	IncrementNotificationCount(string, string) error
	ResetNotificationCount(string, string) error
}

func NewContainer() Container {
	daoType := os.Getenv("DAO_TYPE")
	switch daoType {
	case "memory":
		return NewInMemoryContainer()
	case "mongoDB":
		return NewMongoDBContainer()
	default:
		fmt.Printf("unknown DAO type: %s", daoType)
		return NewInMemoryContainer()
	}
}
