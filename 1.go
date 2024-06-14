package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := []any{1, "2", 10, "11"}
	var err error

	for _, v := range s {
		// 出力用
		var i int

		switch v := v.(type) {
		case int:
			i = v
		case string:
			i, err = strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
			}
		default:
			fmt.Println("unknown!!!!!")
		}
		fmt.Printf("%02d\n", i)
	}
}
