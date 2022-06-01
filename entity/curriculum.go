package entity

type BaseField struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
}

type CurriculumLevels struct {
	ID    string   `json:"id"`
	Units []string `json:"units"`
}

type Curriculum struct {
	BaseField
	Levels []*CurriculumLevels `json:"levels"`
}

type Level struct {
	BaseField
}

type LessonPlan struct {
	BaseField
	ContentID string      `json:"content_id,omitempty"`
	Materials []*Material `json:"materials"`
}

type Unit struct {
	BaseField
	LessonPlans []*LessonPlan `json:"lesson_plans"`
}

type Material struct {
	BaseField
	ContentID string `json:"content_id,omitempty"`
	Data      string `json:"data"`
}
