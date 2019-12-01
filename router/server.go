package router

import (
	"net"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

var serverConfig *ServerConfig

type ServerConfig struct {
	IpAddress string
}

func init() {
	Router = SetupRouter()
}

func StartServer() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				serverConfig = &ServerConfig{
					IpAddress: ipnet.IP.String(),
				}
			}
		}
	}

	Router.Run()
}

func GetServerConfig() *ServerConfig {
	return serverConfig
}