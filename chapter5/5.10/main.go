package main

import "fmt"

//div関数
//関数は、func 関数名(引数,...) 戻り値　で表すよ！
func div(a int, b int) (int, error) {
	// a割るbの結果を返すよ
	if b == 0 {
		// fmt.Errorfはerror型を返すよ
		return 0, fmt.Errorf("0では割れないよ！")
	}
	// a/bとエラーがないからnilを返すよ
	return a / b, nil
}

func main() {
	// 関数divを実行して、結果を表示するよ
	// ans に割り算結果、errにnilが入るよ！
	ans, err := div(8, 2)
	// errはnilだからif文の中には入らないよ
	if err != nil {
		fmt.Println("0ではわれないよ！")
	}
	// 「8/2」だから4が表示されるよ！
	fmt.Println(ans)
	ans2, err := div(8, 0)
	// 0で割れないからerr2は「nil」以外が入っているよ！
	// だから、if文の中身が表示されるよ!
	if err != nil {
		fmt.Println("0ではわれないよ！")
	}
	//「8/0」だと割れないから0を表示するよ
	fmt.Println(ans2)
}
