package main

import (
	"runtime"
	"github.com/wzdxt/lianjia-spider/spider"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	spider.FollowChengjiao()
}


