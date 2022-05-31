package model

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/KL-Engineering/common-log/log"
	"kidsloop-stm-lambda/entity"
	"os"
	"strconv"
)

type IBuilder interface {
	Build(input interface{}, output interface{}) error
}

type Builder struct {
}

func (Builder) curriculums(ctx context.Context, input interface{}) ([]*entity.Curriculum, error) {
	csvFile, err := os.Open("../doc/csv/curriculum.csv")
	if err != nil {
		log.Error(ctx, "curriculums",
			log.Err(err),
			log.Any("input", input))
		return nil, err
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)

	rows, err := csvReader.ReadAll() // `rows` is of type [][]string
	if err != nil {
		log.Error(ctx, "curriculums",
			log.Err(err),
			log.Any("input", input))
		return nil, err
	}
	if len(rows) == 0 {
		log.Warn(ctx, "curriculums: empty",
			log.Any("input", input))
		return []*entity.Curriculum{}, nil
	}
	curriculums := make([]*entity.Curriculum, 0, len(rows)-1)
	for i, row := range rows {
		if i == 0 {
			log.Warn(ctx, "curriculums: empty",
				log.Any("input", input),
				log.Strings("row:", row))
			continue
		}
		var curriculum entity.Curriculum
		curriculum.ID = row[0]
		curriculum.No, _ = strconv.Atoi(row[1])
		curriculum.Thumbnail = row[2]
		curriculum.Description = row[3]
		curriculums = append(curriculums, &curriculum)
	}
	return curriculums, nil
}

func (b Builder) levels(ctx context.Context, input interface{}, output interface{}) (map[string][]*entity.Level, error) {
	csvFile, err := os.Open("../doc/csv/level.csv")
	if err != nil {
		log.Error(ctx, "levels",
			log.Err(err),
			log.Any("input", input))
		return nil, err
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)

	rows, err := csvReader.ReadAll() // `rows` is of type [][]string
	if err != nil {
		log.Error(ctx, "curriculums",
			log.Err(err),
			log.Any("input", input))
		return nil, err
	}
	if len(rows) == 0 {
		log.Warn(ctx, "curriculums: empty",
			log.Any("input", input))
		return nil, nil
	}
	levelMap := make(map[string][]*entity.Level)
	for i, row := range rows {
		if i == 0 {
			log.Warn(ctx, "curriculums: empty",
				log.Any("input", input),
				log.Strings("row:", row))
			continue
		}
		var level entity.Level
		level.ID = row[0]
		level.No, _ = strconv.Atoi(row[1])
		level.Thumbnail = row[2]
		level.Description = row[3]
		if len(row) == 5 {
			continue
		}
		for _, curriculum := range row[5:len(row)] {
			levelMap[curriculum] = append(levelMap[curriculum], &level)
		}
	}
	return levelMap, nil
}

func (b Builder) units(ctx context.Context, input interface{}, output interface{}) () {

}

func (b Builder) Build(ctx context.Context, input interface{}, output interface{}) error {
	// 1.get change data from s3
	// 2.get data from kidsloop2
	// 3.build to json
	curriculums, err := b.curriculums(ctx, input)
	if err != nil {
		log.Error(ctx, "Build: curriculums failed",
			log.Err(err),
			log.Any("input", input))
		return err
	}
	curriculumsJson, err := json.Marshal(curriculums)
	if err != nil {
		log.Error(ctx, "Build: Marshal failed",
			log.Err(err),
			log.Any("input", input))
		return err
	}
	fmt.Println("curriculum:", string(curriculumsJson))
	return nil
}
