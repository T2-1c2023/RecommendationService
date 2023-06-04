package persistence

import (
	"context"
	"fmt"

	"github.com/T2-1c2023/RecommendationService/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IRulesRepository interface {
	GetInterestsRule() (model.InterestsRule, error)
	UpdateInterestsRule(newRule *model.InterestsRule) error
	GetProximityRule() (model.ProximityRule, error)
	UpdateProximityRule(newRule *model.ProximityRule) error
}

type RulesRepository struct {
	Collection *mongo.Collection
}

func (repo *RulesRepository) GetInterestsRule() (model.InterestsRule, error) {
	filter := bson.M{"rule": "interests"}
	result := repo.Collection.FindOne(context.Background(), filter)
	var rule model.InterestsRule
	if result == nil {
		return rule, fmt.Errorf("rule not found")
	}

	result.Decode(&rule)
	return rule, nil
}

func (repo *RulesRepository) UpdateInterestsRule(newRule *model.InterestsRule) error {
	filter := bson.M{"rule": "interests"}
	update := bson.M{
		"$set": bson.M{
			"enabled": newRule.Enabled,
		},
	}
	_, err := repo.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (repo *RulesRepository) GetProximityRule() (model.ProximityRule, error) {
	filter := bson.M{"rule": "proximity"}
	result := repo.Collection.FindOne(context.Background(), filter)
	var rule model.ProximityRule
	if result == nil {
		return rule, fmt.Errorf("rule not found")
	}

	result.Decode(&rule)
	return rule, nil
}

func (repo *RulesRepository) UpdateProximityRule(newRule *model.ProximityRule) error {
	filter := bson.M{"rule": "proximity"}
	update := bson.M{
		"$set": bson.M{
			"radius":  newRule.Radius,
			"enabled": newRule.Enabled,
		},
	}
	_, err := repo.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
