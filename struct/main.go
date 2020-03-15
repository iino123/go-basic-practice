package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// type person struct {
// 	first   string
// 	last    string
// 	contact contactInfo
// }

// emdebbedする時はfield名を省略したら、struct名がfield名になる。
type person struct {
	first string
	last  string
	contactInfo
}

func main() {
	// 初期化の方法二つ
	// alex := person{"Alex", "AAA"}
	// alex := person{first: "Alex", last: "AAA"}

	// var alex person
	// // structのvalueしか出力しないため { } が出力される（空文字列が二つ）
	// // javascriptの連想配列と違って、field名(key名)は出ないので注意が必要
	// fmt.Println(alex)
	// // %+vでstructのfieldも出力する
	// fmt.Printf("%+v", alex)

	ken := person{
		first: "iino",
		last:  "kenjiro",
		contactInfo: contactInfo{
			email:   "iino@gmail.com",
			zipCode: 123,
		},
	}

	// fmt.Printf("%+v", ken)

	// 省略しない書き方 レシーバがポインタになっている
	// kenPointer := &ken // kenというstructのメモリアドレスを参照
	// kenPointer.updateName("ken!!!")

	// 省略記法を使った書き方(レシーバがvalueだけど、goが暗黙的にポインタとして認識してくれる)
	ken.updateName("ken!!!")
	ken.print()
}

// ここの*personは「(structである)personのメモリアドレス(ポインタ)を示す型」という意味。決してポインタに対応する値を参照している訳ではない。
// -> この関数はレシーバにpersonのアドレス(ポインタ)を受けとることができる。
func (pointerToPerson *person) updateName(newFirst string) {
	// pointerToPersonはポインタ(メモリのアドレス)を指す
	// *pointerToPersonとすることでポインタに対応する値を参照する
	(*pointerToPerson).first = newFirst
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

/*
重要: goは関数の引数に値(strct)を渡す時に、デフォルトではメモリの別アドレスに元々の値のコピーを作って渡す仕様

1. &valiable 変数のメモリアドレス(ポインタ)を指す。
2. *pointer ポインタに対するvalueを指す。
3. *type typeに対するポインタを指定する型を示している。
2と3が非常に記法は同じだが全く意味が違うので注意する。*を見た時は、その後ろが「値か型か」をチェックすること！

valueのアドレスを知りたい --> &をつける(ex. &value)
アドレスのvalueを知りたい --> *をつける(ex. *address)
アドレスを示す型を表現したい --> *をつける(ex. *person)
*/

/*
大前提として、goで関数の引数は、渡された値のコピーに対する参照である。

int, float, string, bool, struct --> value type(変数に代入した時にそのメモリのアドレスに対する値に実体が存在する)
--> 関数内でこれらの値を変更する時は、引数としてポインタ型を受け取るようにする

slices, maps, channels, pointers, functions --> reference type(変数に代入した時にそのメモリのアドレスに対する値には実体への参照を保持する。実体は別のメモリアドレスに存在する。)
--> 関数内でこれらの値を更新する時も、rubyなどと同じようにポインタを考慮する必要はない（実体への参照自体がコピーされても、参照先の実体は同じなので）
*/

/*
以下の例では、出力結果が Bye how are you となる（structと大きく異なることに注意）
package main

import (
	"fmt"
)

func main() {
	mySlice := []string{"Hi", "how", "are", "you"}
	update(mySlice)
	fmt.Println(mySlice)
}

func update(s []string) {
	s[0] = "Bye"
}
*/
