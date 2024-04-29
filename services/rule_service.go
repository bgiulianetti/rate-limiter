package services

import (
	"rate-limiter/domain"
)

type RulesContainer interface {
	GetRules() (map[string]*domain.RateLimitRule, error)
	GetRuleByType(string) (*domain.RateLimitRule, error)
}

type RulesService struct {
	rulesContainer RulesContainer
	//notificaionClient
}

func NewRulesService(rulesContainer RulesContainer) *RulesService {
	return &RulesService{
		rulesContainer: rulesContainer,
	}
}

func (rs *RulesService) GetRules() (map[string]*domain.RateLimitRule, error) {
	return rs.rulesContainer.GetRules()
}

func (rs *RulesService) GetRuleByType(notificationType string) (*domain.RateLimitRule, error) {
	return rs.rulesContainer.GetRuleByType(notificationType)
}
