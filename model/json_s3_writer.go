package model

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/KL-Engineering/common-log/log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"kidsloop-stm-lambda/entity"
	"strings"
)

type JsonS3Writer struct {
	svc    *s3.S3
	bucket *string
	prefix string
}

func (jsonS3 JsonS3Writer) writeData(ctx context.Context, key string, data interface{}) error {
	buff, err := json.Marshal(data)
	if err != nil {
		log.Error(ctx, "marshal data", log.Err(err), log.String("key", key), log.Any("data", data))
		return err
	}

	input := s3.PutObjectInput{
		Bucket: jsonS3.bucket,
		Key:    aws.String(key),
		Body:   aws.ReadSeekCloser(bytes.NewBuffer(buff)),
	}
	result, err := jsonS3.svc.PutObject(&input)
	if err != nil {
		log.Error(ctx, "put object",
			log.Err(err),
			log.String("bucket", *jsonS3.bucket),
			log.String("key", key),
			log.String("data", string(buff)))
		return err
	}
	log.Info(ctx, "put object", log.String("key", key), log.String("result", result.GoString()))
	return nil
}

func (jsonS3 JsonS3Writer) Curriculums(ctx context.Context, curriculums []*entity.Curriculum) error {
	if len(curriculums) == 0 {
		log.Info(ctx, "curriculum  slice is empty")
		return nil
	}
	err := jsonS3.writeData(ctx, strings.Join([]string{jsonS3.prefix, entity.CurriculumJSONKey}, "/"), curriculums)
	if err != nil {
		return err
	}
	return nil
}

func (jsonS3 JsonS3Writer) Levels(ctx context.Context, levelMap map[string]*entity.Level) error {
	if len(levelMap) == 0 {
		log.Info(ctx, "level map is empty")
		return nil
	}

	for k, v := range levelMap {
		err := jsonS3.writeData(ctx, strings.Join([]string{jsonS3.prefix, entity.LevelsJSONKey, k + ".json"}, "/"), v)
		if err != nil {
			log.Error(ctx, "write level", log.Err(err), log.String("key", k))
			return err
		}
	}
	return nil
}

func (jsonS3 JsonS3Writer) Units(ctx context.Context, unitMap map[string]*entity.Unit) error {
	if len(unitMap) == 0 {
		log.Info(ctx, "unit map is empty")
		return nil
	}

	for k, v := range unitMap {
		err := jsonS3.writeData(ctx, strings.Join([]string{jsonS3.prefix, entity.UnitsJSONKey, k + ".json"}, "/"), v)
		if err != nil {
			log.Error(ctx, "write unit", log.Err(err), log.String("key", k))
			return err
		}
	}
	return nil
}

func (jsonS3 JsonS3Writer) LessonPlan(ctx context.Context, lessonPlanMap map[string]*entity.LessonPlan) error {
	if len(lessonPlanMap) == 0 {
		log.Info(ctx, "lesson_plan map is empty")
		return nil
	}

	for k, v := range lessonPlanMap {
		err := jsonS3.writeData(ctx, strings.Join([]string{jsonS3.prefix, entity.LessonPlansJSONKey, k + ".json"}, "/"), v)
		if err != nil {
			log.Error(ctx, "write lesson_plan", log.Err(err), log.String("key", k))
			return err
		}
	}
	return nil
}
