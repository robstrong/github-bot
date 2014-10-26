package main

import (
	"github.com/op/go-logging"
)

type App struct {
	Logger *logging.Logger
	Config Config
}
