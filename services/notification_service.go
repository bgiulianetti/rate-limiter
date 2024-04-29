package services

import (
	"fmt"
	"rate-limiter/dao"
	"rate-limiter/domain"
	"rate-limiter/errors"
	"time"
)

type Service interface {
	GetRules() (map[string]*domain.RateLimitRule, error)
	GetRuleByType(string) (*domain.RateLimitRule, error)
	GetNotifications() (map[string]map[string]*domain.Notification, error)
	SendNotification(string, string) error
}

type NotificationService struct {
	container dao.Container
}

func NewService() *NotificationService {
	return &NotificationService{
		container: dao.NewContainer(),
	}
}

func (ns *NotificationService) GetRules() (map[string]*domain.RateLimitRule, error) {
	return ns.container.GetRules()
}

func (ns *NotificationService) GetRuleByType(notificationType string) (*domain.RateLimitRule, error) {
	return ns.container.GetRuleByType(notificationType)
}

func (ns *NotificationService) GetNotifications() (map[string]map[string]*domain.Notification, error) {
	return ns.container.GetNotifications()
}

func (ns *NotificationService) GetNotificationsByType(notificationType string) (map[string]*domain.Notification, error) {
	return ns.container.GetNotificationsByType(notificationType)
}

func (ns *NotificationService) SendNotification(recipient, notificationType string) error {
	rule, err := ns.container.GetRuleByType(notificationType)
	if err != nil {
		return errors.ErrGetRateLimitRule
	}

	if rule == nil {
		sendEmail(recipient)
		return nil
	}

	err = ns.checkRateLimit(recipient, rule)
	if err != nil {
		return err
	}

	sendEmail(recipient)
	return nil
}

func (ns *NotificationService) checkRateLimit(recipient string, rule *domain.RateLimitRule) error {
	notification, err := ns.container.GetNotificationByTypeAndUser(rule.NotificationType, recipient)
	if err != nil {
		return err
	}

	if notification == nil {
		err := ns.container.IncrementNotificationCount(rule.NotificationType, recipient)
		if err != nil {
			return err
		}
	} else {
		interval := time.Since(notification.Timestamp)
		if interval >= rule.TimeInterval.Duration {
			fmt.Println("interval passed. reset counter")
			err := ns.container.ResetNotificationCount(rule.NotificationType, recipient)
			if err != nil {
				return err
			}
		} else if notification.Count >= rule.MaxLimit {
			fmt.Println("within interval. max exceeded")
			return errors.ErrRateLimitExceeded
		} else {
			fmt.Println("within interval. increment counter")
			err := ns.container.IncrementNotificationCount(rule.NotificationType, recipient)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func sendEmail(recipient string) {
	fmt.Printf("Email sent to %s\n", recipient)
}
