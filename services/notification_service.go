package services

import (
	"fmt"
	"rate-limiter/domain"
	"rate-limiter/errors"
)

type NotificationsContainer interface {
	AddNotification(userID, notificationType string) error
	GetNotificationsByUser(userID string) ([]*domain.Notification, error)
	GetNotificationsByUserAndTypeAndInterval(params domain.GetNotificationParams) ([]*domain.Notification, error)
}

//struct titulo, mensaje
// type CommunicationClient struct{
//  	send(string) error/*aca va un strucvt*/ error
// }

type NotificationService struct {
	notificationsContainer NotificationsContainer
	rulesContainer         RulesContainer
	//notificaionClient
}

func NewNotificationService(notificationsContainer NotificationsContainer, rulesContainer RulesContainer) *NotificationService {
	return &NotificationService{
		notificationsContainer: notificationsContainer,
		rulesContainer:         rulesContainer,
	}
}

// type NotificationStorage interface {
// 	add(userID, notificationType string) //, ttl time.Duration)
// 	getNotifications(userID, notificationType string, intervalTime time.Duration) []domain.Notification
// }

func (ns *NotificationService) GetNotificationsByUser(userID string) ([]*domain.Notification, error) {
	return ns.notificationsContainer.GetNotificationsByUser(userID)
}

func (ns *NotificationService) SendNotification(notificationParams domain.SendNotificationParams) error {

	//array de reglas
	rule, err := ns.rulesContainer.GetRuleByType(notificationParams.NotificationType)
	if err != nil {
		return errors.ErrGetRateLimitRule
	}

	if rule == nil {
		sendEmail(notificationParams.UserID)
		return nil
	}

	err = ns.checkRateLimit(notificationParams.UserID, rule)
	if err != nil {
		return err
	}

	sendEmail(notificationParams.UserID)
	return nil
}

func (ns *NotificationService) checkRateLimit(userID string, rule *domain.RateLimitRule) error {
	notifications, err := ns.notificationsContainer.GetNotificationsByUserAndTypeAndInterval(domain.GetNotificationParams{
		UserID:           userID,
		NotificationType: rule.NotificationType,
		TimeInterval:     rule.TimeInterval.Duration,
	})
	if err != nil {
		return err
	}

	if len(notifications) >= rule.MaxLimit {
		fmt.Println("Within interval. max exceeded")
		return errors.ErrRateLimitExceeded
	}

	err = ns.notificationsContainer.AddNotification(userID, rule.NotificationType)
	if err != nil {
		return err
	}

	return nil
}

func sendEmail(recipient string) {
	//get email from, userID
	fmt.Printf("Email sent to %s\n", recipient)
}
