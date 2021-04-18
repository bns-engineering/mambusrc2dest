package middleware

import (
	"net/http"
)

type (
	//MuxHandler is func type to be parsed to http router handler
	MuxHandler func(http.ResponseWriter, *http.Request)
	//Chain is func type chain before throws to main handler
	Chain func(MuxHandler) MuxHandler
	//Handler is func type to standarize main handler function
	Handler func(http.ResponseWriter, *http.Request) (HandlerResponse, error)

	//HandlerResponse is standarized response type for all api handler
	HandlerResponse struct {
		Data         interface{}
		ResponseCode string
	}
)

// APIResponseTemplate is template response API BANK NET SYARIAH
type APIResponseTemplate struct {
	ResponseCode    string      `json:"responseCode" form:"responseCode" query:"responseCode"`
	ResponseMessage string      `json:"responseMessage" form:"responseMessage" query:"responseMessage"`
	Data            interface{} `json:"data"`
}
