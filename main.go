package main

import (
	"fmt"
	"strconv"
)

func main() {
	name, age, likesGo := "Hugo", 32, true

	fmt.Println("My name is " + name + " and I am " + strconv.Itoa(age) + " years old." + " Do I love Go? " + strconv.FormatBool(likesGo))
}