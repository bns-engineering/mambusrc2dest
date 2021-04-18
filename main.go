// File: main.go
package main

import (
	// "datafeeder/common/helper"
	// "datafeeder/common/logging"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bns-engineering/mambusrc2des/controller"
	cf "github.com/bns-engineering/mambusrc2des/tcommon/config"
	"github.com/bns-engineering/mambusrc2des/tcommon/helper"
	"github.com/bns-engineering/mambusrc2des/tcommon/logging"
	"github.com/gorilla/mux"
)

func main() {
	logging.OpenFileLog()

	logging.InfoLn(helper.GetStrDateTime())
	logging.InfoLn("Starting...")

	//load config
	Config := cf.LoadConfig()
	cf.SetConfig(&Config)

	//serve rest
	addrHTTP := Config.Server.Port
	go serveHTTP(addrHTTP)

	// external.InitGrpcExternal()

	// logging.Infof("Debug Mode: %v", Config.Debug)
	// logging.Infof("GRPC account  : \"%s\"", Config.Server.InquiryAddress)
	// logging.Infof("GRPC account  : \"%s\"", Config.Server.InquiryAddress)
	// logging.Infof("GRPC customer : \"%s\"", Config.Server.CustomerAddress)
	// logging.Infof("GRPC param    : \"%s\"", Config.Server.ParamAddress)
	// logging.Infof("GRPC security : \"%s\"", Config.Server.SecurityAddress)
	// logging.Infof("GRPC transfer : \"%s\"", Config.Server.TransferAddress)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
	log.Println("All server stopped!")
}

func serveHTTP(port string) error {
	router := mux.NewRouter()

	// router.HandleFunc("/param/listcity", middleware.ChainHandler(param.HandleCity)).Methods("POST")
	// router.HandleFunc("/param/listdegree", middleware.ChainHandler(param.HandleDegree)).Methods("POST")
	// router.HandleFunc("/param/listincome", middleware.ChainHandler(param.HandleIncome)).Methods("POST")
	// router.HandleFunc("/param/listindustrialsector", middleware.ChainHandler(param.HandleIndustrialSector)).Methods("POST")
	// router.HandleFunc("/param/listopeningpurpose", middleware.ChainHandler(param.HandleOpeningPurpose)).Methods("POST")
	// router.HandleFunc("/param/listprovince", middleware.ChainHandler(param.HandleProvince)).Methods("POST")
	// router.HandleFunc("/param/listsourceincome", middleware.ChainHandler(param.HandleSourceIncome)).Methods("POST")
	// router.HandleFunc("/param/listjobtypes", middleware.ChainHandler(param.HandleJobTypes)).Methods("POST")

	// router.HandleFunc("/security/generatepartnerkey", middleware.ChainHandler(security.HandleGeneratepartnerkey)).Methods("POST")
	// router.HandleFunc("/security/generatepartneronboardingkey", middleware.ChainHandler(security.HandleGeneratepartneronboardingkey)).Methods("POST")
	// router.HandleFunc("/security/login", middleware.ChainHandler(security.HandleLogin)).Methods("POST")

	// router.HandleFunc("/transfer/inqtrfonus", middleware.ChainHandler(transfer.HandleInqTrfOnUs)).Methods("POST")
	// router.HandleFunc("/transfer/inqtrfoffus", middleware.ChainHandler(transfer.HandleInqTrfOffUs)).Methods("POST")
	// router.HandleFunc("/transfer/trfonus", middleware.ChainHandler(transfer.HandleTrfOnUs)).Methods("POST")
	// router.HandleFunc("/transfer/trfoffus", middleware.ChainHandler(transfer.HandleTrfOffUs)).Methods("POST")

	// router.HandleFunc("/account/inqbalance", controller.InqBalance).Methods("POST")
	// router.HandleFunc("/account/inqcif", controller.Inqcif).Methods("POST")
	// router.HandleFunc("/account/inqmutation", controller.Inqmutation).Methods("POST")

	router.HandleFunc("/user/userinfo", controller.DepositoMature).Methods("POST")
	router.HandleFunc("/deposito/mature", controller.DepositoMature).Methods("POST")
	router.HandleFunc("/api/deposits/10092542497/withdrawal-transactions", controller.DepositoMature).Methods("POST")
	//{{server}}/api/deposits/10092542497/withdrawal-transactions

	addr := ":" + port

	logging.Infof("Http Server Started. Listening on port: %s", addr)
	err := http.ListenAndServe(addr, router)
	logging.Error(err, "failed listen and serve rest")

	return err
}
