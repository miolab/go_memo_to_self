package main

import "fmt"

/*
`func init`で一番初めに実行する
*/
func init() {
	fmt.Println("Hello world, init.")
}

/*
`func main`で関数呼び出し実行
*/
func helloInsert() {
	fmt.Println("Hello world, insert.")
}

func main() {
	// 上で定義した`helloInsert`関数を呼び出し
	helloInsert()
	fmt.Println("Hello world.", " Hello im.")
}
