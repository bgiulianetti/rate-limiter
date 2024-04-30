package notifications

import (
	"rate-limiter/domain"
	"strings"
	"sync"
	"time"
)

type InMemoryNotificationsContainer struct {
	notifications map[string][]*domain.Notification
	mutex         *sync.Mutex
}

func NewInMemoryNotificationsContainer() *InMemoryNotificationsContainer {
	return &InMemoryNotificationsContainer{
		notifications: map[string][]*domain.Notification{},
		mutex:         &sync.Mutex{},
	}
}

func (ic *InMemoryNotificationsContainer) GetNotifications() (map[string][]*domain.Notification, error) {
	return ic.notifications, nil
}

func (ic *InMemoryNotificationsContainer) GetNotificationsByUser(params domain.GetNotificationParams) ([]*domain.Notification, error) {
	notificationsToReturn := []*domain.Notification{}
	startTime := time.Now().Add(-params.TimeInterval)
	for _, notification := range ic.notifications[params.UserID] {
		if notification.Timestamp.After(startTime) && notification.Type == params.NotificationType {
			notificationsToReturn = append(notificationsToReturn, notification)
		}
	}
	return notificationsToReturn, nil
}

func (ic *InMemoryNotificationsContainer) AddNotification(userID, notificationType string) error {
	ic.notifications[userID] = append(ic.notifications[userID], &domain.Notification{
		Timestamp: time.Now(),
		UserID:    userID,
		Type:      strings.ToLower(notificationType),
	})
	return nil
}
