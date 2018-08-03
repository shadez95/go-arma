package arma

import (
	"os/exec"
	"strings"

	"github.com/Sirupsen/logrus"
)

var configs = "configs"

var log = logrus.New()

// Platform type represents what platform the server will run on
type Platform string

const (
	// Linux is of type Platform and should be used when running in pure Linux
	Linux Platform = "linux"
	// Windows is of type Platform and used when running in Windows
	Windows Platform = "windows"
	// Wine is of type Platform and should be used when running in linux but using wnie
	Wine Platform = "wine"
)

// BaseConfig is base struct for server and headless clients
type BaseConfig struct {
	// Path is the path to the ARMA directory. Not the path to the executable
	Path string
	// Platform is either linux, windows, or wine
	Platform Platform
	// Port the server will run on
	Port     string
	NoLogs   bool
	EnableHT bool
	Profiles string
	Mod      string
}

// Server is a struct for an ARMA server
type Server struct {
	// Admins is a list of strings of steamID's
	Admins              []string
	AutoInit            bool
	BasicConfig         string
	BEPath              string
	LoadMissionToMemory bool
	ProfileName         string
	ServerConfig        string
	ServerMod           string
	*BaseConfig
}

// HeadlessClient is a struct that will represent a headless client
type HeadlessClient struct {
	Connect     string
	Password    string
	ProfileName string
	*BaseConfig
}

// NewServer starts a new server instance with defaults set. Only Path and Platform are required in BaseConfig struct
func NewServer(server *Server) *Server {

	// Set defaults for BaseConfig portion of the struct
	server.Port = "2302"
	server.EnableHT = true
	server.Profiles = "profiles"

	// Set defaults for Server portion of the struct
	server.BasicConfig = "Default_basic.cfg"
	server.LoadMissionToMemory = true
	server.ServerConfig = "Default_server.cfg"

	return server
}

// NewHeadlessClient starts a new instance of HeadlessClient with defaults set. ProfileName are required
func NewHeadlessClient(hc *HeadlessClient) *HeadlessClient {
	hc.Connect = "127.0.0.1"
	hc.Port = "2302"
	return hc
}

// Start an Arma server
func (s *Server) Start() ([]byte, error) {

	var args []string
	var armaExecutable string

	switch s.Platform {
	case Windows:
		armaExecutable = strings.Join([]string{s.Path, "arma3server.exe"}, "\\")
	case Wine:
		armaExecutable = strings.Join([]string{s.Path, "arma3server.exe"}, "/")
	case Linux:
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
