package main

import (
	"fmt"
	gscratch "github.com/xiongjia/gscratch"
	logger "github.com/xiongjia/gscratch/logger"
)

var (
	log = logger.New("gscratch")
)

func init() {
	fmt.Println("main init")
}

func main() {
	// logLevel := flag.String("logLevel", "info", "Log level (debbug, error, info)")
	// flag.Parse()
	// fmt.Printf("%s\n", *logLevel)
	fmt.Println("gscratch")
	gscratch.GTest()
	log.Infof("test %d", 1)
}
