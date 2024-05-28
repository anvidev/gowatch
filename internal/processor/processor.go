package processor

import (
	"encoding/csv"
	"os"
	"strings"

	"github.com/andreasgylche/gowatch/internal/poster"
)

func ProcessCSV(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	csvData := formatCSVData(records)
	poster.PostData(csvData)

	return nil
}

func formatCSVData(records [][]string) string {
	delimeter := ";"
	var sb strings.Builder
	for _, record := range records {
		sb.WriteString(strings.Join(record, delimeter))
		sb.WriteString("\n")
	}
	return sb.String()
}
