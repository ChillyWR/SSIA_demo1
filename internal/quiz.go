package internal

import (
	"bufio"
	"fmt"
	"github.com/ChillyWR/SSIA_demo1/pkg"
	"math/rand"
	"os"
	"sync"
	"time"
)

type QA struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type Quiz struct {
	qa                   []QA
	time                 int
	correctAnswersAmount int
}

const (
	maxQuestions = 100
)

var (
	InputFileName = "problems.json"
	TimeForQuiz   = 30 // in seconds
	Shuffle       = false
	wg            = sync.WaitGroup{}
)

// CreateQuiz create and return Quiz object
// Configured by variables InputFileName, TimeForQuiz, Shuffle
func CreateQuiz() *Quiz {
	q := new(Quiz)
	q.time = TimeForQuiz
	q.correctAnswersAmount = 0
	q.qa = make([]QA, 0, maxQuestions)
	pkg.ReadAndSaveJSON(InputFileName, &q.qa)
	if Shuffle {
		q.shuffleQuiz()
	}
	return q
}

// StartQuiz set up timer and print questions via askQuestions
func (q *Quiz) StartQuiz() {
	fmt.Printf("Press enter whenever you are ready\nYou will have %d seconds to finish the quiz\n", q.time)
	_, err := fmt.Scanln()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	go q.startTimer()
	q.askQuestions()
}

// Simple shuffle
func (q *Quiz) shuffleQuiz() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(q.qa),
		func(i, j int) {
			q.qa[i], q.qa[j] = q.qa[j], q.qa[i]
		})
}

// Print question and read input
// Compare with answer in parseInput
func (q *Quiz) askQuestions() {
	fmt.Println("Quiz Started")
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < len(q.qa); i++ {
		fmt.Printf("Question %d:\n%s\n", i + 1, q.qa[i].Question)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		wg.Add(1)
		q.checkAnswer(input, q.qa[i].Answer)
	}
	q.printResult()
}

func (q *Quiz) checkAnswer(input, answer string) {
	if pkg.CompareStrings(input, answer) {
		q.correctAnswersAmount++
	}
	wg.Done()
}

// Start timer
// Preferred to run in different go routine
// Shut down program when timer is up
func (q *Quiz) startTimer() {
	timer := time.NewTimer(time.Duration(q.time) * time.Second)
	<-timer.C
	fmt.Println("\nTime is up")
	q.printResult()
	wg.Wait()
	os.Exit(0)
}

func (q *Quiz) printResult() {
	fmt.Printf("You got %d correct answers out of %d\n", q.correctAnswersAmount, len(q.qa))
}
