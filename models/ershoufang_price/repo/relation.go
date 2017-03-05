package repo

import "github.com/wzdxt/lianjia-spider/models/ershoufang_price"

func GetLastByErshoufang(ershoufangId int64) *ershoufang_price.ErshoufangPrice {
	res := executeSelect("where ershoufang_id=? order by id desc limit 1", ershoufangId)
	if len(res) == 0 {
		return nil
	} else {
		return res[0]
	}
}
