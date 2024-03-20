package entities

import "time"

type GIF struct {
	ID         string
	ProviderID string
	Name       string
	URL        string
	Content    string
	User       string
	Tags       []string
	CreatedAt  *time.Time
}

type GIFSearch struct {
	Query    string
	Page     int
	PageSize int
}
