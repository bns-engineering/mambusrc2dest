package model

type VerifySessionIdRp struct {
	ResponseCode    string `json:"responseCode" form:"responseCode" query:"responseCode"`
	ResponseMessage string `json:"responseMessage" form:"responseMessage" query:"responseMessage"`
}

type VerifySessionIdRq struct {
	PartnerID        string `json:"partnerId" form:"partnerId" query:"partnerId"`
	RequestTimestamp string `json:"requestTimestamp" form:"requestTimestamp" query:"requestTimestamp"`
	Signature        string `json:"signature" form:"signature" query:"signature"`
	SessionID        string `json:"sessionId" form:"transactionId" query:"sessionId"`
}
