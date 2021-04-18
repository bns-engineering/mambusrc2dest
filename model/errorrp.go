package model

type ErrorRp struct {
	ResponseCode    string `json:"responseCode" form:"responseCode" query:"responseCode"`
	ResponseMessage string `json:"responseMessage" form:"responseMessage" query:"responseMessage"`
}
