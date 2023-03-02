package tasCSV

import (
	"encoding/csv"
	"log"
	"os"
)

/*
WriteToCSV is a function that writes data to a .csv file. It requires the desired output file name,
csv headers, number of rows/length of entries, and the data to be stored in each row.
*/
func WriteToCSV(outputFile string, headers []string, rowCount int, rowData [][]string) {

	// Create .csv file
	csvFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatalln("[!] Error creating file:", err)
	}

	csvWriter := csv.NewWriter(csvFile)

	var bodyData [][]string
	var headerData [][]string

	// Write headers as first line in .csv
	headerData = append(headerData, headers)
	csvWriter.WriteAll(headerData)

	// Write row for as many rows as specified in rowCount param
	for i := 0; i < rowCount; i++ {
		// Get the row from the array to the .csv file
		row := rowData[i]
		bodyData = append(bodyData, row)
	}

	// Write to CSV
	csvWriter.WriteAll(bodyData)

}
