package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	PRINT_OPTION            = 1
	ADD_OPTION              = 2
	TOGGLE_DONE_OPTION      = 3
	REMOVE_OPTION           = 4
	SAVE_AND_EXIT_OPTION    = 5
	DISCARD_AND_EXIT_OPTION = 6
	MAX_OPTION              = 6
)

var fileName string = "list.csv"

var optionsMessage string = `
Select an option:
  1.- Print list
  2.- Add task
  3.- Toggle DONE
  4.- Remove task
  5.- Save and Exit
  6.- Discard changes and Exit
`

func main() {
	fmt.Println("Welcome to simple todo list!")

	file := OpenFile(fileName)

	defer file.Close()

	reader := csv.NewReader(file)

	taskMap, highestId := loadTaskMap(reader)

	for {
		fmt.Println(optionsMessage)
		userOption, err := readUserOption()
		if err != nil {
			log.Println(err)
		}
		switch userOption {
		case PRINT_OPTION:
			printTasks(taskMap)
		case ADD_OPTION:
			addTask(taskMap, &highestId)
		case TOGGLE_DONE_OPTION:
			markTaskAsDone(taskMap)
		case REMOVE_OPTION:
			taskIdToDelete, err := readTaskIdToDelete()
			if err != nil {
				log.Fatal(err)
			}
			delete(taskMap, taskIdToDelete)
		case SAVE_AND_EXIT_OPTION:
			saveList(taskMap, file)
			fallthrough
		case DISCARD_AND_EXIT_OPTION:
			os.Exit(0)
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

func loadTaskMap(reader *csv.Reader) (taskMap map[int]Task, highestId int) {
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
		if id > highestId {
			highestId = id
		}
	}
	return taskMap, highestId
}

func readTaskIdToDelete() (int, error) {
	fmt.Println("Select the id number to delete the task: ")
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

func markTaskAsDone(taskMap map[int]Task) {
	fmt.Println("Select the id number to toggle Done: ")
	var input string
	_, err := fmt.Scanln(&input)

	if err != nil {
		log.Println(err)
	}

	parsedInput, err := strconv.Atoi(input)
	if err != nil {
		log.Println("Invalid option")
	}

	task := taskMap[parsedInput]
	task.Done = !task.Done
	taskMap[parsedInput] = task
}

func saveList(taskMap map[int]Task, file *os.File) {
	writer := csv.NewWriter(file)
	file.Truncate(0)
	file.Seek(0, 0)

	for _, task := range taskMap {
		writer.Write(taskToRecord(task))
	}
	writer.Flush()
}

func taskToRecord(task Task) (record []string) {
	return []string{
		strconv.Itoa(task.Id),
		task.Description,
		strconv.FormatBool(task.Done),
	}
}

func addTask(taskMap map[int]Task, highestId *int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Task Description: ")
	taskDescription, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}
	taskDescription = strings.TrimSpace(taskDescription)
	*highestId = *highestId + 1
	task := Task{
		Id:          *highestId,
		Description: taskDescription,
		Done:        false,
	}
	taskMap[*highestId] = task
}
