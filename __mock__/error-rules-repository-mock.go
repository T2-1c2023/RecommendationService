package __mock__

import (
	"fmt"

	"github.com/T2-1c2023/RecommendationService/app/model"
)

type ErrorRulesRepositoryMock struct{}

func (repo *ErrorRulesRepositoryMock) GetInterestsRule() (model.InterestsRule, error) {
	var rule model.InterestsRule
	return rule, fmt.Errorf("Test")
}

func (repo *ErrorRulesRepositoryMock) UpdateInterestsRule(newRule *model.InterestsRule) error {
	return fmt.Errorf("Test")
}

func (repo *ErrorRulesRepositoryMock) GetProximityRule() (model.ProximityRule, error) {
	var rule model.ProximityRule
	return rule, fmt.Errorf("Test")
}
func (repo *ErrorRulesRepositoryMock) UpdateProximityRule(newRule *model.ProximityRule) error {
	return fmt.Errorf("Test")
}
