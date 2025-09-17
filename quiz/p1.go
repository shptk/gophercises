package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	var t_limit int
	flag.IntVar(&t_limit, "timeout",30,"pass a timeout in sec")
	flag.Parse()
	timeout := time.NewTimer(time.Duration(t_limit) * time.Second)
	
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
	rand.Shuffle(total, func(i, j int) {csv_data[i], csv_data[j] = csv_data[j], csv_data[i]})
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
		case <-timeout.C:
			break quizloop
		case user_input := <-answerCh:
			if user_input == row_value[1] {
				correct++
			}
		}

	}
	fmt.Printf("score: %d/%d\n", correct, total)
}
