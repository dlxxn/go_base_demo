package main

import (
	"fmt"
	high_level_code "go_base_demo/high.level.code"
)

func main() {
	//fmt.Println("test")
	//test.Test()
	//log.KlogAndError()
	//conf.YamlTest()

	fmt.Println("construct==========")
	high_level_code.Construct()
	fmt.Println("extend===========")
	high_level_code.Extend()

	//service.MsgService

}
