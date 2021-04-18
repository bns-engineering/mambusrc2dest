package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bns-engineering/mambusrc2dest/common/config"
	"github.com/bns-engineering/mambusrc2dest/common/lookup"
	"github.com/bns-engineering/mambusrc2dest/helper"
	"github.com/bns-engineering/mambusrc2dest/logging"
	"github.com/bns-engineering/mambusrc2dest/model"
	"github.com/bns-engineering/mambusrc2dest/tools"
)

//Inqcif Function
func Inqcif(w http.ResponseWriter, r *http.Request) {
	var p model.InquiryCIFRq
	var inqReply interface{}
	var strReply []byte
	var rc = "06"
	var sessionID = ""
	sessionID = r.Header.Get("SessionId")
	if sessionID == "" {
		strlog := fmt.Sprintf("%s [%s]", "SessionID empty", p.Signature)
		logging.LogError(strlog)
		rc = lookup.INVALID_SESSIONID
	}
	resultCheckSessionID := helper.CheckSessionID(sessionID, p.Signature, p.RequestTimestamp, p.PartnerID, config.Config.Server.SecurityAddress)
	if resultCheckSessionID != true {
		strlog := fmt.Sprintf("%s [%s]", "Invalid SessionID", sessionID)
		logging.LogError(strlog)
		rc = lookup.INVALID_SESSIONID
	}

	if rc != lookup.INVALID_SESSIONID {
		err := helper.DecodeJSONBody(w, r, &p)
		strlog := fmt.Sprintf("Request > [%s] [%s]\n%v", "Inqcif", p.Signature, p)
		logging.Log(strlog)

		if err != nil {
			strlog := fmt.Sprintf("%s [%s]\n%v", err.Error(), p.Signature, p)
			logging.LogError(strlog)
			rc = "30" //format error

		} else {
			rc = helper.DoCheckSignature(p.PartnerID, p.RequestTimestamp, p.Signature)

			if rc == "00" {
				rc = helper.DoCheckPartnerID(p.PartnerID)
				if rc == "00" {
					rc, inqReply = doProcessInquiryCIF(p)
				}
			}
		}
	}
	w, strReply = getHTTPReply(w, rc, inqReply)
	strlog := fmt.Sprintf("Reply   < [%s] [%s] [RC=%s]\n%v", "Inqcif", p.Signature, rc, string(strReply))
	logging.Log(strlog)
	fmt.Fprintf(w, string(strReply))
	return
}

//DoProcessInquiryCIF function
func doProcessInquiryCIF(in model.InquiryCIFRq) (string, model.InquiryCIFRp) {
	var r0 = "06"
	var r1 model.InquiryCIFRp
	CIFno := in.CIFNo

	if (strings.Trim(CIFno, " ")) == "" {
		r0 = "30"
		return r0, r1
	}

	r0 = "00"
	if r0 == "00" {
		r1.ResponseCode = r0
		r1.ResponseMessage = lookup.GetResponseDesc(r1.ResponseCode)
		r1.ResponseTimestamp = tools.GetStrDateTime()
		r1.CIFNo = CIFno
		r1.Fullname = "Gunung Semeru"
		r1.Alias = "Semeru"
		r1.Birthdate = "2000-01-05"
		r1.Gender = "L"
		r1.Address = "Jl. Pendakian No 1, Kab.Malang, Jawa Timur. "
		r1.Phone = "02166667777"
		r1.MobilePhone = "08771111155555"
		r1.Email = "Semeru@contoh.com"
		r1.MotherName = "Bunda Semeru"
		r1.Status = "0"

	}

	return r0, r1
}
