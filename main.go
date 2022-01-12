package main

import (
	"fmt"
	"go-staruml/mdjreader"
	"os"
)

func main() {
	fmt.Println("Hello World.")

	f, err := os.Open("pbt.mdj")
	if err != nil {
		panic("no file")
	}

	out, err := mdjreader.ReadMdj(f)
	if err != nil {
		panic(err)
	}

	println(out)
}
