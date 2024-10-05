package service

import "os"

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
)

var Sigch chan os.Signal = make(chan os.Signal)
