package main

import (
	"fmt"
	"time" //二つの処理を同時に走らせる
	// 様々なpackage
	"net/http"
)

//重い処理
func task1(result chan string) {
	time.Sleep(time.Second * 2)
	// 	fmt.Println("task1 finished")
	result <- "task1 result"
}

//軽い処理
func task2() {
	fmt.Println("task2 finished")
}

/*string型を引数にとり出力するメソッド　msg string 出力する値の宣言が必要であることに注意
func hi(name string) (msg string) {
	// 	fmt.Println("hi!!" + name)
	msg = "hi!!" + name
	return
}
*/
//二つの値を同時に返り値ができる（C言語では難しい）
// func swap(a, b int) (int, int) {
// 	return b, a
// }

//構造体のフィールド
// type user struct {
// 	name  string
// 	score int
// }

//インターフェイス:メソッドの一覧を定義したデータ型。type ○○　interface{}に入れる
//空の場合、「すべてのデータ型がこのインターフェイスを満たしている」
//あらゆる方を受け取ることができる関数を作れる
// type greeter interface {
// 	greet()
// }
// type japanese struct{}
// type american struct{}

// //あらゆる方を受け取ることができる
// func show(t interface{}) {
// 	//型アサーション:値と真偽値を返す
// 	// 	_, ok := t.(japanese)
// 	// 	if ok {
// 	// 		fmt.Println("i am japanese")
// 	// 	} else {
// 	// 		fmt.Println("i am not japanese")
// 	// 	}

// 	//型Switch: 渡ってきたデータの名前が返ってくる japanese american
// 	// 	switch t.(type) {
// 	// 	case japanese:
// 	// 		fmt.Println("i am japanese")
// 	// 	default:
// 	// 		fmt.Println("i am not japanese")
// 	// 	}
// }

// func (j japanese) greet() {
// 	fmt.Println("こんにちは")
// }

// func (a american) greet() {
// 	fmt.Println("hello")
// }

//webサーバーの作成
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi %s!", r.URL.Path[1:])
}

