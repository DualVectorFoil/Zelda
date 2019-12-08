package model

type RegisterRespModel struct {
	IsSuccess bool   `json:"is_success"`
	Error     string `json:"error"`
}
