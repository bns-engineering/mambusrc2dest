package controller

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/bns-engineering/mambusrc2dest/common/config"
	"github.com/bns-engineering/mambusrc2dest/common/helper"
	"github.com/bns-engineering/mambusrc2dest/common/lookup"
	"github.com/bns-engineering/mambusrc2dest/logging"
	"github.com/bns-engineering/mambusrc2dest/model"
	service "github.com/bns-engineering/mambusrc2dest/service"
	"github.com/bns-engineering/mambusrc2dest/util"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func Deposits(w http.ResponseWriter, r *http.Request) {
	// var p model.DepositoMature
	var p interface{}
	var resp interface{}
	var strReply []byte
	var rc = "06"
	params := mux.Vars(r)
	id := params["id"]
	category := params["category"]
	strIdempotency := r.Header.Get("Idempotency-Key")
	baseUrl := r.URL.Path

	if rc != lookup.INVALID_SESSIONID {
		err := helper.DecodeJSONBody(w, r, &p)
		logging.Log(baseUrl)
		strlog := fmt.Sprintf("%s Request > Deposits:[%s] [%s] [%s] ", helper.GetStrtimestamp(), category, id, strIdempotency)
		logging.Log(strlog)
		if err != nil {
			strlog := fmt.Sprintf("%s [%s]\n%v", err.Error(), id, strIdempotency)
			logging.LogError(strlog)
			rc = "30"
		} else {
			DoWithdrawal(w, r)
		}
	}
	w, strReply = getHTTPReply(w, rc, resp)
	// strlog := fmt.Sprintf("Reply   < [RC=%s]\n%v", rc, string(strReply))
	strlog := fmt.Sprintf("%s Reply   > Deposits:[%s] [%s] [%s] [RC=%s] ", helper.GetStrtimestamp(), category, id, strIdempotency, rc)
	logging.Log(strlog)
	fmt.Fprintf(w, string(strReply))
	return
}

func DoWithdrawal(w http.ResponseWriter, r *http.Request) {
	// requestBody, _ := json.Marshal("")
	requestBody := r.Body
	timeout := time.Duration(10 * time.Second)
	// baseUrl := config.LoadConfig().Mambu.Endpoint + `/api/gljournalentries?paginationDetails=OFF&from=` + pstrDate + `&to=` + pstrDate + `&offset=` + fmt.Sprint(noffset) + `&limit=` + fmt.Sprint(nlimit)
	baseUrl := r.URL
	fmt.Println(baseUrl)
	header := r.Header
	var hasil map[string]interface{}
	statusCode, responseBody, err := util.SendHTTP(http.MethodPost, baseUrl, int(timeout), &hasil, header, bytes.NewBuffer(requestBody))

	fmt.Println("Statuscode:", statusCode)
	if statusCode != 200 {
		strerror := hasil["errors"]
		strerror = strerror.([]interface{})[0]
		strerrordata := strerror.(map[string]interface{})
		fmt.Println("Statuscode:", strerrordata["errorCode"])
		fmt.Println("errorSource:", strerrordata["errorSource"])
		fmt.Println("errorReason:", strerrordata["errorReason"])
	}

	if err != nil {
		// hasilnya, _ := json.Marshal(err)
		// var strlog = fmt.Sprintf("Err: %v", string(hasilnya))
		// fmt.Println(strlog)
		logging.Log(responseBody)
	}

	if statusCode == 200 {
		logging.Log(responseBody)

	}
	return
}

