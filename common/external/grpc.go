package external

import (
	param "github.com/bns-engineering/grpc/param/client"
	"github.com/bns-engineering/mambusrc2dest/common/config"
	"github.com/bns-engineering/mambusrc2dest/common/logging"
)

var (
	//ParamClient - grpc param client
	ParamClient param.IClient
)

//InitGrpcExternal - initialize client for external grpc
func InitGrpcExternal() {
	paramOption := param.Options{
		Address: config.Config.Server.ParamAddress,
	}

	res, err := param.GetClient(&paramOption)
	if err != nil {
		logging.Error(err, "error get client param")
		return
	}

	ParamClient = res
}
