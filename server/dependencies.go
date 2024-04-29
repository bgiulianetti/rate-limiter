package server

import (
	"rate-limiter/controllers"
	"rate-limiter/dao"
	"rate-limiter/services"
)

func resolveNotificationController() *controllers.NotificationController {

	rulesContainer := dao.NewRulesContainer()
	controller := &controllers.NotificationController{
		NotificationService: services.NewNotificationService(
			dao.NewNotificationContainer(),
			rulesContainer,
		),
		RulesService: services.NewRulesService(
			rulesContainer,
		),
	}
	return controller
}
