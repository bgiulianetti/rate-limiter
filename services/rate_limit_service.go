package services

import (
	"fmt"
	"rate-limiter/domain"
	"rate-limiter/errors"
)

type NotificationsContainer interface {
	AddNotification(userID, notificationType string) error
	GetNotificationsByUser(params domain.GetNotificationParams) ([]*domain.Notification, error)
}

type CommunicationClient interface {
	Send(string) error
}

type RateLimitService struct {
	notificationsContainer NotificationsContainer
	rulesService           *RulesService
}

func NewRateLimitService(notificationsContainer NotificationsContainer, rulesService *RulesService) *RateLimitService {
	return &RateLimitService{
		notificationsContainer: notificationsContainer,
		rulesService:           rulesService,
	}
}

func (ns *RateLimitService) SendNotification(params domain.SendNotificationParams) error {

	//rules preguntar
	//peguntar poer sendEmail
	rule, err := ns.rulesService.GetRuleByType(params.NotificationType)
	if err != nil {
		return errors.ErrGetRateLimitRule
	}

	if rule == nil {
		ns.sendEmail(params.UserID)
		return nil
	}

	//for{
	allow, err := ns.checkRateLimit(params.UserID, rule)
	if err != nil {
		return err
	}

	if !allow {
		return errors.ErrRateLimitExceeded
	}
	//}

	err = ns.sendEmail(params.UserID)
	if err != nil {
		return err
	}
	err = ns.notificationsContainer.AddNotification(params.UserID, params.NotificationType)
	if err != nil {
		fmt.Println("Error registering notification")
	}
	return nil
}

func (ns *RateLimitService) checkRateLimit(userID string, rule *domain.RateLimitRule) (bool, error) {
	notifications, err := ns.notificationsContainer.GetNotificationsByUser(domain.GetNotificationParams{
		UserID:           userID,
		NotificationType: rule.NotificationType,
		TimeInterval:     rule.TimeInterval.Duration,
	})
	if err != nil {
		return false, err
	}

	if len(notifications) >= rule.MaxLimit {
		fmt.Println("Within interval. max exceeded")
		return false, errors.ErrRateLimitExceeded
	}

	return true, nil
}

func (ns *RateLimitService) sendEmail(userID string) error {
	fmt.Printf("Email sent to %s\n", userID)
	return nil
}
