package dto

type GIFResponse struct {
	ID     string            `json:"id"`
	Title  string            `json:"title"`
	Images GIFImagesResponse `json:"images"`
}

type GIFImagesResponse struct {
	Original ImageResponse `json:"original"`
}

type ImageResponse struct {
	URL string `json:"url"`
}
