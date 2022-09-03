package main

import "fmt"

func main() {
	//テストの点数配列
	math := []int{95, 55, 20}
	//名前の配列
	name := []string{"行村先輩", "ゆかちゃん", "GO言語ちゃん"}
	fmt.Println(math)
	fmt.Println(name)
	// 1番目のデータは[0],2番目のデータは[1] でアクセスできるよ!!
	fmt.Println("1番目のデータ : ", name[0])
}
