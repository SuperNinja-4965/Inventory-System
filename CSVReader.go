package main

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

var items []string

func readCSV(filee string) {
	fmt.Println(filee)
	items = nil
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
		for i := 0; i <= len(record)-1; i++ {
			items = append(items, record[i])
		}
	}
}
