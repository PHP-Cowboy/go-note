package main

import "fmt"

func main() {
	defer a()
	defer b()
	defer fmt.Println("ccc")
	defer fmt.Println("ddd")
	return
}

func a() {
	fmt.Println("aaa")
}

func b() {
	fmt.Println("bbb")
}
