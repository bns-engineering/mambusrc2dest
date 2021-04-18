package model

type NotifTransactionRp struct {
	ResponseCode    string `json:"responseCode" form:"responseCode" query:"responseCode"`
	ResponseMessage string `json:"responseMessage" form:"responseMessage" query:"responseMessage"`
}

type NotifTransactionRq struct {
	PartnerID            string `json:"partnerId" form:"partnerId" query:"partnerId"`
	TransactionType      string `json:"transactionType" form:"transactionType" query:"transactionType"`
	ProductCode          string `json:"productCode" form:"productCode" query:"productCode"`
	TransactionTimestamp string `json:"transactionTimestamp" form:"transactionTimestamp" query:"transactionTimestamp"`
	FromAccountNo        string `json:"fromAccountNo" form:"fromAccountNo" query:"fromAccountNo"`
	DestAccountNo        string `json:"destAccountNo" form:"destAccountNo" query:"destAccountNo"`
	FromBankCode         string `json:"fromBankCode"  form:"fromBankCode" query:"fromBankCode"`
	DestBankCode         string `json:"destBankCode" form:"destBankCode" query:"destBankCode"`
	ReferenceNo          string `json:"referenceNo"  form:"referenceNo" query:"referenceNo"`
	Amount               string `json:"amount"  form:"amount" query:"amount"`
	TransactionResponse  string `json:"transactionResponse"  form:"transactionResponse" query:"transactionResponse"`
}
