package entity

import "time"

const (
	CurriculumCSV     = "curriculums.csv"
	LevelUnitCSV      = "level_unit.csv"
	LevelCSV          = "levels.csv"
	UnitLessonPlanCSV = "unit_lesson_plan.csv"
	UnitCSV           = "units.csv"
)

const (
	CurriculumJSONKey  = "curriculums/curriculums.json"
	LevelsJSONKey      = "levels"
	UnitsJSONKey       = "units"
	LessonPlansJSONKey = "lesson_plans"
)

const (
	JsonContentType = "application/json"
	TextContentType = "text/plain"
)

const (
	TokenValidityPeriod = 24 * time.Hour
	TokenRefreshBefore  = 1 * time.Hour
)
