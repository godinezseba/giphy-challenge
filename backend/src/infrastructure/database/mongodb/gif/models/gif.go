package models

type GIF struct {
	ID         string   `bson:"_id,omitempty"`
	ProviderID string   `bson:"providerId"`
	Name       string   `bson:"name"`
	URL        string   `bson:"url"`
	User       string   `bson:"user"`
	Tags       []string `bson:"tags"`
	CreatedAt  string   `bson:"createdAt"`
}
