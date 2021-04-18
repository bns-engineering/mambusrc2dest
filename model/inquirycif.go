package model

type InquiryCIFRp struct {
	ResponseCode      string `json:"responseCode" form:"responseCode" query:"responseCode"`
	ResponseMessage   string `json:"responseMessage" form:"responseMessage" query:"responseMessage"`
	ResponseTimestamp string `json:"responseTimestamp" form:"responseTimestamp" query:"responseTimestamp"`
	CIFNo             string `json:"cifNo"  form:"cifNo" query:"cifNo"`
	NIK               string `json:"nik"  form:"NIK" query:"nik"`
	Fullname          string `json:"fullName" form:"fullName" query:"fullName"`
	Alias             string `json:"alias" form:"alias" query:"alias"`
	Birthdate         string `json:"birthDate" form:"birthDate" query:"birthDate"`
	Gender            string `json:"gender" form:"gender" query:"gender"`
	Address           string `json:"address" form:"address" query:"address"`
	Phone             string `json:"phone" form:"phone" query:"phone"`
	MobilePhone       string `json:"mobilePhone" form:"mobilePhone" query:"mobilePhone"`
	Email             string `json:"email" form:"email" query:"email"`
	MotherName        string `json:"motherName" form:"motherName" query:"motherName"`
	Status            string `json:"status" form:"status" query:"status"`
}

type InquiryCIFRq struct {
	PartnerID        string `json:"partnerId" form:"partnerId" query:"partnerId"`
	RequestTimestamp string `json:"requestTimestamp" form:"requestTimestamp" query:"requestTimestamp"`
	Signature        string `json:"signature" form:"signature" query:"signature"`
	CIFNo            string `json:"cifNo"  form:"cifNo" query:"cifNo"`
}
