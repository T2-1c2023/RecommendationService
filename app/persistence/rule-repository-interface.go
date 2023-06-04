package persistence

import (
	"github.com/T2-1c2023/RecommendationService/app/model"
)

type RulesPersistence interface {
	GetInterestsRule() (model.InterestsRule, error)
	UpdateInterestsRule(newRule *model.InterestsRule) error
	GetProximityRule() (model.ProximityRule, error)
	UpdateProximityRule(newRule *model.ProximityRule) error
}
