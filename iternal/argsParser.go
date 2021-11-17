package iternal

import (
	"fmt"
	"os"
	"strconv"
)

const helpMessage = "Quiz program will ask questions one by one waiting for your reply\n" +
	"If your answers is correct you will receive points that are displayed at the end of a quiz\n" +
	"Be alert the timer will start counting down after you confirm the start of a quiz"
const flagsMassage = "Available flags:\n" +
	"-help: Print instructions\n" +
	"-shuffle: Shuffle quiz questions\n" +
	"--filename: Change default filename (specify new one as new command line argument)\n" +
	"--timer: Change default time for timer to count down (specify new one as new command line argument)"

var (
	flags = map[string]interface{}{
		"-help":    printHelpMessage,
		"-shuffle": shuffleQuizQAs,
	}
	flagsWithArg = map[string]interface{}{
		"--filename": changeDefaultFilename,
		"--time":     changeDefaultTime,
	}
)

func ArgsHandler() {
	args := os.Args

	// For debug
	// fmt.Println(args)
	// args = []string{"./main", "--time", "10", "--filename", "extra.json", "-shuffle"}

	for i := 1; i < len(args); i++ {
		if value, ok := flags[args[i]]; ok {
			value.(func())()
		} else if valueWA, okWA := flagsWithArg[args[i]]; okWA && i+1 < len(args) {
			valueWA.(func(string))(args[i+1])
			i++
		} else {
			fmt.Println("Invalid flag format\n\n" + flagsMassage)
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

func printHelpMessage() {
	fmt.Println(helpMessage + "\n" + flagsMassage)
	os.Exit(0)
}
