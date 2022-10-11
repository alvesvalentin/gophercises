package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	exercises()
}
func exercises() {
	commandFlag := flag.String("csv-file", "problems.csv", "CSV file commandFlag")

	flag.Parse()

	file, err := os.Open("quiz_game/" + *commandFlag)
	if err != nil {
		log.Fatalf("error opening %s", *commandFlag)
	}

	reader := csv.NewReader(file)

	all, err := reader.ReadAll()
	var correctAnswers int

	for _, val := range all {
		var input string
		fmt.Printf("Question: %s... \n", val[0])

		fmt.Scanln(&input)
		if input == val[1] {
			fmt.Println("Correct !")
			correctAnswers++
		} else {
			fmt.Println("Incorrect !")
		}
	}

	fmt.Printf("You have %d/%d correct answers ! \n", correctAnswers, len(all))

}