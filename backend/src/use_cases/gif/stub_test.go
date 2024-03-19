package gif

import (
	"giphy/challenge/src/domain/entities"
	"time"
)

func GIFWithURLStub() *entities.GIF {
	return &entities.GIF{
		Name:    "name",
		Content: "content",
		User:    "user",
		Tags:    []string{"tag1", "tag2"},
		URL:     "http://google.com",
	}
}

func GIFStub() *entities.GIF {
	return &entities.GIF{
		ID:        "123",
		Name:      "name",
		URL:       "http//google.com",
		Content:   "content",
		User:      "user",
		Tags:      []string{"tag1", "tag2"},
		CreatedAt: &time.Time{},
	}
}

func GIFToBeCreatedStub() *entities.GIF {
	return &entities.GIF{
		Name:    "name",
		Content: "content",
		User:    "user",
		Tags:    []string{"tag1", "tag2"},
	}
}

func PaginationStub() *entities.Pagination[entities.GIF] {
	return &entities.Pagination[entities.GIF]{
		TotalRows:   1,
		Page:        1,
		RowsPerPage: 1,
		Results:     []*entities.GIF{GIFStub()},
	}
}

func GIFSearchStub() *entities.GIFSearch {
	return &entities.GIFSearch{
		Query: "query",
		Page:  1,
	}
}
