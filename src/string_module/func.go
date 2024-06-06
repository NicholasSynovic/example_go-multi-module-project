package string_module

import "fmt"

func HelloWorld() {
	fmt.Println(addPunctuation("Hello World"))
}

func addPunctuation(text string) string {
	return string(text + "!")
}
