package main

import (
	"abigenTest/callToken"
)

func main() {
	callToken.Init()
	//查询余额
	callToken.QueryTest()
	//转账
	callToken.Transfer()

}
