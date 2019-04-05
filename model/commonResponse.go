package model

type CommonResponse struct {
	Status   string		`json:"status"`
	Code     uint16		`json:"code"`
	Message  string		`json:"message"`
}
