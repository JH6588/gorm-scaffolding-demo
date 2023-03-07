package main

import (
	. "example/go_dbs/comm"
	"fmt"
)

func main() {
	InitDB()

	//搜年龄18到30岁间，以年龄排序，并且分页
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
