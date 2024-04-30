package notifications

import (
	"rate-limiter/domain"
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
	ic.mutex.Lock()
	defer ic.mutex.Unlock()

	return ic.notifications, nil
}

func (ic *InMemoryNotificationsContainer) GetNotificationsByUser(params domain.GetNotificationParams) ([]*domain.Notification, error) {
	ic.mutex.Lock()
	defer ic.mutex.Unlock()

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
		Type:      notificationType,
	})
	return nil
}
