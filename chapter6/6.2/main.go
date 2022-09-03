package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"
)
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
            Weathers     []string `json:"weathers"` //ここに天気予報が入るよ！
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
    weatherUrl := "https://www.jma.go.jp/bosai/forecast/data/
    forecast/エリアコード.json"
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
	//　構造体の中身を表示
    fmt.Println(getWeathers)
	// 構造体内の明日の天気予報を表示
    fmt.Println(getWeathers[0].TimeSeries[0].Areas[0].Weathers[1])
}