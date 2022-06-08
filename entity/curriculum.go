package entity

type BaseField struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
}

type CSVCurriculum struct {
	BaseField
}

type CSVLevel struct {
	BaseField
	CurriculumID string `json:"curriculum_id"`
}

type CSVLevelUnitRelation struct {
	LevelID string `json:"level_id"`
	UnitID  string `json:"unit_id"`
}

type CSVUnit struct {
	BaseField
}

type CSVUnitLessonPlanRelation struct {
	UnitID       string `json:"unit_id"`
	LessonPlanID string `json:"lesson_plan_id"`
}

type Curriculum struct {
	BaseField
	Levels []*BaseField `json:"levels"`
}

type Level struct {
	BaseField
	Units []*Unit `json:"units"`
}

type Unit struct {
	BaseField
	LessonPlans []*BaseField `json:"lesson_plans"`
}

type LessonPlan struct {
	BaseField
	Materials []*Material `json:"materials"`
}

type Material struct {
	BaseField
	Data string `json:"data"`
}
