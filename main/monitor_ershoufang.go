package main

import (
	"runtime"
	"github.com/wzdxt/lianjia-spider/monitor"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	monitor.MonitorAllErshoufang()
}

