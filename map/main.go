package main

import "fmt"

func main() {
	// mapを初期化する方法は下記の2つ
	// var colors map[string]string
	// colors := make(map[string]string)

	// map[keyの型]valueの型
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf754",
		"white": "#ffffff",
	}

	// mapの要素を追加/削除する方法
	// colors["white"] = "#ffffff"
	// delete(colors, "white")

	// fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("color is", color, ",", "hex is", hex)
	}
}

/*
structとmapの違い

map
- reference type
- key名でiterateできる(keyがindexをサポートしてる
- 全てのkey, 全てのvalueが同じtypeの必要がある

strct
- value type
- オブジェクト指向っぽい使い方する感じか(?)
*/
