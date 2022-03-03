package command

import "fmt"

type JumpCommand struct {
}

func (jc *JumpCommand) Execute() {
	fmt.Println("ジャンプ")
}
