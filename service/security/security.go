package security

import (
	"context"

	"github.com/bns-engineering/mambusrc2des/service"
	"github.com/bns-engineering/mambusrc2des/tcommon/config"
	"github.com/bns-engineering/mambusrc2des/tcommon/logging"
	"github.com/bns-engineering/mambusrc2des/tcommon/lookup"
	"google.golang.org/grpc"
)

//GeneratePartnerKey - get generated parner key
func GeneratePartnerKey(userID, partnerID, requestTimestamp, signature, publicKey string) (publicKeyRes, responseCode string, err error) {
	//init logger for logging
	logger := logging.SetVariable(map[string]interface{}{
		"userID":           userID,
		"partnerID":        partnerID,
		"requestTimestamp": requestTimestamp,
		"signature":        signature,
		"publicKey":        publicKey,
	})

	responseCode = lookup.GENERALERROR

	grpcAddr := config.Config.Server.SecurityAddress
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		logger.Error(err, "error dial up grpc")
		responseCode = lookup.SERVICEERROR
		return
	}
	defer conn.Close()

	c := service.NewSecurityHandlerClient(conn)

	response, err := c.GeneratePartnerKey(context.Background(), &service.GenPartnerKeyRq{
		UserID:           userID,
		PartnerID:        partnerID,
		RequestTimestamp: requestTimestamp,
		Signature:        signature,
		PublicKey:        publicKey,
	})
	if err != nil {
		logger.Error(err, "error generate partner key from grpc")
	}

	responseCode = response.GetResponseCode()
	if responseCode == lookup.SUCCESS {
		publicKeyRes = response.GetPublicKey()
	}

	return
}

//GeneratePartnerOnboardingKey - get generated onboarding parner key
func GeneratePartnerOnboardingKey(partnerID, requestTimestamp, signature string) (publicKeyRes, responseCode string, err error) {
	//init logger for logging
	logger := logging.SetVariable(map[string]interface{}{
		"partnerID":        partnerID,
		"requestTimestamp": requestTimestamp,
		"signature":        signature,
	})

	responseCode = lookup.GENERALERROR
	grpcAddr := config.Config.Server.SecurityAddress
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		logger.Error(err, "error dial up grpc")
		responseCode = lookup.SERVICEERROR
		return
	}
	defer conn.Close()

	c := service.NewSecurityHandlerClient(conn)
	response, err := c.GeneratePartnerOnboardingKey(context.Background(), &service.GenPartnerOnboardKeyRq{
		PartnerID:        partnerID,
		RequestTimestamp: requestTimestamp,
		Signature:        signature,
	})
	if err != nil {
		logger.Error(err, "error generate partner onboarding key from grpc")
	}

	responseCode = response.GetResponseCode()
	if responseCode == lookup.SUCCESS {
		publicKeyRes = response.GetPublicKey()
	}

	return
}

//Login - login function
func Login(userID, partnerID, requestTimestamp, signature, encryptValue string) (sessionID, responseCode string, err error) {
	//init logger for logging
	logger := logging.SetVariable(map[string]interface{}{
		"userID":           userID,
		"partnerID":        partnerID,
		"requestTimestamp": requestTimestamp,
		"signature":        signature,
		"encryptValue":     encryptValue,
	})

	responseCode = lookup.GENERALERROR

	grpcAddr := config.Config.Server.SecurityAddress
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		logger.Error(err, "error dial up grpc")
		responseCode = lookup.SERVICEERROR
		return
	}
	defer conn.Close()

	c := service.NewSecurityHandlerClient(conn)

	response, err := c.Login(context.Background(), &service.LoginRq{
		UserID:           userID,
		PartnerID:        partnerID,
		RequestTimestamp: requestTimestamp,
		Signature:        signature,
		EncryptValue:     encryptValue,
	})
	if err != nil {
		logger.Error(err, "error login from grpc")
	}

	responseCode = response.GetResponseCode()
	if responseCode == lookup.SUCCESS {
		sessionID = response.GetSessionID()
	}

	return
}
