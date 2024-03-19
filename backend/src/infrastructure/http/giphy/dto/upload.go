package dto

type UploadRequest struct {
	ApiKey string `json:"api_key"`
	File   string `json:"file"`
	Tags   string `json:"tags"`
}

type UploadResponse struct {
	Data UploadDataResponse `json:"data"`
	Meta MetaResponse       `json:"meta"`
}

type UploadDataResponse struct {
	ID string `json:"id"`
}
