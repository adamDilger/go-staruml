package main

import (
	"fmt"
	"go-staruml/mdjreader"
	"strings"
)

const input = `
{
	"_type": "Project",
	"_id": "AAAAAAF9KutqdluOyNg=",
	"name": "pbt"
}
`

func main() {
	fmt.Println("Hello World.")

	sr := strings.NewReader(input)

	out, err := mdjreader.ReadMdj(sr)
	if err != nil {
		panic(err)
	}

	println(out)
}
