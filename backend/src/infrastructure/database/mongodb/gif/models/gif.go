package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type GIF struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	ProviderID string             `bson:"providerId"`
	Name       string             `bson:"name"`
	URL        string             `bson:"url"`
	User       string             `bson:"user"`
	Tags       []string           `bson:"tags"`
	CreatedAt  string             `bson:"createdAt"`
}
