package dao

import (
	"encoding/json"
	"fmt"
	"rate-limiter/domain"
	"rate-limiter/utils"
	"strings"
	"sync"
	"time"
)

func NewInMemoryContainer() *InMemoryContainer {
	return &InMemoryContainer{
		rules:         setInitialRules(),
		notifications: map[string]map[string]*domain.Notification{},
		mutex:         &sync.Mutex{},
	}
}

type InMemoryContainer struct {
	rules         map[string]*domain.RateLimitRule
	notifications map[string]map[string]*domain.Notification
	mutex         *sync.Mutex
}

func (ic *InMemoryContainer) GetRules() (map[string]*domain.RateLimitRule, error) {
	return ic.rules, nil
}

func (ic *InMemoryContainer) GetRuleByType(notificationType string) (*domain.RateLimitRule, error) {
	return ic.rules[notificationType], nil
}

func (ic *InMemoryContainer) GetNotifications() (map[string]map[string]*domain.Notification, error) {
	return ic.notifications, nil
}

func (ic *InMemoryContainer) GetNotificationsByType(notificationType string) (map[string]*domain.Notification, error) {
	return ic.notifications[notificationType], nil
}

func (ic *InMemoryContainer) GetNotificationByTypeAndUser(notificationType, userID string) (*domain.Notification, error) {
	return ic.notifications[notificationType][userID], nil
}

func (ic *InMemoryContainer) IncrementNotificationCount(notificationType, userID string) error {

	if ic.notifications[notificationType] == nil {
		ic.notifications[notificationType] = make(map[string]*domain.Notification)
	}

	if _, exists := ic.notifications[notificationType][userID]; !exists {
		ic.notifications[notificationType][userID] = &domain.Notification{
			Count:     1,
			Timestamp: time.Now(),
		}
	} else {
		ic.notifications[notificationType][userID].Count++
		ic.notifications[notificationType][userID].Timestamp = time.Now()

	}
	return nil
}

func (ic *InMemoryContainer) ResetNotificationCount(notificationType, userID string) error {
	if ic.notifications[notificationType] != nil {
		notification := ic.notifications[notificationType][userID]
		if notification != nil {
			notification.Timestamp = time.Now()
			notification.Count = 1
		}
	}

	return nil
}

func setInitialRules() map[string]*domain.RateLimitRule {

	fileData := utils.LoadFile()
	var rules []domain.RateLimitRule
	if err := json.Unmarshal(fileData, &rules); err != nil {
		fmt.Println("error unarshaling rules.json file:::", err)
	}

	ruleMap := make(map[string]*domain.RateLimitRule)
	for _, rule := range rules {
		ruleMap[strings.ToLower(rule.NotificationType)] = &rule
	}
	return ruleMap
}
