package main

import (
	"github.com/jiro4989/ojosama"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[1]

	text, err := ojosama.Convert(args, nil)
	if err != nil {
		panic(err)
	}

	err = exec.Command("git", "commit", "-m", text).Run()
	if err != nil {
		panic(err)
	}
}
