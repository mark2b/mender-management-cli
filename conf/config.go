package conf

import (
	"mender-management-cli/log"
	"os"
	"path/filepath"
)

const (
	VerboseFlag  = "verbose"
	DebugFlag    = "debug"
	UserFlag     = "user"
	PasswordFlag = "password"
	EndpointFlag = "endpoint"
)

type config struct {
	Verbose  bool
	Debug    bool
	ExecDir  string
	User     string
	Password string
	Endpoint string
}

var Config = new(config)

func (self *config) Init() {
	self.init()
}

func (self *config) init() {
	self.ExecDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	if self.Verbose {
		log.SetVerboseMode()
	}
	if self.Debug {
		log.SetDebugMode()
	}
}
