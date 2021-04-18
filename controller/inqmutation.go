package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/bns-engineering/mambusrc2dest/helper"
	"github.com/bns-engineering/mambusrc2dest/logging"
	"github.com/bns-engineering/mambusrc2dest/model"
	"github.com/bns-engineering/mambusrc2dest/tcommon/config"
	"github.com/bns-engineering/mambusrc2dest/tcommon/lookup"
	service "github.com/bns-engineering/mambusrc2dest/tservice"
	"google.golang.org/grpc"
)

//Inqmutation Function
func Inqmutation(w http.ResponseWriter, r *http.Request) {
	var p model.InquiryMutationRq
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
		strlog := fmt.Sprintf("Request > [%s] [%s]\n%v", "InqMutation", p.Signature, p)
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
					rc, inqReply = doProcessInquiryMutation(p)
				}
			}
		}
	}
	w, strReply = getHTTPReply(w, rc, inqReply)
	strlog := fmt.Sprintf("Reply   < [%s] [%s] [RC=%s]\n%v", "Inqmutation", p.Signature, rc, string(strReply))
	logging.Log(strlog)
	fmt.Fprintf(w, string(strReply))
	return
}

//DoProcessInqbalance function
func doProcessInquiryMutation(in model.InquiryMutationRq) (string, model.InquiryMutationRp) {
	var r0 = "06"
	var r1 model.InquiryMutationRp
	Accno := in.AccountNo

	// strlog := fmt.Sprintf("Request > %s[%s] [%s]", strings.Split(runtime.FuncForPC(reflect.ValueOf(doProcessInquiryMutation).Pointer()).Name(), ".")[2], Accno, in.Signature)
	// logging.Log(strlog)

	if (strings.Trim(Accno, " ")) == "" {
		r0 = "30"
		// strlog = fmt.Sprintf("Reply   < [RC=%s] [%s]", r0, in.Signature)
		// logging.Log(strlog)
		return r0, r1
	}

	// var callInqBalanceRPC AccBalanceService
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

	response, err := c.InqMutationRq(context.Background(), &service.InquiryMutationRq{
		AccountNo:        Accno,
		Signature:        in.Signature,
		PartnerID:        in.PartnerID,
		RequestTimestamp: in.RequestTimestamp,
		StartDate:        in.StartDate,
		EndDate:          in.EndDate})

	if err != nil {
		strlog := fmt.Sprintf("***ERROR: %s [%s]", err, in.Signature)
		logging.LogError(strlog)
		r0 = "S91"
		// strlog = fmt.Sprintf("Reply   < [RC=%s] [%s]", r0, in.Signature)
		// logging.Log(strlog)
		return r0, r1
	}

	r0 = response.GetResponseCode()
	dataCount, err := strconv.ParseInt(response.GetDataCount(), 10, 8)
	r1.ResponseCode = r0
	r1.ResponseMessage = lookup.GetResponseDesc(r1.ResponseCode)
	if r0 == "00" {
		r1.ResponseTimestamp = response.GetResponseTimestamp()
		r1.Data.AccountNo = response.GetAccountNo()
		r1.Data.AccountName = response.GetAccountName()
		r1.Data.AccountCurrency = response.GetAccountCurrency()
		r1.Data.StartDate = response.GetStartDate()
		r1.Data.EndDate = response.GetEndDate()
		r1.Data.DataCount = dataCount
		r1.Data.List = make([]model.DetailMutation, 0)
		if len(response.Data) > 0 {
			for _, value := range response.Data {
				r1.Data.List = append(r1.Data.List, model.DetailMutation{
					TransactionDate:   value.TransactionDate,
					TransactionTime:   value.TransactionTime,
					DCFlag:            value.DCFlag,
					TransactionAmount: value.TransactionAmount,
					TransactionName:   value.TransactionName,
					TransactionInfo:   value.TransactionInfo,
				})
			}
		}
	}

	// strlog = fmt.Sprintf("Reply   < [RC=%s] %v [%s]", r0, r1, in.Signature)
	// logging.Log(strlog)
	return r0, r1
}
