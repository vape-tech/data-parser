package data_parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadCSV(filePath string) ([]map[string]string, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var data []map[string]string
	for _, record := range records {
		row := make(map[string]string)
		for i, field := range record {
			row[strings.ToLower(record[i])] = strings.TrimSpace(field)
		}
		data = append(data, row)
	}

	return data, nil
}

func SplitString(s string) []string {
	return strings.Split(s, ",")
}

func IsNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func TrimWhitespace(s string) string {
	return strings.TrimSpace(s)
}

func TrimCSVString(s string) string {
	return strings.TrimRight(s, ",\n")
}

func ToSliceInt64(arr []string) ([]int64, error) {
	var result []int64
	for _, v := range arr {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}

	return result, nil
}

func WriteDataToCSV(filePath string, data [][]string) error {
	csvFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	for _, row := range data {
		err = writer.Write(row)
		if err != nil {
			return err
		}
		err = writer.Flush()
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteTestData(filePath string, data interface{}) error {
	switch data := data.(type) {
	case []interface{}:
		csvData := [][]string{}
		for _, item := range data {
			row := []string{}
			switch item.(type) {
			case map[string]interface{}:
				for k, v := range item.(map[string]interface{}) {
					row = append(row, fmt.Sprintf("%v", v))
				}
			case []interface{}:
				for _, v := range item.([]interface{}) {
					row = append(row, fmt.Sprintf("%v", v))
				}
			}
			csvData = append(csvData, row)
		}
		return WriteDataToCSV(filePath, csvData)
	default:
		return fmt.Errorf("invalid data type: %T", data)
	}
}