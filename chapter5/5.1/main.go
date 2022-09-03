package main

import "fmt"

func main() {
	fmt.Printf("1+2=%d\n", 1+2)       //加算
	fmt.Printf("1-2=%d\n", 1-2)       //減算
	fmt.Printf("1x2=%d\n", 1*2)       //乗算
	fmt.Printf("1/2=%d\n", 1/2)       //除算
	fmt.Printf("1/2=%g\n", 1.0/2.0)   //小数点の計算
	fmt.Printf("1 mod 2 は、%d\n", 1%2) //剰余算
}
