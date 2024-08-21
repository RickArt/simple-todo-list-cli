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

	taskMap := loadTaskMap(reader)

	for {
		fmt.Println(optionsMessage)
		userOption, err := readUserOption()
		if err != nil {
			log.Println(err)
		}
		switch userOption {
		case PRINT_OPTION:
			printTasks(taskMap)
		case REMOVE_OPTION:
			taskIdToDelete, err := readTaskIdToDelete()
			if err != nil {
				log.Fatal(err)
			}
			delete(taskMap, taskIdToDelete)
		case EXIT_OPTION:
			return
		}
	}
}

func readUserOption() (int, error) {
	var input string
	_, err := fmt.Scanln(&input)

	if err != nil {
		return 0, err
	}

	parsedInput, err := strconv.Atoi(input)
	if err != nil || !isValidOption(parsedInput) {
		log.Println("Invalid option")
		return 0, nil
	}

	return parsedInput, nil
}

func isValidOption(input int) bool {
	return input > 0 && input <= MAX_OPTION
}

func printTasks(taskMap map[int]Task) {
	for _, task := range taskMap {
		fmt.Printf("#%v - %v - %v\n", task.Id, task.Description, task.getDoneSymbol())
	}
}

func loadTaskMap(reader *csv.Reader) (taskMap map[int]Task) {
	taskMap = make(map[int]Task)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		done, _ := strconv.ParseBool(record[2])
		data := Task{
			Id:          id,
			Description: record[1],
			Done:        done,
		}
		taskMap[id] = data
	}
	return taskMap
}

func readTaskIdToDelete() (int, error) {
	fmt.Println("Select the id number to delete the task")
	var input string
	_, err := fmt.Scanln(&input)

	if err != nil {
		return 0, err
	}

	parsedInput, err := strconv.Atoi(input)
	if err != nil {
		log.Println("Invalid option")
		return 0, nil
	}

	return parsedInput, nil
}
