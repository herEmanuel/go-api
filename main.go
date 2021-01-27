package main

import (
	"fmt"
	"strconv"
)

func appendToString(str string, number int) string {
	newString := str + strconv.Itoa(number)

	return newString
}

func main() {

	fmt.Println(appendToString("testando", 234))
}