func main() {
	// 	var msg string
	// 	msg = "hello world"
	// 	msg := "hello world"
	// 	fmt.Println(msg)
	// int の書き方
	// 	var a, b int
	// 	a, b = 10, 15
	// 	a, b := 10, 15

	/*変数の宣言方法
	var (
		c int
		d string
	)
	c = 20
	d = "hoge"
	*/

	/*
			// 型の種類
			string "hello"
			int     53
			float64 10.2
			bool    true false
			nil
		    //初期値を把握しておこう
			var s string // ""
		    var a int // 0
		    var f bool // false
	*/

	/*変数の宣言、代入、デフォルト値
	a := 10
	b := 10.2
	c := "hoge"
	var d bool
	fmt.Printf("a: %d, b: %f, c: %s, d: %t\n", a, b, c, d)
	*/

	//定数の宣言、値を変更できないようにするために使用
	// 	const name = "fkoji"　//cannot assign to name
	// 	name = "takuo" //
	// 	fmt.Printf(name)

	/*
		const (
			//iota 宣言した値に+1してさらに定義していく
			//ID?
			sun = iota
			mon
			tue
			wed
			thr
			fri
		)
		fmt.Println(sun, mon, tue)
	*/
	// 	var x int
	// 	x = 10 * 3
	// 	fmt.Println(x)

	/* true false の判定方法
	a := true
	b := false
	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!a)
	*/

	/*
		//ポインタ
		//アドレス指定する変数
		// 演算はできない
		//メモリの節約に使用する
		a := 5
		var pa *int //int型を格納する領域のアドレスを準備する
		pa = &a     //&a =　aのアドレスを表現
		// pa の領域にあるデータの値 = *pa
		fmt.Println(pa)  //0xc00001a118
		fmt.Println(*pa) //5
	*/

	//hi()メソッドの呼び出し
	// 	hi("taguchi")

	//hi()で出力値を指定する場合
	//fmt.Println(hi("taguchi"))
	//func swap()の呼び出しに使用
	//fmt.Println(swap(5, 2))

	/*
		//メソッドを代入
		f := func(a, b int) (int, int) {
			return b, a
		}
		//メソッドを代入
		fmt.Println(f(3, 2))

		//即時関数　関数を宣言してその場で使用する（遠くで引用せず直近の処理で実行する場合に使う）
		func(msg string) {
			fmt.Println(msg)
		}("taguchi")
	*/

	/*配列の使い方
	var a [5]int   //a[0]-a[4]
	fmt.Println(a) //[0 0 0 0 0]
	a[2] = 3
	a[4] = 10
	fmt.Println(a)
	fmt.Println(a[2])
	*/
	/*配列の宣言＆挿入を１行で！
	b := [...]int{1, 3, 5}
	fmt.Println(b)
	fmt.Println(len(b)) //配列の個数
	*/

	/*C言語とは異なり、配列の宣言を行ってもa[0]へのポインタ
	（メモリのアドレス）にはならない。あくまでも配列*/
	//動的に変化させることもできず、関数の配列は丸ごとデータを渡すため、
	//メモリ的にも非効率

	/*スライスの使い方1*/
	/*a := [5]int{1, 10, 5, 18, 24}
	s := a[2:4]
	s2 := a[2:]
	s3 := a[:4]
	s[1] = 12 //10⇒12　もとの配列の要素も変更する
	fmt.Println(a)
	fmt.Println(s)       //[5,18]
	fmt.Println(s2)      //[5~]
	fmt.Println(s3)      //[~18]
	fmt.Println(cap(s3)) //[~18]　切り出しうる最大数の配列要素を出力
	*/
	/*スライスの使い方2
	// 	s := make([]int, 3) //[0,0,0]
	s := []int{1, 3, 5}
	//append 配列に動的に追加
	s = append(s, 8, 2, 1)
	//copy
	t := make([]int, len(s))
	n := copy(t, s)

	fmt.Println(s) //sに8,2,1を追加した配列
	fmt.Println(t) //s のコピー
	fmt.Println(n) //配列数　6
	*/
	/*マップmap : 添え字にキーを使ってハッシュ値のようにキーと値を
	  ペアで管理するデータ型*/
	// 	m := make(map[string]int)
	// 	m["taguchi"] = 200
	// 	m["fkoji"] = 300
	// 	fmt.Println(m) //map[fkoji:300 taguchi:200]
	// 	m := map[string]int{"taguchi": 100, "fkoji": 500}
	// 	fmt.Println(m)      //map[fkoji:300 taguchi:200]
	// 	fmt.Println(len(m)) //2
	// 	delete(m, "taguchi")

	// 	fmt.Println(len(m)) //1

	/*キーの存在を調べる
	v, ok := m["fkoji"]
	fmt.Println(v)  // 存在の有無、値の出力
	fmt.Println(ok) // 存在の有無、値の出力
	*/

	//if　条件分岐
	// 	score := 83
	// 	Score := 40
	// 	if Score > 80 {
	// 		fmt.Println("Great!")
	// 	} else if Score > 60 {
	// 		fmt.Println("OK!")
	// 	} else {
	// 		fmt.Println("soso...")
	// 	}
	// 	fmt.Println(Score)

	/*switch*/
	// 	signal := "green"
	// 	switch signal {
	// 	case "red":
	// 		fmt.Println("stop")

	// 	case "blue", "green":
	// 		fmt.Println("go")

	// 	case "yellow":
	// 		fmt.Println("hey!")
	// 	default:
	// 		fmt.Println("wrong signal")
	// 	}

	/*for*/
	// 	for i := 0; i < 10; i++ {
	// 		// 		if i == 3 {
	// 		// 			break
	// 		// 		}
	// 		if i == 3 {
	// 			continue
	// 		}

	// 		fmt.Println(i)
	// 	}
	// 	i := 0
	// 	for i < 10 {
	// 		fmt.Println(i)
	// 		i++
	// 	}
	// 	i := 0
	// 	for {
	// 		fmt.Println(i)
	// 		i++
	// 		if i == 10 {
	// 			break
	// 		}
	// 	}

	/*配列、スライス、マップの違い*/
	// 配列・・値の配列、定数で後から変更不可、引用時すべての配列データを取得するので目盛りの無駄
	//スライス・・配列を動的に処理できる
	//マップ・・配列にキーと値で管理できる（ハッシュ）

	/*range：指定の要素分だけ処理したい場合*/
	//配列
	// 	s := []int{2, 3, 8}
	// 	for i, v := range s {
	// 		fmt.Println(i, v)
	// 	}
	//ブランク修飾子： _ 何が入ってきてもすべて廃棄する
	// 	for _, v := range s {
	// 		fmt.Println(v)
	// 	}

	// 	m := map[string]int{"taaguchi": 100, "fkoji": 200}
	// 	for k, v := range m {
	// 		fmt.Println(k, v)
	// 		// fmt.Println(k,v)
	// 	}

	//構造体：複数の値を意味のあるまとまりとして新しい型を定義することができる
	//userの構造体を確保し、初期化
	// 	u := new(user)
	// 	u.name = "taguchi"
	//     u.score = 20
	//     fmt.Println(u)//ポインタで返す　&{taguchi 30}
	// 	u := user{name: "taguchi", score: 50} //{taguchi 50}

	// 	fmt.Println(u)

	//メソッド（データ型に紐づいた関数）
	// 	u := user{name: "tagichi", score: 50}
	// 	u.hit()
	// 	u.show() //name:tagichi, score:50

	//インターフェイス
	// 	greeters := []greeter{japanese{}, american{}}
	// 	for _, greeter := range greeters {
	// 		greeter.greet()
	// 		show(greeter)
	// 	}

	//GOの並行処理を実装してみよう
	//そのまま上から処理
	// 	task1()
	// 	task2()

	// 	result := make(chan string)
	// 	//goで並列処理 go だけだとgoroutineが終わる前にmain関数が終わっているから
	// 	go task1(result)
	// 	go task2()

	// 	fmt.Println(<-result) //resultの結果がなければ結果が来るまで待機する
	// 	time.Sleep(time.Second * 3)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

//メソッド（データ型に紐づいた関数）
//メソッドと構造体を紐づける //u レシーバー
// func (u user) show() {
// 	fmt.Printf("name:%s, score:%d\n", u.name, u.score)
// }

// //*user 参照渡しにすると値を更新できる
// func (u *user) hit() {
// 	u.score++
// }

//チャネル：データの受け渡しをするパイプ
// 外部のtask1()から処理終了後、mainの関数に値を渡せる
