package main

import (
	"runtime"
	"log"
	"github.com/wzdxt/lianjia-spider/monitor"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Println("--------- start ------------")
	//monitor.RefreshAllXiaoqu()
	log.Println("--------- 0 ------------")
	//spider.TranvelXiaoquList()
	log.Println("--------- 1 ------------")
	//spider.TravelAllXiaoqu()
	log.Println("--------- 2 ------------")
	monitor.MonitorAllErshoufang()
	log.Println("--------- 3 ------------")
}

