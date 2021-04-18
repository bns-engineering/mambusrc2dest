package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bns-engineering/mambusrc2dest/helper"
	"github.com/bns-engineering/mambusrc2dest/logging"
	"github.com/bns-engineering/mambusrc2dest/model"
	"github.com/bns-engineering/mambusrc2dest/tcommon/config"
	"github.com/bns-engineering/mambusrc2dest/tcommon/lookup"

	"strings"

	service "github.com/bns-engineering/mambusrc2dest/tservice"
	"google.golang.org/grpc"
)

//InqBalance Function
func InqBalance(w http.ResponseWriter, r *http.Request) {
	var p model.InquiryBalanceRq
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
		strlog := fmt.Sprintf("Request > [%s] [%s]\n%v", "InqBalance", p.Signature, p)
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
					rc, inqReply = doProcessInqbalance(p)
				}
			}
		}
	}
	w, strReply = getHTTPReply(w, rc, inqReply)
	strlog := fmt.Sprintf("Reply   < [%s] [%s] [RC=%s]\n%v", "InqBalance", p.Signature, rc, string(strReply))
	logging.Log(strlog)
	fmt.Fprintf(w, string(strReply))
	return
}

//DoProcessInqbalance function
func doProcessInqbalance(in model.InquiryBalanceRq) (string, model.InquiryBalanceRp) {
	var r0 = "06"
	var r1 model.InquiryBalanceRp
	Accno := in.AccountNo

	// strlog := fmt.Sprintf("Request > %s[%s] [%s]", strings.Split(runtime.FuncForPC(reflect.ValueOf(doProcessInqbalance).Pointer()).Name(), ".")[2], Accno, in.Signature)
	// logging.Log(strlog)

	if (strings.Trim(Accno, " ")) == "" {
		r0 = "30"
		// strlog = fmt.Sprintf("Reply   < [RC=%s] [%s]", r0, in.Signature)
		// logging.Log(strlog)
		return r0, r1
	}

	// var callInqBalanceRPC InqBalanceService
	grpcAddr := config.Config.Server.InquiryAddress
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		strlog := fmt.Sprintf("***ERROR: [%s] %s [%s]", grpcAddr, err, in.Signature)
		logging.LogError(strlog)
		r0 = "S91"
		// strlog = fmt.Sprintf("Reply   < [RC=%s] [%s]", r0, in.Signature)
		// logging.Log(strlog)
		return r0, r1
	}
	defer conn.Close()

	c := service.NewAccountHandlerClient(conn)

	response, err := c.InqbalanceRq(context.Background(), &service.InquiryBalanceRq{AccountNo: Accno, Signature: in.Signature, PartnerID: in.PartnerID, RqTimestamp: in.RequestTimestamp})
	if err != nil {
		strlog := fmt.Sprintf("***ERROR: %s [%s]", err, in.Signature)
		logging.LogError(strlog)
		r0 = "S91"
		// strlog = fmt.Sprintf("Reply   < [RC=%s] [%s]", r0, in.Signature)
		// logging.Log(strlog)
		return r0, r1
	}

	r0 = response.GetResponsecode()
	r1.ResponseCode = r0
	r1.ResponseMessage = lookup.GetResponseDesc(r1.ResponseCode)
	if r0 == "00" {
		r1.ResponseTimestamp = response.GetResponsetimestamp()
		r1.Data.AccountNo = response.GetAccountno()
		r1.Data.AccountName = response.GetAccountname()
		r1.Data.AccountBalance = response.GetAccountbalance()
		r1.Data.AccountCurrency = response.GetAccountcurrency()

	}

	// strlog = fmt.Sprintf("Reply   < [RC=%s] %v [%s]", r0, r1, in.Signature)
	// logging.Log(strlog)
	return r0, r1
}
