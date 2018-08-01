package arma

import (
	"runtime"

	"github.com/Sirupsen/logrus"
)

var configs = "configs"
var platform string

var log = logrus.New()

// Server is main object to interact with an Arma server
type Server struct {
	Path string
}

func getPlatform() {

	switch runtime.GOOS {
	case "windows":
		platform = "windows"
	case "linux":
		platform = "linux"
	}

	if platform == "" {
		log.Fatal("Your operating system doesn't support Arma")
	}
}

// StartServer starts an Arma server
func (s *Server) StartServer() error {

	return nil
}
