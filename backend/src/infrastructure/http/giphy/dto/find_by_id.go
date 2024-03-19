package dto

type FindResponse struct {
	Data FindData     `json:"data"`
	Meta MetaResponse `json:"meta"`
}

type FindData struct {
	Images FindDataImages `json:"images"`
}

type FindDataImages struct {
	Original ImagesOriginal `json:"url"`
}

type ImagesOriginal struct {
	URL string `json:"url"`
}
