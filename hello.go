package main

import "hellogo/study03"

func main() {
	// study01.StudyStr()
	// study03.Greeting("hao", "Li", "Zhu", "Wu")
	option1 := new(Options)

	option1.f = 3.14
	option1.s = "hello"
	option1.b = false
	option1.i = 1

	option2 := new(Options)

	option2.i = 2
	option2.f = 23.14
	option2.s = "2hello"
	option2.b = true

	study03.Greeting(option1, option2)

}

type Options struct {
	i int
	f float64
	s string
	b bool
}
