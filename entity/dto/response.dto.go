package entity

type ResponseDto struct {
	Code  string      `json:"code"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}
