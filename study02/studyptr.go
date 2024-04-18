package study02

import "fmt"

func StudyPtr() {
	var i1 = 5
	var intp *int
	fmt.Printf("An integer: %d, it`s location in menory: %p\n", i1, &i1)

	intp = &i1

	fmt.Println(intp, *intp)

	*intp = 8
	fmt.Println(i1)
}

func StudyFor() {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
	}
	str2 := "日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}

}