//UserInfo Function for get User information
func DepositoMature(w http.ResponseWriter, r *http.Request) {
	// var p model.DepositoMature
	var p interface{}
	var resp interface{}
	var strReply []byte
	var rc = "06"
	// var sessionID = ""
	// sessionID = r.Header.Get("SessionId")
	// if sessionID == "" {
	// 	strlog := fmt.Sprintf("%s [%s]", "SessionID empty", p.Signature)
	// 	logging.LogError(strlog)
	// 	rc = lookup.INVALID_SESSIONID
	// }
	// resultCheckSessionID := helper.CheckSessionID(sessionID, p.Signature, p.RequestTimestamp, p.PartnerID, config.Config.Server.SecurityAddress)
	// if resultCheckSessionID != true {
	// 	strlog := fmt.Sprintf("%s [%s]", "Invalid SessionID", sessionID)
	// 	logging.LogError(strlog)
	// 	rc = lookup.INVALID_SESSIONID
	// }

	if rc != lookup.INVALID_SESSIONID {
		err := helper.DecodeJSONBody(w, r, &p)
		strlog := fmt.Sprintf("Request > [%s] ", p)
		// strlog := fmt.Sprintf("Request > [%s] [%s]\n%v", p.ProcessID, p.ProcessCode, p)
		logging.Log(strlog)
		if err != nil {
			// strlog := fmt.Sprintf("%s [%s]\n%v", err.Error(), p.Signature, p)
			// logging.LogError(strlog)
			rc = "30" //format error
		} else {
			// rc = helper.DoCheckSignature(p.PartnerID, p.RequestTimestamp, p.Signature)

			// if rc == "00" {
			// 	rc = helper.DoCheckPartnerID(p.PartnerID)
			// 	if rc == "00" {
			// 		rc, resp = getUserInfo(p)
			// 	}
			// }
		}
	}
	w, strReply = getHTTPReply(w, rc, resp)
	strlog := fmt.Sprintf("Reply   < [RC=%s]\n%v", rc, string(strReply))
	// strlog := fmt.Sprintf("Reply   < [%s] [%s] [RC=%s]\n%v", p.ProcessID, p.ProcessCode, rc, string(strReply))
	logging.Log(strlog)
	fmt.Fprintf(w, string(strReply))
	return
}

//DoProcessInqbalance function
func getUserInfo(in model.UserInfoRq) (string, model.UserInfoRp) {
	var r0 = "06"
	var r1 model.UserInfoRp
	grpcAddr := config.Config.Server.CustomerAddress
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		strlog := fmt.Sprintf("***ERROR: [%s] %s [%s]", grpcAddr, err, in.Signature)
		logging.LogError(strlog)
		r0 = "S91"
		return r0, r1
	}
	defer conn.Close()

	c := service.NewUserHandlerClient(conn)

	response, err := c.UserInfo(context.Background(), &service.UserInfoRq{
		UserID:           in.UserID,
		PartnerID:        in.PartnerID,
		RequestTimestamp: in.RequestTimestamp,
		Signature:        in.Signature,
	})

	if err != nil {
		strlog := fmt.Sprintf("***ERROR: %s [%s]", err, in.Signature)
		logging.LogError(strlog)
		r0 = "S91"
		return r0, r1
	}

	r0 = response.GetResponseCode()
	r1.ResponseCode = r0
	r1.ResponseMessage = lookup.GetResponseDesc(r1.ResponseCode)
	if r0 == "00" {
		r1.Data.CifNo = response.GetCifNo()
		r1.Data.FullName = response.GetFullName()
		r1.Data.Alias = response.GetAlias()
		r1.Data.LastLogin = response.GetLastLogin()
		r1.Data.DefaultCardNo = response.GetDefaultCardNo()
		r1.Data.DefaultAccountNo = response.GetDefaultAccountNo()
		r1.Data.Cards = make([]model.Card, 0)
		for _, value := range response.GetCards() {
			var card = model.Card{}
			card.CardNo = value.CardNo
			card.CardSts = value.CardSts
			card.Accounts = make([]model.Account, 0)
			for _, valueAcc := range value.GetAccounts() {
				var account = model.Account{}
				account.AccountNo = valueAcc.AccountNo
				account.AccountSts = valueAcc.AccountSts
				card.Accounts = append(card.Accounts, account)
			}
			r1.Data.Cards = append(r1.Data.Cards, card)
		}

	}
	return r0, r1
}
