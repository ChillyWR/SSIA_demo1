package iternal

import (
	"fmt"
	"os"
	"strconv"
)

var (
	flags           = map[string]interface{}{
		"-shuffle": shuffleQuizQAs,
	}
	flagsWithArg = map[string]interface{}{
		"--filename": changeDefaultFilename,
		"--timer":    changeDefaultTime,
	}
)

func ArgsHandler() {
	args := os.Args
	for i := 1; i < len(args); i++ {
		if value, ok := flags[args[i]]; ok {
			value.(func())()
		}
		if valueWA, okWA := flagsWithArg[args[i]]; okWA && i+1 < len(args) {
			valueWA.(func(string))(args[i+1])
			i++
		} else {
			fmt.Println("Invalid flag format")
			os.Exit(1)
		}

	}
}

func changeDefaultTime(newTime string) {
	newTimeInt, err := strconv.Atoi(newTime)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Time changed to %d\n", newTimeInt)
	TimeForQuiz = newTimeInt
}
func changeDefaultFilename(newFileName string) {
	fmt.Printf("Filename changed to %s\n", newFileName)
	InputFileName = newFileName
}
func shuffleQuizQAs() {
	Shuffle = true
}