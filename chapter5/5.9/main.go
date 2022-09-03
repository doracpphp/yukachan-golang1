package main

import "fmt"

//div関数
//関数は、func 関数名(引数,...) 戻り値　で表すよ！
func div(a int, b int) int {
	//a割るbの結果を返すよ
	return a / b
}

func main() {
	//関数divを実行して、結果を表示するよ
	fmt.Println(div(8, 2))
}
