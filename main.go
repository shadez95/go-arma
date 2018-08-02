package arma

import (
	"os/exec"
	"strings"

	"github.com/Sirupsen/logrus"
)

var configs = "configs"

var log = logrus.New()

// BaseServer is base struct for server and headless clients
type BaseServer struct {
	// Path is the path to the ARMA directory. Not the path to the executable
	Path string
	// Platform is either linux, windows, or wine
	Platform string
	// Port the server will run on
	Port        string
	NoLogs      bool
	EnableHT    bool
	ProfileName string
	Profiles    string
	Mod         string
}

// Server is a struct for an ARMA server
type Server struct {
	// Admins is a list of strings of steamID's
	Admins              []string
	AutoInit            bool
	BasicConfig         string
	BEPath              string
	LoadMissionToMemory bool
	ServerConfig        string
	ServerMod           string
	BaseServer
}

// HeadlessClient is a struct that will represent a headless client
type HeadlessClient struct {
	connect  string
	password string
	BaseServer
}

// NewServer should be used when you want to create a server
func NewServer() *Server {
	var admins = []string{}
	server := &Server{
		Admins: admins,
	}
	server.Port = "2302"
	server.NoLogs = false
	server.EnableHT = true

	// TODO: Contine defaults

	return server
}

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
	default:
		log.Error("Platform not specified, should be windows, wine, or linux")
	}

	armaServer := exec.Cmd{
		Path: armaExecutable,
		Args: args,
		Dir:  s.Path,
	}

	return armaServer.CombinedOutput()
}
