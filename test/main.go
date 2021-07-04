package main

import (
	"fmt"
)

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
)

const (
	a = iota + 1
	b = iota + 2
	c
	d
	e
	f
)

func main() {
	//fmt.Println(math.Sqrt(24*24 + 32*32))
	fmt.Println(B, KB, MB, GB, TB)
	fmt.Println(a, b, c, d, e, f)
	return
}
