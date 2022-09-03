package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Webhookメッセージ用の構造体(jsonを利用)
type Webhooks struct {
	Content string `json:"content"`
}

func main() {
	//DiscordのWebHookのURLを入れてね
	var discordUrl = "https://discord.com/api/webhooks/xxxx/yyyy"
	// メッセージ用の構造体を作成
	webhook := &Webhooks{
		// 【重要】Contentにメッセージを送りたい文書を入れてね
		Content: "Go言語最高！！",
	}
	//jsonを[]byteに変換
	j, err := json.Marshal(webhook)
	if err != nil {
		fmt.Println("jsonエラー :", err)
		return
	}
	//httpリクエストの作成
	req, err := http.NewRequest("POST", discordUrl, bytes.NewBuffer(j))
	//エラー処理
	if err != nil {
		fmt.Println("リクエストエラー : ", err)
		return
	}
	//Content-Typeにjsonを設定
	req.Header.Set("Content-Type", "application/json")
	// httpクライアントの作成
	client := new(http.Client)
	// クライアントで設定されたポリシー (リダイレクト、クッキー、認証など) に、
	// 従って、HTTP リクエストを送信し、HTTP レスポンスを返す
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http clientエラー : ", err)
		return
	}
	//ステータスコードが204なら送信成功だよ！
	if resp.StatusCode == 204 {
		fmt.Println("送信成功!!") //成功
	} else {
		fmt.Printf("送信失敗 : %#v\n", resp) //失敗
	}
}
