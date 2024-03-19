package models

type GIF struct {
	ID        string   `bson:"_id,omitempty"`
	Name      string   `bson:"name"`
	URL       string   `bson:"url"`
	User      string   `bson:"user"`
	Tags      []string `bson:"tags"`
	CreatedAt string   `bson:"createdAt"`
}
