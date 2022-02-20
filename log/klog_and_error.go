package log

import (
	"errors"
	"flag"
	"fmt"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
)

func KlogAndError() {
	/**
	在 go 1.13 版本发布之后，针对 fmt.Errorf 增加了 wraping 功能，并在 errors 包中添加了 Is() 和 As() 函数
	*/

	fmt.Errorf("test fmt error is %v", "testFmt流派")
	e1 := errors.New("error1")
	efmt := fmt.Errorf("error2: %w", e1)
	fmt.Println(efmt)
	fmt.Println(errors.Is(efmt, e1))

	klog.InitFlags(nil)
	klog.Errorf("test is %v", "test")

	//输出到文件
	// By default klog writes to stderr. Setting logtostderr to false makes klog
	// write to a log file.
	flag.Set("logtostderr", "false") //日志输出到stderr，不输出到日志文件。false为关闭
	flag.Set("log_file", "myfile.log")
	flag.Parse()
	klog.Info("nice to meet you")
	klog.Flush()

	klog.InitFlags(nil)
	flag.Set("v", "3")
	flag.Parse()
	log := klogr.New().WithName("MyName").WithValues("user", "you")
	log.Info("hello", "val1", 1, "val2", map[string]int{"k": 1})
	log.V(1).Info("nice to meet you")
	log.Error(nil, "uh oh", "trouble", true, "reasons", []float64{0.1, 0.11, 3.14})
	klog.Flush()
}
