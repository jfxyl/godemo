package main

import "godemo/log/demo0/logger"

func main() {
	logger.SetFile("demo0")
	logger.SetLevel(logger.DebugLevel)

	logger.Debug("hello world")
	logger.Info("hello world")
}
