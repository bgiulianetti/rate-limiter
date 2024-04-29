package services

import (
	"fmt"
	"rate-limiter/dao"
	"rate-limiter/domain"
	"rate-limiter/errors"
	"rate-limiter/utils"
	"time"
)

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
	fmt.Println("i get here")
	return ns.container.GetNotificationsByType(notificationType)
}

func (ns *NotificationService) SendNotification(recipient, notificationType string) error {
	rule, err := ns.container.GetRuleByType(notificationType)
	fmt.Println("rule", utils.SerializeObject(rule))
	if err != nil {
		return errors.ErrGetRateLimitRule
	}

	if rule == nil {
		// no rule for this notification type. send email
		sendEmail(recipient)
		return nil
	}

	// Check if the notification should be rate limited
	err = ns.checkRateLimit(recipient, rule)
	if err != nil {
		return err
	}

	// If rate limit check passes, send email
	sendEmail(recipient)
	return nil
}

func (ns *NotificationService) checkRateLimit(recipient string, rule *domain.RateLimitRule) error {
	notification, err := ns.container.GetNotificationByTypeAndUser(rule.NotificationType, recipient)
	fmt.Println("notification by type and user", utils.SerializeObject(notification))
	if err != nil {
		return err
	}

	if notification == nil {
		fmt.Println("Create Notification")
		err := ns.container.IncrementNotificationCount(rule.NotificationType, recipient)
		fmt.Println("notification was created")
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Notification exists")
		// Check if the time interval has elapsed, if yes, reset the count
		interval := time.Now().Sub(notification.Timestamp)
		if interval >= rule.TimeInterval.Duration {
			fmt.Println("interval passed. reset counter")
			err := ns.container.ResetNotificationCount(rule.NotificationType, recipient)
			if err != nil {
				return err
			}
		} else if notification.Count >= rule.MaxLimit {
			// If within the time interval and count exceeds max limit, reject notification
			fmt.Println("within interval. max exceeded")
			return errors.ErrRateLimitExceeded
		} else {
			// Increment the notification count
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
	//get recipient email
	//send email to recipient
	fmt.Printf("Email sent to %s\n", recipient)
}
