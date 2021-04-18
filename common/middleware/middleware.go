package middleware

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bns-engineering/mambusrc2dest/common/logging"
	"github.com/bns-engineering/mambusrc2dest/common/lookup"
)

//ChainHandler is middleware function to manage flow chain filters and main handler function
func ChainHandler(endHandler Handler, chains ...Chain) MuxHandler {
	if len(chains) == 0 {
		return handlerWrapper(endHandler)
	}

	return chains[0](ChainHandler(endHandler, chains[1:]...))
}

//handlerWrapper is global wrapper handler function, converting from handler return type to http response, tracking, and so on
var handlerWrapper = func(h Handler) MuxHandler {
	return func(pw http.ResponseWriter, r *http.Request) {
		//start handler
		resp, err := h(pw, r)

		var data = resp.Data
		if err != nil || resp.ResponseCode != "00" {
			data = nil
		}

		standardAPIResponse(pw, data, resp.ResponseCode)
	}
}

// standardAPIResponse writes header with standard response API BANK NET SYARIAH
var standardAPIResponse = func(pw http.ResponseWriter, modreply interface{}, responseCode string) {
	var response APIResponseTemplate

	response.Data = modreply
	response.ResponseCode = responseCode
	response.ResponseMessage = lookup.GetResponseDesc(responseCode)

	strReply, err := json.Marshal(response)
	if err != nil {
		logging.Error(err, "Failed marshal response API")
	}

	pw.Header().Set("Content-Type", "application/json")
	if pw.Header().Get("Access-Control-Allow-Origin") == "" {
		pw.Header().Set("Access-Control-Allow-Origin", "*")
	}

	httpstatus, err := strconv.Atoi(lookup.GetResponseHTTP(responseCode))
	if err != nil {
		httpstatus = 500
	}
	pw.WriteHeader(httpstatus)

	pw.Write(strReply)
}
