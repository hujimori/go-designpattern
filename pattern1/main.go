package main

import (
	"v0/button"
	"v0/handler"
)

func main() {
	ih := handler.New()
	ih.HandleInput(button.BUTTON_A)
}
