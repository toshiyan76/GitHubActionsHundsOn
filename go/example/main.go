package main

import "fmt"

var version string // ビルド時にldflagsフラグ経由でバージョンを渡す

func main() {
	fmt.Printf("Example %s\n", version)
}
