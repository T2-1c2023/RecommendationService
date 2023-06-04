package __mock__

import "github.com/T2-1c2023/RecommendationService/app/model"

type RulesRepositoryMock struct {
	InterestsRuleEnabled bool
	ProximityRuleEnabled bool
}

func NewRulesRepositoryMock(interestsRuleEnabled bool, proximityRuleEnabled bool) RulesRepositoryMock {
	return RulesRepositoryMock{
		InterestsRuleEnabled: interestsRuleEnabled,
		ProximityRuleEnabled: proximityRuleEnabled,
	}
}

func (repo *RulesRepositoryMock) GetInterestsRule() (model.InterestsRule, error) {
	rule := model.InterestsRule{
		Enabled: repo.InterestsRuleEnabled,
	}
	return rule, nil
}

func (repo *RulesRepositoryMock) UpdateInterestsRule(newRule *model.InterestsRule) error {
	return nil
}

func (repo *RulesRepositoryMock) GetProximityRule() (model.ProximityRule, error) {
	rule := model.ProximityRule{
		Radius:  3,
		Enabled: repo.ProximityRuleEnabled,
	}
	return rule, nil
}
func (repo *RulesRepositoryMock) UpdateProximityRule(newRule *model.ProximityRule) error {
	return nil
}
