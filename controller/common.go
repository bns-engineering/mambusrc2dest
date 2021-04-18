package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bns-engineering/mambusrc2dest/common/helper"
	"github.com/bns-engineering/mambusrc2dest/common/lookup"
)

//getHTTPReply function
func getHTTPReply(pw http.ResponseWriter, rc string, modreply interface{}) (http.ResponseWriter, []byte) {

	if rc != "00" {
		modreply = helper.DoReplyError(rc)
	}

	strReply, err := json.Marshal(modreply)
	if err != nil {
		fmt.Println(err)
	}

	pw.Header().Set("Content-Type", "application/json")
	httpstatus, err := strconv.Atoi(lookup.GetResponseHTTP(rc))
	if err != nil {
		httpstatus = 500
	}
	pw.WriteHeader(httpstatus)

	return pw, strReply
}
