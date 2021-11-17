package main

import (
	"github.com/ChillyWR/SSIA_demo1/iternal"
)

// TODO: cobra
func main() {
	iternal.ArgsHandler()
	quiz := iternal.CreateQuiz()
	quiz.StartQuiz()
}
