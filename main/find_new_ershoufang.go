package main

import (
	"runtime"
	"github.com/wzdxt/lianjia-spider/spider"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	spider.TravelAllXiaoqu()
	time.Sleep(13 * time.Hour)
}

