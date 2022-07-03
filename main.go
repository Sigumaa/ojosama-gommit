package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jiro4989/ojosama"
)

func main() {
	args := os.Args[1]

	text, err := ojosama.Convert(args, nil)
	if err != nil {
		fmt.Println("メッセージの変換に失敗してしまいましたわ。")
		os.Exit(1)
	}

	out, err := exec.Command("git", "commit", "-m", text).Output()
	if err != nil {
		fmt.Println("commitに失敗してしまいましたわ。")
		os.Exit(1)
	}
	fmt.Println(string(out))
}
