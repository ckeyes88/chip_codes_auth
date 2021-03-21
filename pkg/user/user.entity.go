package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	PasswordHash string             `json:"-" bson:"passwordHash"`
	RefreshToken string             `json:"-" bson:"refreshToken"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`
	DeletedAt    *time.Time         `json:"deletedAt" bson:"deletedAt"`
}
