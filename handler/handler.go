package handler

import (
	"v0/button"
	"v0/command"
)

type InputHandler struct {
	commandFns map[button.ButtonType]command.ICommand
}

func New() *InputHandler {
	ih := &InputHandler{}
	ih.commandFns = make(map[button.ButtonType]command.ICommand)
	ih.registerCommand(button.BUTTON_A, &command.JumpCommand{})

	return ih
}

func (ih *InputHandler) registerCommand(buttonType button.ButtonType, c command.ICommand) {
	ih.commandFns[buttonType] = c
}

func (ih *InputHandler) HandleInput(buttonType button.ButtonType) {
	c := ih.commandFns[buttonType]
	c.Execute()
}
