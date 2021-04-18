package model

type DetailMutation struct {
	TransactionDate   string `json:"transactionDate"  form:"transactionDate" query:"transactionDate"`
	TransactionTime   string `json:"transactionTime"  form:"transactionTime" query:"transactionTime"`
	DCFlag            string `json:"dcFlag"  form:"dcFlag" query:"dcFag"`
	TransactionAmount string `json:"transactionAmount"  form:"transactionAmount" query:"transactionAmount"`
	TransactionName   string `json:"transactionName"  form:"transactionName" query:"transactionName"`
	TransactionInfo   string `json:"transactionInfo"  form:"transactionInfo" query:"transactionInfo"`
}
