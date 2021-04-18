package model

type InquiryMutationRp struct {
	ResponseCode      string                `json:"responseCode" form:"responseCode" query:"responseCode"`
	ResponseMessage   string                `json:"responseMessage" form:"responseMessage" query:"responseMessage"`
	ResponseTimestamp string                `json:"responseTimestamp" form:"responseTimestamp" query:"responseTimestamp"`
	Data              InquiryMutationDataRp `json:"data" form:"data" query:"data"`
}

type InquiryMutationDataRp struct {
	AccountNo       string           `json:"accountNo"  form:"accountNo" query:"accountNo"`
	AccountName     string           `json:"accountName" form:"accountName" query:"accountName"`
	AccountCurrency string           `json:"accountCurrency" form:"accountCurrency" query:"accountCurrency"`
	StartDate       string           `json:"startDate"  form:"startDate" query:"startDate"`
	EndDate         string           `json:"endDate"  form:"endDate" query:"endDate"`
	DataCount       int64            `json:"dataCount"  form:"dataCount" query:"dataCount"`
	List            []DetailMutation `json:"list" form:"list" query:"list"`
}

type InquiryMutationRq struct {
	PartnerID        string `json:"partnerId" form:"partnerId" query:"partnerId"`
	RequestTimestamp string `json:"requestTimestamp" form:"requestTimestamp" query:"requestTimestamp"`
	Signature        string `json:"signature" form:"signature" query:"signature"`
	AccountNo        string `json:"accountNo"  form:"accountNo" query:"accountNo"`
	StartDate        string `json:"startDate"  form:"startDate" query:"startDate"`
	EndDate          string `json:"endDate"  form:"endDate" query:"endDate"`
}
