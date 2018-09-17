package main

import (
	"fmt"
	"math"
)

var (aa = 3
ss = "kkk"
bb = true
)

func variable(){
	var a int
	var s string
	fmt.Printf("%d %q\n",a , s)
}

func variableInitialValue(){
	var a,b int = 3,4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction(){
	var a,b,c,s = 3,4,true, "def"
	fmt.Println(a,b,c,s)
}

func variableShorter(){
	a,b,c,s := 3,4,true, "def"
	b = 5
	fmt.Println(a,b,c,s)
}

func consts(){
	const(
		filename = "abc.txt"
		a,b = 3, 4
	)
	var c int
	c = int(math.Sqrt(a * a + b * b))
	fmt.Println(filename,c)
}

func enums(){
	const(
		cpp = iota
		_
		python
		golang
		javascript
	)
	fmt.Println(cpp,python,golang,javascript)

	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b,kb,mb,gb,tb,pb)
}

func main() {
	fmt.Println("Hello world")
	variable()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa,bb,ss)
	enums()
}
