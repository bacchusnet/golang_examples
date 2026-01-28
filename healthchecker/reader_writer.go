package main

import (
	"bufio"
	"os"
)

func File_reader(filename string) ([]string, error) {

	// We open the file
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	// We tell Golang to close the file once this function has finished running
	defer file.Close()

	// We create a buffer in memory for the contents of the file.
	scanner := bufio.NewScanner(file)

	// Create a slice to hold our items. Alternatively, if you know the size of the list, initialize with make([]string,0,100)
	urls := []string{}

	// We begin looping through the file content separated by new line characters and storing the contents in the buffer
	for scanner.Scan() {

		// Write to our urls slice by appending each line to the current slice
		urls = append(urls, scanner.Text())
	}

	return urls, err

}

func File_writer(filename string, data string) (string, error) {

	// We tell Golang to open the file for writing and appending and if the file does not exist, create it, and give read-write permission
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		return "Failed", err
	}

	// We tell Golang to close the file once this function has finished running
	defer file.Close()

	// We create the buffer
	bufferedWriter := bufio.NewWriter(file)

	// We write our data to the buffer
	_, err = bufferedWriter.WriteString(data + "\n")

	if err != nil {
		return "Failed", err
	}

	// We flush the buffered data to the file essentially writing to the file
	err = bufferedWriter.Flush()

	if err != nil {
		return "Failed", err
	}

	return "Success", err

}
