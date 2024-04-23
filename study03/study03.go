package study03

import "fmt"

func Greeting(tpyeInfos ...interface{}) {
	fmt.Println(len(tpyeInfos))
	for _, value := range tpyeInfos {
		switch v := value.(type) {
		case int:
			print("int")
		case float64:
			print("float64")
		case string:
			print("string")
		case bool:
			print("bool")
		default:
			fmt.Println(v, "interface{}")
		}
	}

}
