package entity

type BaseField struct {
	ID          string `json:"id"`
	No          int    `json:"no,omitempty"`
	Name        string `json:"name"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
}

type Curriculum struct {
	BaseField
}

type Level struct {
	BaseField
}

type LessonPlan struct {
	BaseField
	ContentID string `json:"content_id"`
}

type Unit struct {
	BaseField
	LessonPlans []*LessonPlan `json:"lesson_plans"`
}

type Material struct {
	BaseField
	ContentID string `json:"content_id"`
	Data      string `json:"data"`
}
