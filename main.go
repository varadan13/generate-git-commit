package main

import (
	"fmt"
	"generate-git-commit/internal/gemini"
)

func main() {
	fmt.Println("===================generate-git-command initialised===================")

	gemini.StreamContent()

}
