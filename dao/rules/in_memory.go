package rules

import (
	"encoding/json"
	"fmt"
	"rate-limiter/domain"
	"rate-limiter/utils"
	"strings"
	"sync"
)

type InMemoryRulesContainer struct {
	rules map[string][]*domain.RateLimitRule
	mutex *sync.Mutex
}

func NewInMemoryRulesContainer() *InMemoryRulesContainer {
	rules := setInitialRules()
	return &InMemoryRulesContainer{
		rules: rules,
		mutex: &sync.Mutex{},
	}
}

func (ic *InMemoryRulesContainer) GetRules() (map[string][]*domain.RateLimitRule, error) {
	ic.mutex.Lock()
	defer ic.mutex.Unlock()

	return ic.rules, nil
}

func (ic *InMemoryRulesContainer) GetRuleByType(notificationType string) ([]*domain.RateLimitRule, error) {
	ic.mutex.Lock()
	defer ic.mutex.Unlock()

	return ic.rules[notificationType], nil
}

func setInitialRules() map[string][]*domain.RateLimitRule {
	fileData := utils.LoadRulesFile()
	var rules []*domain.RateLimitRule
	if err := json.Unmarshal(fileData, &rules); err != nil {
		fmt.Println("error unarshaling rules.json file", err)
	}

	ruleMap := make(map[string][]*domain.RateLimitRule)
	for _, rule := range rules {
		rule.NotificationType = strings.ToLower(rule.NotificationType)
		_, ok := ruleMap[strings.ToLower(rule.NotificationType)]
		if !ok {
			ruleMap[strings.ToLower(rule.NotificationType)] = []*domain.RateLimitRule{
				rule,
			}
		} else {
			ruleMap[strings.ToLower(rule.NotificationType)] = append(ruleMap[strings.ToLower(rule.NotificationType)], rule)
		}
	}
	fmt.Println("initial rules:", utils.SerializeObject(ruleMap))
	return ruleMap
}
