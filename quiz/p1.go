package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	csv_reader := csv.NewReader(f)
	csv_data, err := csv_reader.ReadAll()
	correct := 0
	total := len(csv_data)
	if err == io.EOF {
		return
	}
	if err != nil {
		log.Fatal(err)

	}
	timer := time.NewTimer(15 * time.Second)
	answerCh := make(chan string)
quizloop:
	for _, row_value := range csv_data {
		// fmt.Printf("%+v\n", l)

		fmt.Print(row_value[0], ": ")
		go func() {

			var user_input string
			fmt.Scanln(&user_input)
			answerCh <- user_input
		}()

		select {
		case <-timer.C:
			break quizloop
		case user_input := <-answerCh:
			if user_input == row_value[1] {
				correct++
			}
		}

	}
	fmt.Printf("score: %d/%d\n", correct, total)
}
