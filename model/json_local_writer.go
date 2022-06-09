package model

import (
	"context"
	"encoding/json"
	"github.com/KL-Engineering/common-log/log"
	"io/ioutil"
	"kidsloop-stm-lambda/entity"
	"os"
	"strings"
)

type LocalJsonWriter struct {
	dir string
}

//var mockDir = "/Users/yanghui/kidsloop/kidsloop-stm-lambda/doc/json2"

func (jsonLoc LocalJsonWriter) Curriculums(ctx context.Context, curriculums []*entity.Curriculum) error {
	data, err := json.Marshal(curriculums)
	if err != nil {
		log.Error(ctx, "marshal curriculums", log.Err(err), log.Any("curriculums", curriculums))
		return err
	}
	filename := strings.Join([]string{jsonLoc.dir, entity.CurriculumJSONKey}, "/")
	err = ioutil.WriteFile(filename, data, os.ModePerm)
	if err != nil {
		log.Error(ctx, "write curriculums file", log.Err(err), log.String("data", string(data)))
		return err
	}
	return nil
}

func (jsonLoc LocalJsonWriter) Levels(ctx context.Context, levelMap map[string]*entity.Level) error {
	for k, v := range levelMap {
		data, err := json.Marshal(v)
		if err != nil {
			log.Error(ctx, "marshal level", log.Err(err), log.Any("level", v))
			return err
		}
		filename := strings.Join([]string{jsonLoc.dir, entity.LevelsJSONKey, k + ".json"}, "/")
		err = ioutil.WriteFile(filename, data, os.ModePerm)
		if err != nil {
			log.Error(ctx, "write level file", log.Err(err), log.String("data", string(data)))
			return err
		}
	}
	return nil
}

func (jsonLoc LocalJsonWriter) Units(ctx context.Context, unitMap map[string]*entity.Unit) error {
	for k, v := range unitMap {
		data, err := json.Marshal(v)
		if err != nil {
			log.Error(ctx, "marshal unit", log.Err(err), log.Any("unit", v))
			return err
		}
		filename := strings.Join([]string{jsonLoc.dir, entity.UnitsJSONKey, k + ".json"}, "/")
		err = ioutil.WriteFile(filename, data, os.ModePerm)
		if err != nil {
			log.Error(ctx, "write unit file", log.Err(err), log.String("data", string(data)))
			return err
		}
	}
	return nil
}

func (jsonLoc LocalJsonWriter) LessonPlan(ctx context.Context, lessonPlanMap map[string]*entity.LessonPlan) error {
	for k, v := range lessonPlanMap {
		data, err := json.Marshal(v)
		if err != nil {
			log.Error(ctx, "marshal lesson_plan", log.Err(err), log.Any("lesson_plan", v))
			return err
		}
		filename := strings.Join([]string{jsonLoc.dir, entity.LessonPlansJSONKey, k + ".json"}, "/")
		err = ioutil.WriteFile(filename, data, os.ModePerm)
		if err != nil {
			log.Error(ctx, "write lesson_plan file", log.Err(err), log.String("data", string(data)))
			return err
		}
	}
	return nil
}
