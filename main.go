package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"strconv"
)

const (
	PRINT_OPTION     = 1
	ADD_OPTION       = 2
	MARK_DONE_OPTION = 3
	REMOVE_OPTION    = 4
	EXIT_OPTION      = 5
	MAX_OPTION       = 5
)

var fileName string = "list.csv"
var optionsMessage string = `
Select an option:
  1.- Print list
  2.- Add task
  3.- Mark as DONE
  4.- Remove task
  5.- Exit
`

func main() {
	fmt.Println("Welcome to simple todo list!")

	file := OpenFile(fileName)

	defer file.Close()

	reader := csv.NewReader(file)

	tasks, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Println(optionsMessage)
		userOption, err := readUserInput()
		if err != nil {
			log.Println(err)
		}
		switch userOption {
		case 1:
			printTasks(tasks)
		case EXIT_OPTION:
			return
		}
	}
}

func readUserInput() (int, error) {
	var input string
	_, err := fmt.Scanln(&input)

	if err != nil {
		return 0, err
	}

	parsedInput, err := strconv.Atoi(input)
	if err != nil || !isValidInput(parsedInput) {
		log.Println("Invalid option")
		return 0, nil
	}

	return parsedInput, nil
}

func isValidInput(input int) bool {
	return input > 0 && input <= MAX_OPTION
}

func printTasks(tasks [][]string) {
	for _, task := range tasks {
		fmt.Println(task)
	}
}
