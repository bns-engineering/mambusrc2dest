package helper

import (
	"context"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bns-engineering/mambusrc2dest/common/lookup"
	"github.com/bns-engineering/mambusrc2dest/logging"
	"github.com/bns-engineering/mambusrc2dest/model"
	"github.com/bns-engineering/mambusrc2dest/service"
	"github.com/golang/gddo/httputil/header"
	"google.golang.org/grpc"
)

type MalformedRequest struct {
	Status int
	Msg    string
}

func (mr *MalformedRequest) Error() string {
	return mr.Msg
}

func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			return &MalformedRequest{Status: http.StatusUnsupportedMediaType, Msg: msg}
		}
	}
	// rlog := r
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	// strlog := fmt.Sprintf("decodeJSONBody: %+v", dst)
	// logging.Log(strlog)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		msg := "Request body must only contain a single JSON object"
		return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
	}

	return nil
}

func DoReplyError(rc string) model.ErrorRp {
	var JError model.ErrorRp
	JError.ResponseCode = rc
	JError.ResponseMessage = lookup.GetResponseDesc(rc)
	return JError
}

//DoCheckSignature untuk pengecekan Signature.
func DoCheckSignature(PartnerID string, RqTimestamp string, Signature string) string {

	h := sha256.New()
	h.Write([]byte(PartnerID + "|" + RqTimestamp))

	sEnc := b64.StdEncoding.EncodeToString(h.Sum(nil))
	if sEnc == Signature {
		return "00"
	}

	return "63"
}

//DoCheckPartnerID pengecekan PartnerID
func DoCheckPartnerID(PartnerID string) string {
	r0 := "06"
	if PartnerID == "Partner-001" {
		r0 = "00"
	}

	return r0
}

func CheckSessionID(sessionID string, signature string, requestTimeStamp string, partnerID string, securityAddress string) bool {
	grpcAddr := securityAddress
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		strlog := fmt.Sprintf("***ERROR: [%s] %s [%s]", grpcAddr, err, signature)
		logging.LogError(strlog)
		return false
	}
	defer conn.Close()

	c := service.NewSecurityHandlerClient(conn)

	response, err := c.VerifySessionId(context.Background(), &service.VerifySessionIdRq{
		PartnerID:        partnerID,
		RequestTimestamp: requestTimeStamp,
		Signature:        signature,
		SessionID:        sessionID,
	})

	if err != nil {
		strlog := fmt.Sprintf("***ERROR: %s [%s]", err, signature)
		logging.LogError(strlog)
		return false
	}

	result := response.GetResponseCode()
	if result == lookup.SUCCESS {
		return true
	} else {
		return false
	}
}
