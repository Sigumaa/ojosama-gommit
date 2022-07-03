package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jiro4989/ojosama"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("引数の数が間違っておりましてよ。")
		os.Exit(1)
	}
	args := os.Args[1]

	text, err := ojosama.Convert(args, nil)
	if err != nil {
		fmt.Println("メッセージの変換に失敗してしまいましたわ。")
		os.Exit(1)
	}

	out, err := exec.Command("git", "commit", "-m", text).Output()
	if err != nil {
		fmt.Println("commitに失敗してしまいましたわ。addはしましたの？")
		os.Exit(1)
	}
	fmt.Println(string(out))
}
