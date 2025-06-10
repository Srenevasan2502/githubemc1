package main

import "fmt"
// m := map[int]int{1:1,2:2,3:3}

func main() {
	m := map[int]int{1:1,2:2,3:3}
	val,ok := m[1]
	for k,v:=range m{
		fmt.Printf("%v:%v\n",k,v)
	}
	fmt.Println(val,ok)
}
