package main

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var items []string
var Value []string
var ItemsTotal []int

func readCSV(filee string) {
	var conv1 int = 0
	var conv2 int = 0
	var err1 error = nil
	var err2 error = nil
	//fmt.Println(filee)
	items = nil
	Value = nil
	ItemsTotal = nil
	// Open the file
	csvfile, err := os.Open(ExecPath + "/data/" + filee + ".csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(len(record))
		//fmt.Printf("Question: %s Answer %s\n", record[0], record[1])
		items = append(items, record[0])
		Value = append(Value, record[1])
		if record[2] == "" {
			conv1 = 0
		} else {
			conv1, err1 = strconv.Atoi(record[2])
		}
		if record[3] == "" {
			conv2 = 0
		} else {
			conv2, err2 = strconv.Atoi(record[3])
		}

		if err1 != nil {
			fmt.Println("Error converting to string.")
		} else if err2 != nil {
			fmt.Println("Error converting to string.")
		} else {
			ItemsTotal = append(ItemsTotal, conv1+conv2)
		}
	}
}

func readCSVItemSearch(filee string, ItemToFind string) (string, string, int, int, int, string) {
	var conv1 int = 0
	var conv2 int = 0
	var err1 error = nil
	var err2 error = nil
	//fmt.Println(filee)
	//items = nil
	//Value = nil
	//ItemsTotal = nil
	// Open the file
	csvfile, err := os.Open(ExecPath + "/data/" + filee + ".csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(len(record))
		//fmt.Printf("Question: %s Answer %s\n", record[0], record[1])
		//items = append(items, record[0])
		//Value = append(Value, record[1])
		if record[2] == "" {
			conv1 = 0
		} else {
			conv1, err1 = strconv.Atoi(record[2])
		}
		if record[3] == "" {
			conv2 = 0
		} else {
			conv2, err2 = strconv.Atoi(record[3])
		}

		if err1 != nil {
			fmt.Println("Error converting to string.")
		} else if err2 != nil {
			fmt.Println("Error converting to string.")
		} else {
			if record[0] == ItemToFind {
				return record[0], record[1], conv1, conv2, conv1 + conv2, record[4]
			}
			//ItemsTotal = append(ItemsTotal, conv1+conv2)
		}
	}
	return "Nothing to find", "Error", 0, 0, 0, "Error"
}
