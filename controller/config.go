package controller

import "github.com/bns-engineering/mambusrc2des/model"

//Config Variable for Configuration file
var Config *model.Configuration

//SetConfig function
func SetConfig(cfg *model.Configuration) {
	Config = cfg
}
