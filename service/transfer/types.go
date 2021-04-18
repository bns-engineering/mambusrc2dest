package transfer

//InqtrfOnUsRp - model response inqury transfer on us (internal partner)
type InqtrfOnUsRp struct {
	ResponseTimestamp string `json:"responseTimestamp" form:"responseTimestamp" query:"responseTimestamp"`
	DestAccountNo     string `json:"destAccountNo" form:"destAccountNo" query:"destAccountNo"`
	DestAccountName   string `json:"destAccountName" form:"destAccountName" query:"destAccountName"`
	DestAccountCurr   string `json:"destAccountCurr" form:"destAccountCurr" query:"destAccountCurr"`
	DestAccountType   string `json:"destAccountType" form:"destAccountType" query:"destAccountType"`
}

//InqtrfOffUsRp - model response inqury transfer off us (external partner)
type InqtrfOffUsRp struct {
	ResponseTimestamp string `json:"responseTimestamp" form:"responseTimestamp" query:"responseTimestamp"`
	DestInstutionID   string `json:"destinstutionId" form:"destinstutionId" query:"destinstutionId"`
	DestInstutionName string `json:"destinstutionName" form:"destinstutionName" query:"destinstutionName"`
	DestAccountNo     string `json:"destAccountNo" form:"destAccountNo" query:"destAccountNo"`
	DestAccountName   string `json:"destAccountName" form:"destAccountName" query:"destAccountName"`
	DestAccountCurr   string `json:"destAccountCurr" form:"destAccountCurr" query:"destAccountCurr"`
	DestAccountType   string `json:"destAccountType" form:"destAccountType" query:"destAccountType"`
}

//TrfOnUsRq - model request param for transfer on us (internal partner)
type TrfOnUsRq struct {
	PartnerID        string
	RequestTimestamp string
	Signature        string
	TransactionID    string
	AccountNo        string
	CurrencyCode     string
	DestAccountNo    string
	Amount           string
	Remark           string
}

//TrfOnUsRp - model response transfer on us (internal partner)
type TrfOnUsRp struct {
	ResponseTimestamp string `json:"responseTimestamp" form:"responseTimestamp" query:"responseTimestamp"`
	RetReferenceID    string `json:"retReferenceId" form:"retReferenceId" query:"retReferenceId"`
}

//TrfOffUsRq - model request param for transfer off us (external partner)
type TrfOffUsRq struct {
	PartnerID        string `json:"partnerId" form:"partnerId" query:"partnerId"`
	RequestTimestamp string `json:"requestTimestamp" form:"requestTimestamp" query:"requestTimestamp"`
	Signature        string `json:"signature" form:"signature" query:"signature"`
	TransactionID    string `json:"transactionId" form:"transactionId" query:"transactionId"`
	ReferenceID      string `json:"referenceId" form:"referenceid" query:"referenceid"`
	AccountNo        string `json:"accountNo"  form:"accountNo" query:"accountNo"`
	CurrencyCode     string `json:"currencyCode" form:"currencyCode" query:"currencyCode"`
	DestInstutionID  string `json:"destinstutionId" form:"destinstutionId" query:"destinstutionId"`
	DestAccountNo    string `json:"destAccountNo"  form:"destAccountNo" query:"destAccountNo"`
	Amount           string `json:"amount"  form:"amount" query:"amount"`
	Remark           string `json:"remark"  form:"remark" query:"remark"`
}

//TrfOffUsRp - model response transfer off us (external partner)
type TrfOffUsRp struct {
	ResponseTimestamp string `json:"responseTimestamp" form:"responseTimestamp" query:"responseTimestamp"`
	RetReferenceID    string `json:"retReferenceID" form:"retReferenceID" query:"retReferenceID"`
}
