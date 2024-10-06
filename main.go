package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type JsonRecord struct {
	Preview bool      `json:"preview"`
	Result  ResultObj `json:"result"`
}

type ResultObj struct {
	ThreadName       string `json:"thread_name"`
	WhereClause      string `json:"where_clause"`
	TimeElapsed      string `json:"time_elapsed"`
	RetrievedRecords int    `json:"retrieved_records"`
}

func main() {
	filepath := os.ExpandEnv("$HOME/Dev/Go/Projects/dynasearchtools/resources/test.json")
	jsonFile, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Successfully opened file")
	defer jsonFile.Close()

	// Because input is stupid and only a valid json per line.
	scanner := bufio.NewScanner(jsonFile)

	var jsonLine []string
	for scanner.Scan() {
		// read line and store to jsonLine
		jsonLine = append(jsonLine, scanner.Text())
	}

	for _, line := range jsonLine {
		fmt.Println("Handling line: ", line)
	}

}
