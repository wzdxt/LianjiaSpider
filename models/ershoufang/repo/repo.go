package repo

import (
	"github.com/wzdxt/lianjia-spider/db"
	"github.com/wzdxt/lianjia-spider/models/ershoufang"
	"time"
)

func executeSelect(where string, bindings... interface{}) []*ershoufang.Ershoufang {
	rows, err := db.DBInstance().Query("select id, page_id, name, size, xiaoqu, bankuai, qu, xiaoqu_page_id, sold_date from ershoufang " + where, bindings...)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var ret []*ershoufang.Ershoufang = nil
	for rows.Next() {
		ershoufang := new(ershoufang.Ershoufang)
		rows.Columns()
		rows.Scan(
			&ershoufang.Id,
			&ershoufang.PageId,
			&ershoufang.Name,
			&ershoufang.Size,
			&ershoufang.Xiaoqu,
			&ershoufang.Bankuai,
			&ershoufang.Qu,
			&ershoufang.XiaoquPageId,
			ershoufang.SoldDate,
		)
		ret = append(ret, ershoufang)
	}
	return ret
}

func Get(id int64) *ershoufang.Ershoufang {
	return executeSelect("where id=?", id)[0]
}

func GetByPageId(pageId string) (ret *ershoufang.Ershoufang) {
	res := executeSelect("where page_id=?", pageId)
	if len(res) > 0 {
		ret = res[0]
	}
	return
}

func All() []*ershoufang.Ershoufang {
	return executeSelect("")
}

func Unsold() []*ershoufang.Ershoufang {
	return executeSelect("where sold_date is null")
}

func UnsoldBatchAfter(id int64) []*ershoufang.Ershoufang {
	return executeSelect("where sold_date is null and id>? order by id limit 1000", id)
}

func Save(ershoufang *ershoufang.Ershoufang) *ershoufang.Ershoufang {
	if ershoufang.Id != 0 {
		Update(ershoufang)
		return ershoufang
	}
	res, err := db.DBInstance().Exec("insert into ershoufang (page_id, name, size, xiaoqu, bankuai, qu, xiaoqu_page_id, sold_date) values(?,?,?,?,?,?,?,?)",
		ershoufang.PageId, ershoufang.Name, ershoufang.Size, ershoufang.Xiaoqu, ershoufang.Bankuai, ershoufang.Qu, ershoufang.XiaoquPageId, ershoufang.SoldDate)
	if err != nil {
		panic(err)
	}
	ershoufang.Id, _ = res.LastInsertId()
	return ershoufang
}

func Update(ershoufang *ershoufang.Ershoufang) int64 {
	res, err := db.DBInstance().Exec("update ershoufang set page_id=?, name=?, size=?, xiaoqu=?, bankuai=?, qu=?, xiaoqu_page_id=?, sold_date=? where id=?",
		ershoufang.PageId, ershoufang.Name, ershoufang.Size, ershoufang.Xiaoqu, ershoufang.Bankuai, ershoufang.Qu, ershoufang.XiaoquPageId, ershoufang.SoldDate, ershoufang.Id, )
	if err != nil {
		panic(err)
	}
	ret, _ := res.RowsAffected()
	return ret
}

func New(pageId, name string, size float64, xiaoqu, bankuai, qu, xiaoquPageId string, soldDate *time.Time) *ershoufang.Ershoufang {
	return &ershoufang.Ershoufang{
		PageId:       pageId,
		Name:            name,
		Size:            size,
		XiaoquPageId: xiaoquPageId,
		Xiaoqu: xiaoqu,
		Bankuai: bankuai,
		Qu: qu,
		SoldDate:soldDate,
	}
}

func Create(pageId, name string, size float64, xiaoqu, bankuai, qu, xiaoquPageId string, soldDate *time.Time) *ershoufang.Ershoufang {
	return Save(New(pageId, name, size, xiaoqu, bankuai, qu, xiaoquPageId, soldDate))
}


