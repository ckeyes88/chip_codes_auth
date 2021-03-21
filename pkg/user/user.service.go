package user

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IUser interface {
	CreateUser(ctx context.Context, u User) error
}

type UserService struct {
	userCollection *mongo.Collection
}

func (us *UserService) CreateUser(ctx context.Context, u User) error {
	now := time.Now()
	t := true

	_, err := us.userCollection.UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{"$set": bson.M{
		"name":         u.Name,
		"passwordHash": u.PasswordHash,
		"refreshToken": u.RefreshToken,
		"createdAt":    now,
		"updatedAt":    now,
	}}, &options.UpdateOptions{Upsert: &t})
	return err
}
