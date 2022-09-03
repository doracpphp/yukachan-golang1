package main

import "fmt"

func main() {
	a := 1000
	b := 2000
	// aとbが一致するとき
	if a == b {
		fmt.Println("a == b")
	} else if a > b { // aよりbが小さいとき
		fmt.Println("a > b")
	} else { // 上の条件に当てはまらないとき
		fmt.Println("a < b")
	}
}
