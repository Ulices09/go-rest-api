package app

type ErrorField struct {
	Field   string `json:"field"`
	Code    string `json:"code"`
	Message string `json:"message"`
}
