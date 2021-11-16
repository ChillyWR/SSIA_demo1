package main

import(
	"github.com/ChillyWR/SSIA_demo1/iternal"
)

func main() {
	iternal.ArgsHandler()
	quiz := iternal.CreateQuiz()
	quiz.StartQuiz()
	return
}

//type QA struct {
//	Question string `json:"question"`
//	Answer   string `json:"answer"`
//}
//
//type Quiz struct {
//	qa                   []QA
//	correctAnswersAmount int
//}
//
//func (q *Quiz) startQuiz() {
//	fmt.Printf("Press enter whenever you are ready\nYou will have %d seconds to finish the quiz\n", defaultTime)
//	_, err := fmt.Scanln()
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//	go q.startTimer()
//	q.askQuestions()
//}
//
//func (q *Quiz) askQuestions() {
//	fmt.Println("Quiz Started")
//	reader := bufio.NewReader(os.Stdin)
//	for i := 0; i < len(q.qa); i++ {
//		fmt.Println(q.qa[i].Question)
//		input, err := reader.ReadString('\n')
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		q.parseInputAnswer(input, q.qa[i].Answer)
//	}
//	q.printResult()
//}
//
//// deletes spaces anyway?
//func (q *Quiz) parseInputAnswer(input, answer string) {
//	input = strings.Join(strings.Fields(input), " ")
//	if strings.ToLower(input) == strings.ToLower(answer) {
//		q.correctAnswersAmount++
//	}
//}
//
//func (q Quiz) startTimer() {
//	timer := time.NewTimer(time.Duration(defaultTime) * time.Second)
//	<-timer.C
//	fmt.Println("\nTime ended")
//	q.printResult()
//	os.Exit(0)
//}
//
//func (q Quiz) printResult() {
//	fmt.Printf("You got %d correct answers\n", q.correctAnswersAmount)
//}
//
//func createQuiz(filename string, shuffle bool) *Quiz {
//	quiz := new(Quiz)
//	quiz.correctAnswersAmount = 0
//	quiz.qa = make([]QA, 0, maxQuestions)
//	byteValue := readInputJSON(filename)
//	err := json.Unmarshal(byteValue, &quiz.qa)
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//	if shuffle {
//		rand.Seed(time.Now().UnixNano())
//		rand.Shuffle(len(quiz.qa),
//			func(i, j int) {
//				quiz.qa[i], quiz.qa[j] = quiz.qa[j], quiz.qa[i]
//			})
//	}
//	return quiz
//}
//
//func readInputJSON(filename string) []byte {
//	jsonFile, err := os.Open(filename)
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//	byteValue, _ := ioutil.ReadAll(jsonFile)
//	err = jsonFile.Close()
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//	return byteValue
//}
//
//// TODO: -help
//var (
//	shuffle         = false
//	defaultFileName = "problems.json"
//	defaultTime     = 30
//	flags           = map[string]interface{}{
//		"-shuffle": shuffleQuizQAs,
//	}
//	flagsWithArg = map[string]interface{}{
//		"--filename": changeDefaultFilename,
//		"--timer":    changeDefaultTime,
//	}
//)
//
//const (
//	maxQuestions = 100
//)
//
//func argsHandler() {
//	args := os.Args
//	for i := 1; i < len(args); i++ {
//		if value, ok := flags[args[i]]; ok {
//			value.(func())()
//		}
//		if valueWA, okWA := flagsWithArg[args[i]]; okWA && i+1 < len(args) {
//			valueWA.(func(string))(args[i+1])
//			i++
//		} else {
//			fmt.Println("Invalid flag format")
//			os.Exit(1)
//		}
//
//	}
//}
//
//func changeDefaultTime(newTime string) {
//	newTimeInt, err := strconv.Atoi(newTime)
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//	fmt.Printf("Time changed to %d\n", newTimeInt)
//	defaultTime = newTimeInt
//}
//func changeDefaultFilename(newFileName string) {
//	fmt.Printf("Filename changed to %s\n", newFileName)
//	defaultFileName = newFileName
//}
//func shuffleQuizQAs() {
//	shuffle = true
//}
//
//// TODO: cobra + git
//func main() {
//	argsHandler()
//	quiz := createQuiz(defaultFileName, shuffle)
//	quiz.startQuiz()
//	return
//}
