package main

import (
	. "example/go_dbs/comm"
	"fmt"
)

type People struct {
	ID    uint32 `gorm:"privateKey,column:id" json:"id,omitempty"`
	Name  string `gorm:"column:name" json:"name,omitempty"`
	Age   uint8  `gorm:"column:age" json:"age,omitempty"`
	Email string `gorm:"column:email" json:"email,omitempty"`
}

func GetPeopleTable() string {
	return "people"
}

type PersonQuery struct {
	AgeStart uint8 `sqlK:"age >"`
	AgeEnd   uint  `sqlK:"age <"`
}

// 前端查询条件
type PersonReq struct {
	PersonQuery
	OrderInfo
	PageInfo
}

// 查询请求
func (p *PersonReq) FindPeople() (pl []People) {
	sqlMap := GetSqlBuildMapFromStruct(p)
	cond, val, _ := WhereBuild(sqlMap)
	p.GetPageSql(p.GetOrderSql(DBX.Table(GetPeopleTable()).Where(cond, val...))).Find(&pl)
	return
}

func main() {
	InitDB()

	pr := PersonReq{
		PersonQuery: PersonQuery{
			AgeStart: 18,
			AgeEnd:   30,
		},
		OrderInfo: OrderInfo{
			SortKey:  "age",
			SortType: 0,
		},
		PageInfo: PageInfo{
			Page: 1,
			Size: 3,
		},
	}

	fmt.Println(pr.FindPeople())
}
