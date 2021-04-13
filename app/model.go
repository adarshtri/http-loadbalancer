package app

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Server struct {
	ServerID uint `json:"server_id" gorm:"primary_key"`
	ServerHost string `json:"server_host"`
	ServerPort uint `json:"server_port"`
	ServerEndpoints []ServerEndpoint `json:"server_endpoints" gorm:"foreignkey:ServerID"`
}

type ServerEndpoint struct {
	ServerEndPointID uint `json:"server_endpoint_id" gorm:"primary_key"`
	EndpointName string `json:"endpoint_name"`
	EndPointMethod string `json:"endpoint_method"`
	ServerID uint `json:"-"`
}
