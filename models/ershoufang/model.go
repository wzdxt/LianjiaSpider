package ershoufang

import (
	"github.com/wzdxt/lianjia-spider/models/xiaoqu"
	"sync"
	"fmt"
	"github.com/wzdxt/lianjia-spider/models/ershoufang_price"
	ershoufang_price_repo "github.com/wzdxt/lianjia-spider/models/ershoufang_price/repo"
	"time"
)

type Ershoufang struct {
	Id           int64
	PageId       string
	Name         string
	Size         float64
	XiaoquPageId string
	SoldDate     *time.Time

	xiaoqu       *xiaoqu.Xiaoqu
	xiaoquOnce   sync.Once
}

func (this *Ershoufang) GetXiaoqu() *xiaoqu.Xiaoqu {
	this.xiaoquOnce.Do(func() {
	})
	return this.xiaoqu
}

func (this *Ershoufang) GetUrl() string {
	return fmt.Sprintf("http://sh.lianjia.com/ershoufang/%s.html", this.PageId)
}

func (this *Ershoufang) GetLastPrice() *ershoufang_price.ErshoufangPrice {
	return ershoufang_price_repo.GetLastByErshoufang(this.Id)
}

func (this *Ershoufang) String() string {
	return fmt.Sprint([]interface{}{this.Id, this.PageId, this.Name, this.Size, this.XiaoquPageId, this.SoldDate})
}

func Save() {}
