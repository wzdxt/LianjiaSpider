package main

import (
	"runtime"
	"github.com/wzdxt/lianjia-spider/spider"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	spider.TranvelXiaoquList()
	time.Sleep(31 * time.Hour)
}

