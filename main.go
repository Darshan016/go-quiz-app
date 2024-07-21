package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func problemPuller(fileName string) ([]problem, error) {
	fileObj, err := os.Open(fileName)
	handleError(err)
	csvR := csv.NewReader(fileObj)
	cLines, err := csvR.ReadAll()
	handleError(err)
	return parseProblem(cLines), nil
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

type problem struct {
	question string
	answer   string
}

func main() {
	fileName := flag.String("f", "quiz.csv", "path of csv file")
	timer := flag.Int("t", 30, "timer of the quiz")
	problems, err := problemPuller(*fileName)
	handleError(err)
	flag.Parse()

	correctAns := 0
	tObj := time.NewTimer(time.Duration(*timer) * time.Second)
	ansC := make(chan string)

problemLoop:
	for i, p := range problems {
		var answer string
		fmt.Printf("Problem %d: %s=", i+1, p.question)

		go func() {
			fmt.Scanf("%s", &answer)
			ansC <- answer
		}()
		select {
		case <-tObj.C:
			fmt.Println()
			break problemLoop
		case iAns := <-ansC:
			if iAns == p.answer {
				correctAns++
			}
			if i == len(problems)-1 {
				close(ansC)
			}
		}
	}
	fmt.Printf("result: %d out of %d\n", correctAns, len(problems))
	fmt.Println("Press enter to exit")
	<-ansC

}

func parseProblem(lines [][]string) []problem {
	r := make([]problem, len(lines))
	for i := 0; i < len(lines); i++ {
		r[i] = problem{question: lines[i][0], answer: lines[i][1]}
	}
	return r
}
