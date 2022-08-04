package main

import "logrus/logrus_log"

func main() {

	logger := logrus_log.GetLogger()
	logger.Info("created Log file and save log")
	logger.Error("No error here")
	logger.Debug("No Debug mode")
	logger.Fatal("No Fatal mode")
	logger.Panic("No Panic mode")

}
