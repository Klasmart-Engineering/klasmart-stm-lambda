package entity

type BaseField struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
}

type Curriculum struct {
	BaseField
}

type Level struct {
	BaseField
	CurriculumID string `json:"curriculum_id"`
}

type LevelUnitRelation struct {
	LevelID string `json:"level_id"`
	UnitID  string `json:"unit_id"`
}

type Unit struct {
	BaseField
}

type UnitLessonPlanRelation struct {
	UnitID       string `json:"unit_id"`
	LessonPlanID string `json:"lesson_plan_id"`
}
type LessonPlan struct {
	BaseField
	Materials []*Material `json:"materials"`
}

type Material struct {
	BaseField
	Data string `json:"data"`
}
