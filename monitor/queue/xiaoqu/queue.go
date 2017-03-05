package xiaoqu

import (
	"log"
	"github.com/wzdxt/lianjia-spider/inspector"
	"github.com/wzdxt/lianjia-spider/models/xiaoqu"
	"github.com/wzdxt/lianjia-spider/models/xiaoqu/repo"
	"time"
)

var finishChannel = make(chan struct{})

var xiaoquQueue = make([]*xiaoqu.Xiaoqu, 0, 1000)

func Add(xiaoqu *xiaoqu.Xiaoqu) {
	xiaoquQueue = append(xiaoquQueue, xiaoqu)
	log.Printf("added to xiaoqu pool: %s", xiaoqu)
}

func Process() chan struct{} {
	go run()
	return finishChannel
}

var n = 10
var cc = make(chan chan struct{}, n)

func run() {
	for i := 0; i < n; i++ {
		cc <- make(chan struct{})
	}
	for len(xiaoquQueue) > 0 {
		c := <-cc
		xiaoqu := xiaoquQueue[0]
		log.Printf("start process %s", xiaoqu)
		xiaoquQueue = xiaoquQueue[1:]
		go checkXiaoqu(xiaoqu, c)
		time.Sleep(10 * time.Millisecond)
	}
	for i := 0; i < n; i++ {
		<-cc
	}
	finishChannel <- struct{}{}
}

func checkXiaoqu(xiaoqu *xiaoqu.Xiaoqu, c chan struct{}) {
	area := inspector.InspectXiaoquFromUrl(xiaoqu.GetUrl())
	area.Id = xiaoqu.Id
	repo.Update(area)
	cc <- c
}


