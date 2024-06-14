package main

import "fmt"

type MyIntSlice []int

func (ms MyIntSlice) Unique() MyIntSlice{
	var unique_ms MyIntSlice

	return unique_ms
}


func main(){
	m := MyIntSlice{1, 2, 2, 3, 3, 3, 4, 5}
	fmt.Println(m.Unique()) // [1, 2, 3, 4, 5]
}