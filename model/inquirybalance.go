package model

type InquiryBalanceRp struct {
	ResponseCode      string               `json:"responseCode" form:"responseCode" query:"responseCode"`
	ResponseMessage   string               `json:"responseMessage" form:"responseMessage" query:"responseMessage"`
	ResponseTimestamp string               `json:"responseTimestamp" form:"responseTimestamp" query:"responseTimestamp"`
	Data              InquiryBalanceDataRp `json:"data" form:"data" query:"data"`
}

type InquiryBalanceDataRp struct {
	AccountNo       string `json:"accountNo"  form:"accountNo" query:"accountNo"`
	AccountName     string `json:"accountName" form:"accountName" query:"accountName"`
	AccountCurrency string `json:"accountCurrency" form:"accountCurrency" query:"accountCurrency"`
	AccountBalance  string `json:"accountBalance" form:"accountBalance" query:"accountBalance"`
}

type InquiryBalanceRq struct {
	PartnerID        string `json:"partnerId" form:"partnerId" query:"partnerId"`
	RequestTimestamp string `json:"requestTimestamp" form:"requestTimestamp" query:"requestTimestamp"`
	Signature        string `json:"signature" form:"signature" query:"signature"`
	AccountNo        string `json:"accountNo"  form:"accountNo" query:"accountNo"`
}
