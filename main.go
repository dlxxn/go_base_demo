package main

import (
	"fmt"
	"go_base_demo/examples"
)

func main() {
	//fmt.Println("test")
	//test.Test()
	//log.KlogAndError()
	//conf.YamlTest()

	//test interface
	//fmt.Println("test interface===========")
	//examples.Measure(examples.GetRect())

	//test rang array map
	//fmt.Println("test array map rang===========")
	//examples.TestRange()

	//test goroutine
	//fmt.Println("test goroutine=============")
	//examples.TestGoroutines()

	//test channels
	fmt.Println("test channels=============")
	//examples.TestChannels()

	//done := make(chan bool, 1)
	/**
	Start a worker goroutine, giving it the channel to notify on. Block until we receive a notification from
	the worker on the channel.
	*/
	//go examples.Worker(done)
	/**
	If you removed the <- done line from this program, the program would exit before the worker even
	started.
	*/
	//<-done

	//examples.TestPingPong()
	examples.TestSelect()

	/*	fmt.Println("construct==========")
		high_level_code.Construct()
		fmt.Println("extend===========")
		high_level_code.Extend()*/

	//service.MsgService

}
