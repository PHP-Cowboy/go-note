package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	count,amount := 5,500
	ret := make([]int,0,count)
	for i := 0; i < count; i++ {
		money := DoubleMean(count-i,amount)
		amount -= money
		ret = append(ret,money)
	}
	fmt.Println(ret)
}

var min = 1

func DoubleMean(count,amount int) int {
	if count == 1 {
		return amount
	}
	//计算最大可以金额
	max := amount - count * min
	//计算最大可用平均值
	avg := max / count
	//计算二倍均值 再加 最小值 防止出现 0 值 比如 count 和 amount 相等时
	avg2 := avg * 2 + min
	//随机种子
	rand.Seed(time.Now().UnixNano())
	//随机 最大为二倍均值
	return rand.Intn(avg2) + min
}