package main

import (
	gscratch "github.com/xiongjia/gscratch"
	logger "github.com/xiongjia/gscratch/logger"
)

var (
	log = logger.New("gscratch")
)

func init() {
	logger.SetLevel(logger.DebugLevel)
}

func main() {
	log.Debugf("loading gscratch")
	gscratch.GTest()
}
