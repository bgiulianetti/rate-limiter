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
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://bgiulianetti:mongodb@cluster0.0nvl0.mongodb.net/minesweper?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	mongoDBContainer := &MongoDBContainer{
		Client:     client,
		DB:         "minesweper",
		Collection: "notifications",
	}
	return mongoDBContainer
}

func (mc *MongoDBContainer) GetRules() (map[string]*domain.RateLimitRule, error) {

	// var rules map[string]*domain.RateLimitRule
	// collection := mdb.Client.Database(mdb.DB).Collection(mdb.Collection)
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// cursor, err := collection.Find(ctx, bson.M{})
	// if err != nil {
	// 	return nil, err
	// }
	// defer cursor.Close(ctx)
	// for cursor.Next(ctx) {
	// 	var rule *domain.RateLimitRule
	// 	cursor.Decode(rule)
	// 	rules[rule.NotificationType] = rule
	// }
	// if err := cursor.Err(); err != nil {
	// 	return nil, err
	// }
	// return rules, nil
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
