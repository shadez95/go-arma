package arma

import (
	"os/exec"
	"strings"

	"github.com/Sirupsen/logrus"
)

var configs = "configs"

var log = logrus.New()

// Server is main object to interact with an Arma server
type Server struct {
	// Path is the path to the ARMA directory. Not the path to the executable
	Path string

	// Admins is a list of strings of steamID's
	Admins []string

	// Platform is either linux, windows, or wine
	Platform string
}

// func getPlatform() {

// 	switch runtime.GOOS {
// 	case "windows":
// 		platform = "windows"
// 	case "linux":
// 		platform = "linux"
// 	}

// 	if platform == "" {
// 		log.Fatal("Your operating system doesn't support Arma")
// 	}

// 	return platform
// }

// Start an Arma server
func (s *Server) Start() ([]byte, error) {

	var args []string
	var armaExecutable string

	switch s.Platform {
	case "windows":
		armaExecutable = strings.Join([]string{s.Path, "arma3server.exe"}, "\\")
	case "wine":
		armaExecutable = strings.Join([]string{s.Path, "arma3server.exe"}, "/")
	case "linux":
		armaExecutable = strings.Join([]string{s.Path, "arma3server"}, "/")
	}

	armaServer := exec.Cmd{
		Path: armaExecutable,
		Args: args,
		Dir:  s.Path,
	}

	return armaServer.CombinedOutput()
}
