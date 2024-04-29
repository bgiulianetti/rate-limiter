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

type CommunicationClient interface {
	Send(string) error
}

type NotificationService struct {
	notificationsContainer NotificationsContainer
	rulesContainer         RulesContainer
	communicationClient    CommunicationClient
}

func NewNotificationService(notificationsContainer NotificationsContainer, rulesContainer RulesContainer) *NotificationService {
	return &NotificationService{
		notificationsContainer: notificationsContainer,
		rulesContainer:         rulesContainer,
	}
}

func (ns *NotificationService) GetNotificationsByUser(userID string) ([]*domain.Notification, error) {
	return ns.notificationsContainer.GetNotificationsByUser(userID)
}

func (ns *NotificationService) SendNotification(notificationParams domain.SendNotificationParams) error {

	rule, err := ns.rulesContainer.GetRuleByType(notificationParams.NotificationType)
	if err != nil {
		return errors.ErrGetRateLimitRule
	}

	if rule == nil {
		ns.sendEmail(notificationParams.UserID)
		return nil
	}

	err = ns.checkRateLimit(notificationParams.UserID, rule)
	if err != nil {
		return err
	}

	ns.sendEmail(notificationParams.UserID)
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

func (ns *NotificationService) sendEmail(userID string) {
	err := ns.communicationClient.Send(userID)
	if err != nil {
		fmt.Printf("Error sending email to %s\n", userID)
		return
	}
	fmt.Printf("Email sent to %s\n", userID)
}
