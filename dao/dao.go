package dao

import (
	"fmt"
	"os"
	"rate-limiter/dao/notifications"
	"rate-limiter/dao/rules"
	"rate-limiter/services"
)

func NewRulesContainer() services.RulesContainer {
	daoType := getDAOType()
	fmt.Printf("Container DAO_TYPE: %s\n", daoType)
	switch daoType {
	case "memory":
		return rules.NewInMemoryContainer()
	case "redis":
		return rules.NewInMemoryContainer()
	default:
		fmt.Printf("unknown DAO type: '%s'. Load default in memory\n", daoType)
		return rules.NewInMemoryContainer()
	}
}

func NewNotificationContainer() services.NotificationsContainer {
	return notifications.NewInMemoryNotificationsContainer()
}

func getDAOType() string {
	return os.Getenv("DAO_TYPE")
}
