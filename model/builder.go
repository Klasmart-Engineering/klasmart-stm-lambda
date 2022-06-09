package model

import (
	"context"
	"github.com/KL-Engineering/common-log/log"
	"kidsloop-stm-lambda/entity"
	"kidsloop-stm-lambda/utils"
	"sync"
)

type IBuilder interface {
	Build(ctx context.Context, input interface{}, output interface{}) error
}

type Builder struct {
}

//func (b Builder) getLessonPlans(ctx context.Context, IDs []string) (map[string]*entity.LessonPlan, error) {
//	return GetContentProvider(ctx).MapContents(ctx, IDs)
//}

func (b Builder) Build(ctx context.Context, input interface{}, output interface{}) error {
	csvCurriculums, err := GetCSVReader(ctx).Curriculums(ctx)
	if err != nil {
		log.Error(ctx, "curriculums", log.Err(err))
		return err
	}
	csvLevels, err := GetCSVReader(ctx).Levels(ctx)
	if err != nil {
		log.Error(ctx, "levels", log.Err(err))
		return err
	}
	csvUnits, err := GetCSVReader(ctx).Units(ctx)
	if err != nil {
		log.Error(ctx, "units", log.Err(err))
		return err
	}
	csvLevelUnit, err := GetCSVReader(ctx).LevelUnitRelation(ctx)
	if err != nil {
		log.Error(ctx, "level unit relation", log.Err(err))
		return err
	}
	csvUnitLessonPlan, err := GetCSVReader(ctx).UnitLessonPlanRelation(ctx)
	if err != nil {
		log.Error(ctx, "unit lesson_plan relation", log.Err(err))
		return err
	}

	var lessonPlanIDs []string
	unitIDKeyLessonPlanIDMap := make(map[string][]string)
	for _, ul := range csvUnitLessonPlan {
		lessonPlanIDs = append(lessonPlanIDs, ul.LessonPlanID)
		unitIDKeyLessonPlanIDMap[ul.UnitID] = append(unitIDKeyLessonPlanIDMap[ul.UnitID], ul.LessonPlanID)
	}
	lessonPlanIDs = utils.SliceDeduplicationExcludeEmpty(lessonPlanIDs)
	lessonPlanMap, err := GetContentProvider(ctx).MapContents(ctx, lessonPlanIDs)
	if err != nil {
		log.Error(ctx, "lesson_plans", log.Err(err), log.Strings("ids", lessonPlanIDs))
		return err
	}

	unitMap := make(map[string]*entity.Unit)
	for _, u := range csvUnits {
		if _, ok := unitMap[u.ID]; ok {
			continue
		}
		unit := entity.Unit{BaseField: u.BaseField}
		for _, l := range unitIDKeyLessonPlanIDMap[u.ID] {
			unit.LessonPlans = append(unit.LessonPlans, &lessonPlanMap[l].BaseField)
		}
		unitMap[u.ID] = &unit
	}

	leveIDKeyUnitIDMap := make(map[string][]string)
	for _, lur := range csvLevelUnit {
		leveIDKeyUnitIDMap[lur.LevelID] = append(leveIDKeyUnitIDMap[lur.LevelID], lur.UnitID)
	}

	curriculumIDKeyLevel := make(map[string][]*entity.BaseField)
	levelMap := make(map[string]*entity.Level)
	for _, l := range csvLevels {
		if _, ok := levelMap[l.ID]; ok {
			continue
		}

		curriculumIDKeyLevel[l.CurriculumID] = append(curriculumIDKeyLevel[l.CurriculumID], &l.BaseField)

		level := entity.Level{BaseField: l.BaseField}
		level.Units = make([]*entity.Unit, len(leveIDKeyUnitIDMap[l.ID]))
		for i, u := range leveIDKeyUnitIDMap[l.ID] {
			if unitMap[u] == nil {
				log.Error(ctx, "unit not exists", log.String("unit", u))
				return entity.ErrRecordNotExist
			}
			level.Units[i] = unitMap[u]
		}
		levelMap[l.ID] = &level
	}

	curriculums := make([]*entity.Curriculum, len(csvCurriculums))
	for i, c := range csvCurriculums {
		curriculum := entity.Curriculum{BaseField: c.BaseField}
		curriculum.Levels = curriculumIDKeyLevel[c.ID]
		curriculums[i] = &curriculum
	}

	err = GetJsonWriter(ctx).Curriculums(ctx, curriculums)
	if err != nil {
		log.Error(ctx, "upload curriculums", log.Err(err))
		return err
	}
	err = GetJsonWriter(ctx).Levels(ctx, levelMap)
	if err != nil {
		log.Error(ctx, "upload levels", log.Err(err))
		return err
	}
	err = GetJsonWriter(ctx).Units(ctx, unitMap)
	if err != nil {
		log.Error(ctx, "upload units", log.Err(err))
		return err
	}
	err = GetJsonWriter(ctx).LessonPlan(ctx, lessonPlanMap)
	if err != nil {
		log.Error(ctx, "upload lesson_plan", log.Err(err))
		return err
	}
	return nil
}

var (
	_builder     IBuilder
	_builderOnce sync.Once
)

func GetBuilder(ctx context.Context) IBuilder {
	_builderOnce.Do(func() {
		_builder = &Builder{}
	})
	return _builder
}
