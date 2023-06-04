package __mock__

import "github.com/T2-1c2023/RecommendationService/app/model"

type RulesRepositoryMock struct{}

func (repo *RulesRepositoryMock) GetInterestsRule() (model.InterestsRule, error) {
	rule := model.InterestsRule{
		Enabled: true,
	}
	return rule, nil
}

func (repo *RulesRepositoryMock) UpdateInterestsRule(newRule *model.InterestsRule) error {
	return nil
}

func (repo *RulesRepositoryMock) GetProximityRule() (model.ProximityRule, error) {
	rule := model.ProximityRule{
		Radius:  3,
		Enabled: true,
	}
	return rule, nil
}
func (repo *RulesRepositoryMock) UpdateProximityRule(newRule *model.ProximityRule) error {
	return nil
}
