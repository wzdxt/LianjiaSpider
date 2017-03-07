package monitor

import (
	ershoufang_repo "github.com/wzdxt/lianjia-spider/models/ershoufang/repo"
	ershoufang_queue "github.com/wzdxt/lianjia-spider/monitor/queue/ershoufang"
	xiaoqu_repo "github.com/wzdxt/lianjia-spider/models/xiaoqu/repo"
	xiaoqu_queue "github.com/wzdxt/lianjia-spider/monitor/queue/xiaoqu"
	"github.com/wzdxt/lianjia-spider/db"
)

func MonitorAllErshoufang() {
	lastId := db.GetLastErshoufangProcessId();
	ershoufangs := ershoufang_repo.UnsoldBatchAfter(lastId)
	var maxId int64 = 0
	for _, ershoufang := range ershoufangs {
		ershoufang_queue.Add(ershoufang)
		if ershoufang.Id > maxId {
			maxId = ershoufang.Id
		}
	}
	c := ershoufang_queue.Process()
	<- c
	db.SetLastErshoufangProcessId(maxId)
}

func RefreshAllXiaoqu() {
	xiaoqus := xiaoqu_repo.All()
	for _, xiaoqu := range xiaoqus {
		xiaoqu_queue.Add(xiaoqu)
	}
	c := xiaoqu_queue.Process()
	<- c
}


