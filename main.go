package main

import (
	"fmt"
	"github.com/jiro4989/ojosama"
	"os"
)

func main() {
	args := os.Args[1]

	text, err := ojosama.Convert(args, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
}
