package dto

type FindByIDResponse struct {
	Data GIFResponse  `json:"data"`
	Meta MetaResponse `json:"meta"`
}

type FindBySearchResponse struct {
	Data []*GIFResponse `json:"data"`
	Meta MetaResponse   `json:"meta"`
}
