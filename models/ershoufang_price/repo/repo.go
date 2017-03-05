package repo

import (
	"github.com/wzdxt/lianjia-spider/db"
	"github.com/wzdxt/lianjia-spider/models/ershoufang_price"
)

func executeSelect(where string, bindings... interface{}) []*ershoufang_price.ErshoufangPrice {
	rows, err := db.Instance().Query("select id, ershoufang_id, xiaoqu_id, price, unit_price, prev_id from ershoufang_price " + where, bindings...)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var ret []*ershoufang_price.ErshoufangPrice = nil
	for rows.Next() {
		ershoufang := new(ershoufang_price.ErshoufangPrice)
		rows.Columns()
		rows.Scan(
			&ershoufang.Id,
			&ershoufang.ErshoufangId,
			&ershoufang.XiaoquId,
			&ershoufang.Price,
			&ershoufang.UnitPrice,
			&ershoufang.PrevId,
		)
		ret = append(ret, ershoufang)
	}
	return ret
}

func Get(id string) *ershoufang_price.ErshoufangPrice {
	return executeSelect("where id=?", id)[0]
}

func Save(price *ershoufang_price.ErshoufangPrice) *ershoufang_price.ErshoufangPrice {
	res, err := db.Instance().Exec("insert into ershoufang_price (ershoufang_id, xiaoqu_id, price, unit_price, prev_id) values(?,?,?,?,?)",
		price.ErshoufangId, price.XiaoquId, price.Price, price.UnitPrice, price.PrevId, )
	if err != nil {
		panic(err)
	}
	price.Id, _ = res.LastInsertId()
	return price
}

func New(ershoufangId, xiaoquId int64, price int, unitPrice int) *ershoufang_price.ErshoufangPrice {
	return &ershoufang_price.ErshoufangPrice{
		ErshoufangId:ershoufangId,
		XiaoquId:xiaoquId,
		Price:price,
		UnitPrice:unitPrice,
	}
}

func Create(ershoufangId, xiaoquId int64, price int, unitPrice int) *ershoufang_price.ErshoufangPrice {
	inst := New(ershoufangId, xiaoquId, price, unitPrice)
	return Save(inst)
}
