package main

import (
	"log"
	"os"
)

func OpenFile(path string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return CreateFile(path)
		}
		log.Fatal("Error reading the file: ", err)
	}
	return file
}

func CreateFile(path string) *os.File {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal("Error creating the file: ", err)
	}
	return file
}
