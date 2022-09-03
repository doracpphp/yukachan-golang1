package main

import "fmt"

// プロフィール構造体
type Profile struct {
	Age    int     //年齢
	Name   string  //名前
	height float32 //身長(cm)
}

func main() {
	//ゆかちゃんのプロフィールを定義
	yuka := Profile{
		Age:    15,      //年齢
		Name:   "ゆかちゃん", //名前
		height: 150.4,   //身長(cm)
	}
	// %#v : 構造体のフィールド名も表示されるよ！！
	fmt.Printf("%#v\n", yuka)
	// 名前だけ表示したい時は「変数名.Age」でアクセスするんだ
	fmt.Printf("%#v\n", yuka.Age)
}
