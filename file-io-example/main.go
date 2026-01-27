package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file_writer("urls.txt", "http://www.test.com")

	file_reader("urls.txt")
}

func file_reader(filename string) {

	// We open the file
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	// We tell Golang to close the file once this function has finished running
	defer file.Close()

	// We create a buffer in memory for the contents of the file.
	scanner := bufio.NewScanner(file)

	// We begin looping through the file content separated by new line characters and storing the contents in the buffer
	for scanner.Scan() {

		// We read the current string in the buffer and assign it to the line variable
		line := scanner.Text()

		fmt.Println(line)
	}
}

func file_writer(filename string, data string) {

	// Set file permission to Read-Write
	err := os.Chmod("urls.txt", 0644)

	if err != nil {
		log.Fatal(err)
	}

	// We tell Golang to open the file for writing and appending and if the file does not exist, create it, and give read-write permission
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(err)
	}

	// We tell Golang to close the file once this function has finished running
	defer file.Close()

	// We create the buffer
	bufferedWriter := bufio.NewWriter(file)

	// We write our data to the buffer
	_, err = bufferedWriter.WriteString("\n" + data)

	if err != nil {
		panic(err)
	}

	// We flush the buffered data to the file essentially writing to the file
	err = bufferedWriter.Flush()

	if err != nil {
		log.Fatal(err)
	}

	// We set permissions back to Read-Only
	err = os.Chmod("urls.txt", 0400)

}
