package main

import "fmt"

type MyInt int

func (m MyInt) String() string {
	return "hoge"
}

func main() {
	var m MyInt = 3
	fmt.Println(m) // hogeと出力させるように修正せよ。ただしmain関数に変更を加えないこと。
}