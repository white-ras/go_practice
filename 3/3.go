package main

import "fmt"

type MyIntSlice []int

func contains(sl []int, i int) bool {
	for _, v := range sl {
		if i == v {
			return true
		}
	}
	return false
}

func (ms MyIntSlice) Unique() MyIntSlice{
	var unique_ms MyIntSlice

	for _, v := range ms {
		if !contains(unique_ms, v){
			unique_ms = append(unique_ms, v)
		}
	}
	return unique_ms
}

func main(){
	m := MyIntSlice{1, 2, 2, 3, 3, 3, 4, 5}
	fmt.Println(m.Unique()) // [1, 2, 3, 4, 5]
}
