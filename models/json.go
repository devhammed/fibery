package models

type JsonResponse struct {
	Ok      bool        `json:"ok"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type JsonError struct {
	Ok      bool          `json:"ok"`
	Message string        `json:"message,omitempty"`
	Errors  []interface{} `json:"errors,omitempty"`
}
