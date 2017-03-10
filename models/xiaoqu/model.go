package xiaoqu

import "fmt"

type Xiaoqu struct {
	Id      int64
	PageId  string
	Name    string
	Number  int
	Qu      string
	Bankuai string
}

func (this *Xiaoqu) GetUrl() string {
	return fmt.Sprintf("http://sh.lianjia.com/xiaoqu/%s.html", this.PageId)
}

func (this *Xiaoqu) GetTravelUrl(page int) string {
	return fmt.Sprintf("http://sh.lianjia.com/ershoufang/d%dq%s", page, this.PageId)
}
