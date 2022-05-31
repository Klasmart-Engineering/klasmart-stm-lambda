package model

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"kidsloop-stm-lambda/entity"
	"os"
	"strconv"
	"testing"
)

func TestCSV(t *testing.T) {
	csvFile, err := os.Open("./curriculum.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)

	rows, err := csvReader.ReadAll() // `rows` is of type [][]string
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) == 0 {
		// t.
		return
	}
	curriculums := make([]*entity.Curriculum, 0, len(rows)-1)
	for i, row := range rows {
		if i == 0 {
			fmt.Println("row:", i, row)
			continue
		}
		var curriculum entity.Curriculum
		curriculum.ID = row[0]
		curriculum.No, _ = strconv.Atoi(row[1])
		curriculum.Thumbnail = row[2]
		curriculum.Description = row[3]
		curriculums = append(curriculums, &curriculum)
	}

	result, err := json.Marshal(curriculums)
	fmt.Println("result:", string(result))
}
