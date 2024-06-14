package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	users := []User{
		{"tarou", 33},
		{"zirou", 22},
		{"itirou", 11},
	}

	for i, _ := range users {
		users[i].Age = 44
	}

	fmt.Printf("%v", users) // どうなる？

	/* 別解 スライスをポインタの構造体にする
		users := []*User{
		{"tarou", 33},
		{"zirou", 22},
		{"itirou", 11},
	}
	*/

}