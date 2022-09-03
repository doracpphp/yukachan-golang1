package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Webhookメッセージ用の構造体(jsonを利用)
type Webhooks struct {
	Content string `json:"content"`
}

// JSONから自動でできた天気予報の構造体
type Weathers []struct {
	PublishingOffice string    `json:"publishingOffice"`
	ReportDatetime   time.Time `json:"reportDatetime"`
	TimeSeries       []struct {
		TimeDefines []time.Time `json:"timeDefines"`
		Areas       []struct {
			Area struct {
				Name string `json:"name"`
				Code string `json:"code"`
			} `json:"area"`
			WeatherCodes []string `json:"weatherCodes"`
			Weathers     []string `json:"weathers"`
			Winds        []string `json:"winds"`
			Waves        []string `json:"waves"`
		} `json:"areas"`
	} `json:"timeSeries"`
	TempAverage struct {
		Areas []struct {
			Area struct {
				Name string `json:"name"`
				Code string `json:"code"`
			} `json:"area"`
			Min string `json:"min"`
			Max string `json:"max"`
		} `json:"areas"`
	} `json:"tempAverage,omitempty"`
	PrecipAverage struct {
		Areas []struct {
			Area struct {
				Name string `json:"name"`
				Code string `json:"code"`
			} `json:"area"`
			Min string `json:"min"`
			Max string `json:"max"`
		} `json:"areas"`
	} `json:"precipAverage,omitempty"`
}

func main() {
	//天気予報のURL
	weatherUrl := "https://www.jma.go.jp/bosai/forecast/
	data/forecast/エリアコード.json"
	// HTTPのGETメソッドと接続先のURLを指定する
	req, _ := http.NewRequest(http.MethodGet, weatherUrl, nil)
	// http.Clientを生成
	client := new(http.Client)
	// reqを元に気象庁のサイトにアクセス
	resp, err := client.Do(req)
	// エラー処理
	if err != nil {
		fmt.Println("HTTP Requestエラー", err)
		return
	}
	// TCPコネクションを閉じる(関数がreturnするまで実行されない)
	defer resp.Body.Close()
	// Bodyの内容を全て読み込む
	body, err := io.ReadAll(resp.Body)
	// エラー処理
	if err != nil {
		fmt.Println("io.ReadALLエラー ", err)
		return
	}
	// 天気予報の構造体を定義
	var getWeathers Weathers
	// JSONを構造体に変換
	json.Unmarshal(body, &getWeathers)

	//DiscordのWebHookのURLを入れてね
	var discordUrl = "https://discord.com/api/webhooks/xxxx/yyyy"
	// メッセージ用の構造体を作成
	webhook := &Webhooks{
		//【重要】天気予報をContentに格納するよ
		Content: getWeathers[0].TimeSeries[0].Areas[0].Weathers[0],
	}
	//jsonを[]byteに変換
	j, err := json.Marshal(webhook)
	if err != nil {
		fmt.Println("jsonエラー :", err)
		return
	}
	//httpリクエストの作成
	req2, err := http.NewRequest("POST", discordUrl, bytes.NewBuffer(j))
	// エラー処理
	if err != nil {
		fmt.Println("リクエストエラー : ", err)
		return
	}
	//Content-Typeにjsonを設定
	req2.Header.Set("Content-Type", "application/json")
	// httpクライアントの作成
	client2 := new(http.Client)
	// クライアントで設定されたポリシー (リダイレクト、クッキー、認証など) に従って、
	// HTTP リクエストを送信し、HTTP レスポンスを返す
	resp2, err := client2.Do(req2)
	if err != nil {
		fmt.Println("http clientエラー : ", err)
		return
	}
	//ステータスコードが204なら送信成功だよ！
	if resp2.StatusCode == 204 {
		fmt.Println("送信成功!!") //成功
	} else {
		fmt.Printf("送信失敗 : %#v\n", resp2) //失敗
	}
}
