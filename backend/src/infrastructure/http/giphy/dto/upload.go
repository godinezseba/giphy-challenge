package dto

type UploadResponse struct {
	Data UploadDataResponse `json:"data"`
	Meta MetaResponse       `json:"meta"`
}

type UploadDataResponse struct {
	ID string `json:"id"`
}
