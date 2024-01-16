package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	a := os.Args[1]
	data, err := ioutil.ReadFile(a + ".txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}
