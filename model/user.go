package model

type DepositoMature struct {
	ProcessID         string `json:"processID"`
	ProcessCode       string `json:"processCode"`
	RequestTimestamp  string `json:"requestTimestamp"`
	FixedDepositAccNo string `json:"fixedDepositAccNo"`
	SavingAccNo       string `json:"savingAccNo"`
	Amount            string `json:"amount"`
}

type UserInfoRq struct {
	PartnerID        string `json:"partnerId" form:"partnerId" query:"partnerId"`
	RequestTimestamp string `json:"requestTimestamp" form:"requestTimestamp" query:"requestTimestamp"`
	Signature        string `json:"signature" form:"signature" query:"signature"`
	UserID           string `json:"userId" form:"userId" query:"userId"`
}

type UserInfoRp struct {
	ResponseCode      string         `json:"responseCode" form:"responseCode" query:"responseCode"`
	ResponseMessage   string         `json:"responseMessage" form:"responseMessage" query:"responseMessage"`
	ResponseTimestamp string         `json:"responseTimestamp" form:"responseTimestamp" query:"responseTimestamp"`
	Data              UserInfoDataRp `json:"data" form:"data" query:"data"`
}
type UserInfoDataRp struct {
	CifNo            string `json:"cifNo" form:"cifNo" query:"cifNo"`
	FullName         string `json:"fullName" form:"fullName" query:"fullName"`
	Alias            string `json:"alias" form:"alias" query:"alias"`
	LastLogin        string `json:"lastLogin" form:"lastLogin" query:"lastLogin"`
	DefaultCardNo    string `json:"defaultCardNo" form:"defaultCardNo" query:"defaultCardNo"`
	DefaultAccountNo string `json:"defaultAccountNo" form:"defaultAccountNo" query:"defaultAccountNo"`
	Cards            []Card `json:"cards" form:"cards" query:"cards"`
}

type Card struct {
	CardNo   string    `json:"cardNo" form:"cardNo" query:"cardNo"`
	CardSts  string    `json:"cardSts" form:"cardSts" query:"cardSts"`
	Accounts []Account `json:"accounts" form:"accounts" query:"accounts"`
}

type Account struct {
	AccountNo  string `json:"accountNo" form:"accountNo" query:"accountNo"`
	AccountSts string `json:"accountSts" form:"accountSts" query:"accountSts"`
}

type UserDataRq struct {
	PartnerID        string `json:"partnerId" form:"partnerId" query:"partnerId"`
	RequestTimestamp string `json:"requestTimestamp" form:"requestTimestamp" query:"requestTimestamp"`
	Signature        string `json:"signature" form:"signature" query:"signature"`
	UserID           string `json:"userId" form:"userId" query:"userId"`
}

type UserDataRp struct {
	ResponseCode      string `json:"responseCode" form:"responseCode" query:"responseCode"`
	ResponseMessage   string `json:"responseMessage" form:"responseMessage" query:"responseMessage"`
	ResponseTimestamp string `json:"responseTimestamp" form:"responseTimestamp" query:"responseTimestamp"`
	CifNo             string `json:"cifNo" form:"cifNo" query:"cifNo"`
	EktpID            string `json:"ektpID" form:"ektpID" query:"ektpID"`
	FullName          string `json:"fullName" form:"fullName" query:"fullName"`
	Alias             string `json:"alias" form:"alias" query:"alias"`
	MotherName        string `json:"motherName" form:"motherName" query:"motherName"`
	BirthDate         string `json:"birthDate" form:"birthDate" query:"birthDate"`
	Gender            string `json:"gender" form:"gender" query:"gender"`
	Address           string `json:"address" form:"address" query:"address"`
	Phone             string `json:"phone" form:"phone" query:"phone"`
	MobileNo          string `json:"mobileNo" form:"mobileNo" query:"mobileNo"`
	Email             string `json:"email" form:"email" query:"email"`
	Status            string `json:"status" form:"status" query:"status"`
}
