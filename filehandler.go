package main

import (
	"log"
	"os"
)

func OpenFile(path string) *os.File {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("Error reading the file: ", err)
	}
	return file
}
