package repo

import (
	"github.com/wzdxt/lianjia-spider/db"
	"github.com/wzdxt/lianjia-spider/models/xiaoqu"
)

func executeSelect(where string, bindings... interface{}) []*xiaoqu.Xiaoqu {
	rows, err := db.DBInstance().Query("select id, page_id, name, number from xiaoqu " + where, bindings...)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var ret []*xiaoqu.Xiaoqu = nil
	for rows.Next() {
		xiaoqu := new(xiaoqu.Xiaoqu)
		rows.Columns()
		rows.Scan(
			&xiaoqu.Id,
			&xiaoqu.PageId,
			&xiaoqu.Name,
			&xiaoqu.Number,
		)
		ret = append(ret, xiaoqu)
	}
	return ret
}

func Get(id string) *xiaoqu.Xiaoqu {
	return executeSelect("where id=?", id)[0]
}

func GetByPageId(pageId string) *xiaoqu.Xiaoqu {
	data := executeSelect("where page_id=?", pageId)
	if len(data) > 0 {
		return data[0]
	} else {
		return nil
	}
}

func All() []*xiaoqu.Xiaoqu {
	return executeSelect("")
}

func AllWithErshoufang() []*xiaoqu.Xiaoqu {
	return executeSelect("where number>0")
}

func WithErshoufangBatchAfter(id int64) []*xiaoqu.Xiaoqu {
	return executeSelect("where number>0 and id>? order by id limit 1000", id)
}

func Save(xiaoqu *xiaoqu.Xiaoqu) *xiaoqu.Xiaoqu {
	if xiaoqu.Id != 0 {
		Update(xiaoqu)
		return xiaoqu
	}
	res, err := db.DBInstance().Exec("insert into xiaoqu (page_id, name, number) values(?,?,?)",
		xiaoqu.PageId, xiaoqu.Name, xiaoqu.Number)
	if err != nil {
		panic(err)
	}
	xiaoqu.Id, _ = res.LastInsertId()
	return xiaoqu
}

func Update(xiaoqu *xiaoqu.Xiaoqu) int64 {
	res, err := db.DBInstance().Exec("update xiaoqu set page_id=?, name=?, number=? where id=?",
		xiaoqu.PageId, xiaoqu.Name, xiaoqu.Number, xiaoqu.Id, )
	if err != nil {
		panic(err)
	}
	ret, _ := res.RowsAffected()
	return ret
}

func New(pageId, name string, number int) *xiaoqu.Xiaoqu {
	return &xiaoqu.Xiaoqu{
		Name:name,
		PageId:pageId,
		Number:number,
	}
}

func Create(pageId, name string, number int) *xiaoqu.Xiaoqu {
	inst := New(pageId, name, number)
	return Save(inst)
}
