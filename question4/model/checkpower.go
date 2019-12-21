package model

import (
	"fmt"
)

func Checkpower(username string) int {
	sqlStr := "select power from user where username=?"
	var power int
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, username).Scan(&power)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
	}
	fmt.Printf("%d",power)
	return power
}
