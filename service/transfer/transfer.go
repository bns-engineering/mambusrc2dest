package transfer

import (
	"context"
	"strings"

	"github.com/bns-engineering/mambusrc2dest/service"
	"github.com/bns-engineering/mambusrc2dest/tcommon/config"
	"github.com/bns-engineering/mambusrc2dest/tcommon/logging"
	"github.com/bns-engineering/mambusrc2dest/tcommon/lookup"
	"google.golang.org/grpc"
)

//ProcessInqTrfOnUs function get inquery transfer on us (internal call)
func ProcessInqTrfOnUs(partnerID, requestTimestamp, signature, destAccountNo string) (result InqtrfOnUsRp, responseCode string, err error) {
	//init logger for logging
	logger := logging.SetVariable(map[string]interface{}{
		"partnerID":        partnerID,
		"requestTimestamp": requestTimestamp,
		"signature":        signature,
		"destAccountNo":    destAccountNo,
	})

	responseCode = lookup.GENERALERROR
	result = InqtrfOnUsRp{}

	if (strings.Trim(destAccountNo, " ")) == "" {
		logger.Error(err, "DestAccountNo is empty")
		responseCode = lookup.FORMATERROR
		return
	}

	grpcAddr := config.Config.Server.TransferAddress
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		logger.Error(err, "error dial up grpc")
		responseCode = lookup.SERVICEERROR
		return
	}
	defer conn.Close()

	c := service.NewTransferHandlerClient(conn)

	response, err := c.InqTransferOnUsRq(context.Background(), &service.InqTrfOnUsRq{
		PartnerID:        partnerID,
		RequestTimestamp: requestTimestamp,
		Signature:        signature,
		DestAccountNo:    destAccountNo})
	if err != nil {
		logger.Error(err, "error inquiry transfer on us from grpc")
	}

	responseCode = response.GetResponseCode()
	if responseCode == lookup.SUCCESS {
		result.ResponseTimestamp = response.GetResponseTimestamp()
		result.DestAccountNo = response.GetDestAccountNo()
		result.DestAccountName = response.GetDestAccountName()
		result.DestAccountCurr = response.GetDestAccountCurr()
		result.DestAccountType = response.GetDestAccountType()
	}

	return
}

//ProcessInqTrfOffUs - function get inquery transfer off us (external call)
func ProcessInqTrfOffUs(partnerID, requestTimestamp, signature, destInstutionID, destAccountNo string) (result InqtrfOffUsRp, responseCode string, err error) {
	//init logger for logging
	logger := logging.SetVariable(map[string]interface{}{
		"partnerID":        partnerID,
		"requestTimestamp": requestTimestamp,
		"signature":        signature,
		"destInstutionID":  destInstutionID,
		"destAccountNo":    destAccountNo,
	})

	responseCode = lookup.GENERALERROR
	result = InqtrfOffUsRp{}

	if (strings.Trim(destAccountNo, " ")) == "" {
		logger.Error(err, "DestAccountNo is empty")
		responseCode = lookup.FORMATERROR
		return
	}

	grpcAddr := config.Config.Server.TransferAddress
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		logger.Error(err, "error dial up grpc")
		responseCode = lookup.SERVICEERROR
		return
	}
	defer conn.Close()

	c := service.NewTransferHandlerClient(conn)

	response, err := c.InqTransferOffUsRq(context.Background(), &service.InqTrfOffUsRq{
		PartnerID:        partnerID,
		RequestTimestamp: requestTimestamp,
		Signature:        signature,
		DestInstutionID:  destInstutionID,
		DestAccountNo:    destAccountNo,
	})
	if err != nil {
		logger.Error(err, "error inquiry transfer off us from grpc")
	}

	responseCode = response.GetResponseCode()

	if responseCode == lookup.SUCCESS {
		result.ResponseTimestamp = response.GetResponseTimestamp()
		result.DestAccountNo = response.GetDestAccountNo()
		result.DestInstutionID = response.GetDestInstutionID()
		result.DestInstutionName = response.GetDestInstutionName()
		result.DestAccountNo = response.GetDestAccountNo()
		result.DestAccountName = response.GetDestAccountName()
		result.DestAccountCurr = response.GetDestAccountCurr()
		result.DestAccountType = response.GetDestAccountType()
	}

	return
}

//ProcessTrfOnUs - func for transfer on us
func ProcessTrfOnUs(param TrfOnUsRq) (result TrfOnUsRp, responseCode string, err error) {
	//init logger for logging
	logger := logging.SetVariable(map[string]interface{}{
		"param": param,
	})

	responseCode = lookup.GENERALERROR
	result = TrfOnUsRp{}

	if (strings.Trim(param.AccountNo, " ")) == "" {
		logger.Error(err, "AccountNo is empty")
		responseCode = lookup.FORMATERROR
		return
	}

	grpcAddr := config.Config.Server.TransferAddress
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		logger.Error(err, "error dial up grpc")
		responseCode = lookup.SERVICEERROR
		return
	}
	defer conn.Close()

	c := service.NewTransferHandlerClient(conn)

	response, err := c.TransferOnUsRq(context.Background(), &service.TrfOnUsRq{
		PartnerID:        param.PartnerID,
		RequestTimestamp: param.RequestTimestamp,
		Signature:        param.Signature,
		TransactionID:    param.TransactionID,
		AccountNo:        param.AccountNo,
		CurrencyCode:     param.CurrencyCode,
		DestAccountNo:    param.DestAccountNo,
		Amount:           param.Amount,
		Remark:           param.Remark,
	})

	if err != nil {
		logger.Error(err, "error transfer on us from grpc")
	}

	responseCode = response.GetResponseCode()

	if responseCode == lookup.SUCCESS {
		result.ResponseTimestamp = response.GetResponseTimestamp()
		result.RetReferenceID = response.GetRetReferenceID()
	}

	return
}

//ProcessTrfOffUs - func for transfer off us
func ProcessTrfOffUs(param TrfOffUsRq) (result TrfOffUsRp, responseCode string, err error) {
	//init logger for logging
	logger := logging.SetVariable(map[string]interface{}{
		"param": param,
	})

	responseCode = lookup.GENERALERROR
	result = TrfOffUsRp{}

	if (strings.Trim(param.AccountNo, " ")) == "" {
		logger.Error(err, "AccountNo is empty")
		responseCode = lookup.FORMATERROR
		return
	}

	grpcAddr := config.Config.Server.TransferAddress
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		logger.Error(err, "error dial up grpc")
		responseCode = lookup.SERVICEERROR
		return
	}
	defer conn.Close()

	c := service.NewTransferHandlerClient(conn)

	response, err := c.TransferOffUsRq(context.Background(), &service.TrfOffUsRq{
		PartnerID:        param.PartnerID,
		RequestTimestamp: param.RequestTimestamp,
		Signature:        param.Signature,
		TransactionID:    param.TransactionID,
		ReferenceID:      param.ReferenceID,
		AccountNo:        param.AccountNo,
		CurrencyCode:     param.CurrencyCode,
		DestInstutionID:  param.DestInstutionID,
		DestAccountNo:    param.DestAccountNo,
		Amount:           param.Amount,
		Remark:           param.Remark,
	})

	if err != nil {
		logger.Error(err, "error transfer off us from grpc")
	}

	responseCode = response.GetResponsecode()
	if responseCode == lookup.SUCCESS {
		result.ResponseTimestamp = response.GetResponsetimestamp()
		result.RetReferenceID = response.GetRetReferenceID()
	}

	return
}
