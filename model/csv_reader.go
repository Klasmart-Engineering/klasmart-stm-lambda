package model

import (
	"context"
	"encoding/csv"
	"github.com/KL-Engineering/common-log/log"
	"kidsloop-stm-lambda/entity"
	"os"
	"strings"
	"sync"
)

type ICSVRead interface {
	Curriculums(ctx context.Context) ([]*entity.CSVCurriculum, error)
	Levels(ctx context.Context) ([]*entity.CSVLevel, error)
	Units(ctx context.Context) ([]*entity.CSVUnit, error)
	LevelUnitRelation(ctx context.Context) ([]*entity.CSVLevelUnitRelation, error)
	UnitLessonPlanRelation(ctx context.Context) ([]*entity.CSVUnitLessonPlanRelation, error)
}

type LocalCSVReader struct{}

var csvDir = "/Users/yanghui/kidsloop/kidsloop-stm-lambda/doc/csv"

func (locCSV LocalCSVReader) getData(ctx context.Context, filePath string) ([][]string, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Error(ctx, "open csv",
			log.Err(err),
			log.String("file", filePath))
		return nil, err
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)

	rows, err := csvReader.ReadAll() // `rows` is of type [][]string
	if err != nil {
		log.Error(ctx, "read csv",
			log.Err(err),
			log.String("file", filePath))
		return nil, err
	}
	return rows, nil
}

func (locCSV LocalCSVReader) Curriculums(ctx context.Context) ([]*entity.CSVCurriculum, error) {
	rows, err := locCSV.getData(ctx, strings.Join([]string{csvDir, entity.CurriculumCSV}, "/"))
	if err != nil {
		log.Error(ctx, "curriculum rows", log.Err(err))
		return nil, err
	}
	if len(rows) == 0 {
		log.Info(ctx, "curriculum zero rows", log.Err(err))
		return []*entity.CSVCurriculum{}, nil
	}
	curriculums := make([]*entity.CSVCurriculum, len(rows)-1)
	for i, r := range rows[1:] {
		base := entity.BaseField{
			ID:          r[0],
			Name:        r[1],
			Thumbnail:   r[2],
			Description: r[3],
		}
		curriculums[i] = &entity.CSVCurriculum{BaseField: base}
	}
	return curriculums, nil
}

func (locCSV LocalCSVReader) Levels(ctx context.Context) ([]*entity.CSVLevel, error) {
	rows, err := locCSV.getData(ctx, strings.Join([]string{csvDir, entity.LevelCSV}, "/"))
	if err != nil {
		log.Error(ctx, "level rows", log.Err(err))
		return nil, err
	}
	if len(rows) == 0 {
		log.Info(ctx, "level zero rows", log.Err(err))
		return []*entity.CSVLevel{}, nil
	}
	levels := make([]*entity.CSVLevel, len(rows)-1)
	for i, r := range rows[1:] {
		base := entity.BaseField{
			ID:          r[0],
			Name:        r[1],
			Thumbnail:   r[2],
			Description: r[3],
		}
		levels[i] = &entity.CSVLevel{BaseField: base, CurriculumID: r[4]}
	}
	return levels, nil
}
func (locCSV LocalCSVReader) Units(ctx context.Context) ([]*entity.CSVUnit, error) {
	rows, err := locCSV.getData(ctx, strings.Join([]string{csvDir, entity.UnitCSV}, "/"))
	if err != nil {
		log.Error(ctx, "unit rows", log.Err(err))
		return nil, err
	}
	if len(rows) == 0 {
		log.Info(ctx, "unit zero rows", log.Err(err))
		return []*entity.CSVUnit{}, nil
	}
	units := make([]*entity.CSVUnit, len(rows)-1)
	for i, r := range rows[1:] {
		base := entity.BaseField{
			ID:          r[0],
			Name:        r[1],
			Thumbnail:   r[2],
			Description: r[3],
		}
		units[i] = &entity.CSVUnit{BaseField: base}
	}
	return units, nil
}
func (locCSV LocalCSVReader) LevelUnitRelation(ctx context.Context) ([]*entity.CSVLevelUnitRelation, error) {
	rows, err := locCSV.getData(ctx, strings.Join([]string{csvDir, entity.LevelUnitCSV}, "/"))
	if err != nil {
		log.Error(ctx, "level unit relation rows", log.Err(err))
		return nil, err
	}
	if len(rows) == 0 {
		log.Info(ctx, "relation zero rows", log.Err(err))
		return []*entity.CSVLevelUnitRelation{}, nil
	}
	relations := make([]*entity.CSVLevelUnitRelation, len(rows)-1)
	for i, r := range rows[1:] {
		relation := entity.CSVLevelUnitRelation{
			LevelID: r[0],
			UnitID:  r[1],
		}
		relations[i] = &relation
	}
	return relations, nil
}
func (locCSV LocalCSVReader) UnitLessonPlanRelation(ctx context.Context) ([]*entity.CSVUnitLessonPlanRelation, error) {
	rows, err := locCSV.getData(ctx, strings.Join([]string{csvDir, entity.UnitLessonPlanCSV}, "/"))
	if err != nil {
		log.Error(ctx, "unit lesson_plan relation rows", log.Err(err))
		return nil, err
	}
	if len(rows) == 0 {
		log.Info(ctx, "relation zero rows", log.Err(err))
		return []*entity.CSVUnitLessonPlanRelation{}, nil
	}
	relations := make([]*entity.CSVUnitLessonPlanRelation, len(rows)-1)
	for i, r := range rows[1:] {
		relation := entity.CSVUnitLessonPlanRelation{
			UnitID:       r[0],
			LessonPlanID: r[1],
		}
		relations[i] = &relation
	}
	return relations, nil
}

var (
	_csvReader     ICSVRead
	_csvReaderOnce sync.Once
)

func GetCSVReader(ctx context.Context) ICSVRead {
	_csvReaderOnce.Do(func() {
		_csvReader = &LocalCSVReader{}
	})
	return _csvReader
}
