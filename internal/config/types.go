package config

import "errors"

var (
	ErrObjectNotFound     = errors.New("object not found")
	ErrOrderStatusUnknown = errors.New("got unknown order status")
	ErrInvalidInput       = errors.New("invalid input")
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	SSLMode  string
}
