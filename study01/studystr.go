package study01

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func StudyStr() {
	str := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str3 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined by ;: %s\n", str3)
}

func StudyStrconv() {
	val, _ := strconv.Atoi("666")
	fmt.Println(val)
}

func StudyTime() {
	t := time.Now()
	fmt.Printf("%02d/%02d/%4d\n", t.Day(), t.Month(), t.Year())

	fmt.Println(t.Format("02 Jan2006 15:04"))
	fmt.Println(t.Format(time.RFC822))

}
