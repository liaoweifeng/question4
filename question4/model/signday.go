package model

import "fmt"

// 更新数据
func Signday(username string, message string) bool {
	if message == "qiandao" {
		sqlStr := "update user set signday=signday+1 where username=?"
		ret, err := db.Exec(sqlStr,username)
		if err != nil {
			fmt.Printf("update failed, err:%v\n", err)
			return true
		}
		n, err := ret.RowsAffected() // 操作影响的行数
		if err != nil {
			fmt.Printf("get RowsAffected failed, err:%v\n", err)
			return true
		}
		fmt.Printf("update success, affected rows:%d\n", n)
	} else {
		return false
	}
	return true
}
