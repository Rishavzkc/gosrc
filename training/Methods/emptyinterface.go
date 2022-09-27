package main

import "fmt"

func main() {

	student := make([]interface{}, 5)

	student[0] = 251
	student[1] = "String"
	student[2] = true

	for i, v := range student {
		fmt.Println(i, v)
	}
}
