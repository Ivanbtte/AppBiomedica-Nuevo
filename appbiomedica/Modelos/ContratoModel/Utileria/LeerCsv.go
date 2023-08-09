package Utileria

import (
	"encoding/csv"
	"os"
)

func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	r := csv.NewReader(f)
	r.Comma = '\t'
	r.FieldsPerRecord = 12
	lines, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
