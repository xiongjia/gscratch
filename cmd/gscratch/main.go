package main

import (
	"fmt"
	gscratch "github.com/xiongjia/gscratch"
)

var (
	log = gscratch.LogCreate()
)

func main() {
	fmt.Println("gscratch")
	gscratch.GTest()

	log.Infof("Log test %s", "abc")
}
