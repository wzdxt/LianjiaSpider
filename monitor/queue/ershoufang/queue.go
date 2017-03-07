package ershoufang

import (
	"github.com/wzdxt/lianjia-spider/models/ershoufang"
	ershoufang_repo "github.com/wzdxt/lianjia-spider/models/ershoufang/repo"
	ershoufang_price_repo "github.com/wzdxt/lianjia-spider/models/ershoufang_price/repo"
	"log"
	"github.com/wzdxt/lianjia-spider/inspector"
	"time"
	"sync"
)

var finishChannel = make(chan struct{})

var ershoufangQueue = make([]*ershoufang.Ershoufang, 0, 1000)

func Add(ershoufang *ershoufang.Ershoufang) {
	ershoufangQueue = append(ershoufangQueue, ershoufang)
	log.Printf("added to ershoufang pool: %s", ershoufang)
}

func Process() chan struct{} {
	go run()
	return finishChannel
}

func run() {
	wg := sync.WaitGroup{}
	wg.Add(len(ershoufangQueue))
	for len(ershoufangQueue) > 0 {
		ershoufang := ershoufangQueue[0]
		log.Printf("start process %s", ershoufang)
		ershoufangQueue = ershoufangQueue[1:]
		go checkErshoufang(ershoufang, &wg)
	}
	wg.Wait()
	finishChannel <- struct{}{}
}

var n = 10;
var limit = make(chan struct{}, n)

func init() {
	for i := 0; i < n; i++ {
		limit <- struct{}{}
	}
}

func checkErshoufang(ershoufang *ershoufang.Ershoufang, wg *sync.WaitGroup) {
	<-limit
	defer func() {
		limit <- struct{}{}
	}()
	house, price := inspector.InspectErshoufangFromUrl(ershoufang.GetUrl())
	house.Id = ershoufang.Id
	log.Printf("get ershoufang(%#v) price(%#v)", house, price)
	ershoufang_repo.Save(house)
	price.ErshoufangId = house.Id
	lastPrice := ershoufang.GetLastPrice()

	if lastPrice == nil {
		ershoufang_price_repo.Save(price)
	} else if price.Price != lastPrice.Price {
		price.PrevId = lastPrice.Id
		ershoufang_price_repo.Save(price)
	} else {
		price = lastPrice
	}
	wg.Done()
	time.Sleep(1 * time.Second)
}


