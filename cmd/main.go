package main

import (
	"os"
	"github.com/ChillyWR/SSIA_demo1/internal"
)

// TODO: cobra
func main() {
	internal.ArgsHandler(os.Args)
	quiz := internal.CreateQuiz()
	quiz.StartQuiz()
}
