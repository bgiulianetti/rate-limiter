package dao

import (
	"fmt"
	"os"
	"rate-limiter/dao/notifications"
	"rate-limiter/dao/rules"
	"rate-limiter/services"
)

func NewRulesContainer() services.RulesContainer {
	return rules.NewInMemoryRulesContainer()
}

func NewNotificationContainer() services.NotificationsContainer {
	daoType := getNotificationsDAOType()
	fmt.Printf("Container Notifications DAO_TYPE: %s\n", daoType)
	switch daoType {
	case "memory":
		return notifications.NewInMemoryNotificationsContainer()
	case "redis":
		return notifications.NewRedisContainer()
	default:
		fmt.Printf("unknown Notifications DAO type: '%s'. Load default in memory\n", daoType)
		return notifications.NewInMemoryNotificationsContainer()
	}
}

func getNotificationsDAOType() string {
	return os.Getenv("NOTIFICATIONS_DAO_TYPE")
}
