package monitor

import (
	ershoufang_repo "github.com/wzdxt/lianjia-spider/models/ershoufang/repo"
	ershoufang_queue "github.com/wzdxt/lianjia-spider/monitor/queue/ershoufang"
	xiaoqu_repo "github.com/wzdxt/lianjia-spider/models/xiaoqu/repo"
	xiaoqu_queue "github.com/wzdxt/lianjia-spider/monitor/queue/xiaoqu"
)

func MonitorAllErshoufang() {
	ershoufangs := ershoufang_repo.Unsold()
	for _, ershoufang := range ershoufangs {
		ershoufang_queue.Add(ershoufang)
	}
	c := ershoufang_queue.Process()
	<- c
}

func RefreshAllXiaoqu() {
	xiaoqus := xiaoqu_repo.All()
	for _, xiaoqu := range xiaoqus {
		xiaoqu_queue.Add(xiaoqu)
	}
	c := xiaoqu_queue.Process()
	<- c
}


