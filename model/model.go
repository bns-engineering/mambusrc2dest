package model

//Configuration Struct
type Configuration struct {
	Debug    bool
	Server   Server
	Context  Context
	Database Database
}

//Server Struct
type Server struct {
	Port            string
	InquiryAddress  string
	TransferAddress string
	SecurityAddress string
	UserAddress     string
	ParamAddress    string
	CustomerAddress string
}

//Context Struct
type Context struct {
	Timeout int
}

//Database Struct
type Database struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}
