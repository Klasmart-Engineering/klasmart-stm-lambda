package model

import (
	"context"
	"encoding/json"
	"github.com/KL-Engineering/common-log/log"
	"io/ioutil"
	"kidsloop-stm-lambda/entity"
	"strings"
	"sync"
)

type IContent interface {
	MapContents(ctx context.Context, IDs []string) (map[string]*entity.LessonPlan, error)
}

type ContentProvider struct{}

func (cp ContentProvider) MapContents(ctx context.Context, IDs []string) (map[string]*entity.LessonPlan, error) {
	dir := "/Users/yanghui/kidsloop/kidsloop-stm-lambda/doc/json/lesson_plans"
	result := make(map[string]*entity.LessonPlan)
	for _, id := range IDs {
		data, err := ioutil.ReadFile(strings.Join([]string{dir, id + ".json"}, "/"))
		if err != nil {
			log.Error(ctx, "read lesson_plan", log.Err(err), log.String("id", id))
			return nil, err
		}
		var lessonPlan entity.LessonPlan
		err = json.Unmarshal(data, &lessonPlan)
		if err != nil {
			log.Error(ctx, "unmarshal lesson_plan", log.Err(err), log.String("data", string(data)))
			return nil, err
		}
		result[id] = &lessonPlan
	}
	return result, nil
}

var (
	_contentProvider IContent
	_contentOnce     sync.Once
)

func GetContentProvider(ctx context.Context) IContent {
	_contentOnce.Do(func() {
		_contentProvider = &ContentProvider{}
	})
	return _contentProvider
}
