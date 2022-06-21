package model

import (
	"context"
	"encoding/csv"
	"github.com/KL-Engineering/common-log/log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"kidsloop-stm-lambda/entity"
)

type CSVS3Reader struct {
	svc    *s3.S3
	bucket *string
}

func (csvS3 CSVS3Reader) getData(ctx context.Context, key string) ([][]string, error) {
	input := &s3.GetObjectInput{
		Bucket: csvS3.bucket,
		Key:    aws.String(key),
	}
	result, err := csvS3.svc.GetObject(input)
	if err != nil {
		log.Error(ctx, "get object", log.Err(err), log.String("bucket", *csvS3.bucket), log.String("key", key))
		return nil, err
	}

	csvReader := csv.NewReader(result.Body)
	defer result.Body.Close()

	rows, err := csvReader.ReadAll() // `rows` is of type [][]string
	if err != nil {
		log.Error(ctx, "read csv",
			log.Err(err),
			log.String("key", key))
		return nil, err
	}
	return rows, nil
}

func (csvS3 CSVS3Reader) Curriculums(ctx context.Context) ([]*entity.CSVCurriculum, error) {
	rows, err := csvS3.getData(ctx, entity.CurriculumCSV)
	if err != nil {
		log.Error(ctx, "curriculum rows", log.Err(err))
		return nil, err
	}
	if len(rows) == 0 {
		log.Info(ctx, "curriculum zero rows", log.Err(err))
		return []*entity.CSVCurriculum{}, nil
	}
	exists := make(map[string]bool)
	var curriculums []*entity.CSVCurriculum
	for i, r := range rows[1:] {
		if exists[r[0]] {
			log.Info(ctx, "curriculum already exists", log.Int("index", i), log.Strings("curriculum", r))
			continue
		}
		base := entity.BaseField{
			ID:          r[0],
			Name:        r[1],
			Thumbnail:   r[2],
			Description: r[3],
		}
		curriculums = append(curriculums, &entity.CSVCurriculum{BaseField: base})
	}
	return curriculums, nil
}

func (csvS3 CSVS3Reader) Levels(ctx context.Context) ([]*entity.CSVLevel, error) {
	rows, err := csvS3.getData(ctx, entity.LevelCSV)
	if err != nil {
		log.Error(ctx, "level rows", log.Err(err))
		return nil, err
	}
	if len(rows) == 0 {
		log.Info(ctx, "level zero rows", log.Err(err))
		return []*entity.CSVLevel{}, nil
	}
	exists := make(map[string]bool)
	var levels []*entity.CSVLevel
	for i, r := range rows[1:] {
		if exists[r[0]] {
			log.Info(ctx, "level already exists", log.Int("index", i), log.Strings("level", r))
			continue
		}
		base := entity.BaseField{
			ID:          r[0],
			Name:        r[1],
			Thumbnail:   r[2],
			Description: r[3],
		}
		levels = append(levels, &entity.CSVLevel{BaseField: base, CurriculumID: r[4]})
	}
	return levels, nil
}
func (csvS3 CSVS3Reader) Units(ctx context.Context) ([]*entity.CSVUnit, error) {
	rows, err := csvS3.getData(ctx, entity.UnitCSV)
	if err != nil {
		log.Error(ctx, "unit rows", log.Err(err))
		return nil, err
	}
	if len(rows) == 0 {
		log.Info(ctx, "unit zero rows", log.Err(err))
		return []*entity.CSVUnit{}, nil
	}
	exists := make(map[string]bool)
	var units []*entity.CSVUnit
	for i, r := range rows[1:] {
		if exists[r[0]] {
			log.Info(ctx, "unit already exists", log.Int("index", i), log.Strings("unit", r))
			continue
		}
		base := entity.BaseField{
			ID:          r[0],
			Name:        r[1],
			Thumbnail:   r[2],
			Description: r[3],
		}
		units = append(units, &entity.CSVUnit{BaseField: base})
	}
	return units, nil
}
func (csvS3 CSVS3Reader) LevelUnitRelation(ctx context.Context) ([]*entity.CSVLevelUnitRelation, error) {
	rows, err := csvS3.getData(ctx, entity.LevelUnitCSV)
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
func (csvS3 CSVS3Reader) UnitLessonPlanRelation(ctx context.Context) ([]*entity.CSVUnitLessonPlanRelation, error) {
	rows, err := csvS3.getData(ctx, entity.UnitLessonPlanCSV)
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
