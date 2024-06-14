package main

import (
	"errors"
	"fmt"
)

type MyMap map[int]string

func findKeyByValue (m MyMap, s string) (int,error){
	var return_value int
	var err error

  for key,v := range m {
    if v == s {
      return_value = key
    }
  }

  if return_value == 0 {
    err = errors.New("Not Found!!!")
  }

	return return_value,err
}

func main() {
  m := MyMap{
    1: "01",
    2: "02",
    3: "03",
  }
  key1, err1 := findKeyByValue(m, "03") // key→3, err→nil
  key2, err2 := findKeyByValue(m, "05") // key→0にすること(初期値なので), errはある
	fmt.Println(key1,err1)
	fmt.Println(key2,err2)
}