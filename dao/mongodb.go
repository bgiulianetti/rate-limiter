package dao

import (
	"context"
	"log"
	"rate-limiter/domain"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBContainer struct {
	Client     *mongo.Client
	DB         string
	Collection string
}

func NewMongoDBContainer() *MongoDBContainer {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("uri-here"))
	if err != nil {
		log.Fatal(err)
	}

	mongoDBContainer := &MongoDBContainer{
		Client:     client,
		DB:         "rate-limit",
		Collection: "notifications",
	}
	return mongoDBContainer
}

func (mc *MongoDBContainer) GetRules() (map[string]*domain.RateLimitRule, error) {
	return nil, nil
}

func (mc *MongoDBContainer) GetRuleByType(notificationType string) (*domain.RateLimitRule, error) {
	return nil, nil
}

func (mc *MongoDBContainer) GetNotifications() (map[string]map[string]*domain.Notification, error) {
	return nil, nil
}

func (mc *MongoDBContainer) GetNotificationsByType(notificationType string) (map[string]*domain.Notification, error) {
	return nil, nil
}

func (mc *MongoDBContainer) GetNotificationByTypeAndUser(notificationType, userID string) (*domain.Notification, error) {
	return nil, nil
}

func (mc *MongoDBContainer) IncrementNotificationCount(userID, notificationType string) error {
	return nil
}

func (mc *MongoDBContainer) ResetNotificationCount(notificationType, userID string) error {
	return nil
}
