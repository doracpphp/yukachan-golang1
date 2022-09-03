package main

import (
	"encoding/json"
	"fmt"
)

// プロフィール構造体
type Profile struct {
	Age    int     `json:"age"`    //年齢
	Name   string  `json:"name"`   //名前
	Height float32 `json:"height"` //高さ(cm)
}

func main() {
	//yukachanのプロフィール(JSON形式)
	//「`」←を使うと「"」を文字列として扱えるよ！
	yukachan := `{ "name" : "ゆかちゃん", "age" : 15, "height" : 150.4}`
	//JSONのデータを格納するプロフィールの構造体を定義
	var profile Profile
	//「json.Unmarshal」でJOSNを構造体に変換するよ!
	err := json.Unmarshal([]byte(yukachan), &profile)
	//エラー処理(エラーハンドリングを参照)
	if err != nil {
		fmt.Println("jsonエラー:", err)
	}
	//プロフィールの構造体を表示
	fmt.Printf("%#v\n", profile)
}
