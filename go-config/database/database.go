package database

import (
	"errors"
	"fmt"
	"time"

	"github.com/GuIDeZaK/ecom-app/go-config/utils"
)

type Options struct {
	Host            string
	Port            int
	UserName        string
	Password        string
	DatabaseName    string
	Protocol        string
	ConnMaxLifetime time.Duration
	MaxIdleConns    int
	MaxOpenConns    int
	PARAM           string
}

type Database interface {
	Open(option Options)
	Get() interface{}
	Close()
	Ping() error
}

func BuildDns(options Options) (string, error) {
	handleError := func(msg string) (string, error) { return "", errors.New(msg) }
	if utils.IsBlank(options.UserName) {
		return handleError("user name cannot be empty")
	}
	if utils.IsBlank(options.Password) {
		return handleError("password cannot be empty")
	}
	if utils.IsBlank(options.Host) {
		return handleError("host cannot be empty")
	}
	if options.Port <= 0 {
		return handleError("port cannot be 0 or negative")
	}
	if utils.IsBlank(options.DatabaseName) {
		return handleError("database name cannot be empty")
	}
	var protocol string
	fmt.Println(protocol)
	if utils.IsBlank(options.Protocol) {
		protocol = "tcp"
	} else {
		protocol = options.Protocol
	}
	var param string
	if utils.IsBlank(options.PARAM) {
		param = "parseTime=true"
	} else {
		param = options.PARAM
	}
	fmt.Println(param)
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d ", options.UserName, options.Password, options.DatabaseName, options.Host, options.Port), nil
}
