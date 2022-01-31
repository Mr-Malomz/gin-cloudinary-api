package dtos

type MediaDto struct {
	StatusCode int                    `json:"statusCode"`
	Message    string                 `json:"message"`
	Data       map[string]interface{} `json:"data"`
}
