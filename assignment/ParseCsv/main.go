package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type tableRow struct {
	row []string
}

func parse_csv(data [][]string) []tableRow {
	var table []tableRow
	for _, line := range data {
		row := tableRow{row: line}
		table = append(table, row)
	}
	return table
}

func main() {
	var filePath string
	fmt.Println("Enter the path of csv file: ")
	fmt.Scanln(&filePath)
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvreader := csv.NewReader(f)
	data, err := csvreader.ReadAll()
	if err != nil {
		log.Fatal("Unable to read the csv file")
	}
	table := parse_csv(data)
	for _, ele := range table {
		for _, curr := range ele.row {
			fmt.Printf("\t%s", curr)
		}
		fmt.Printf("\n")
	}
}
