package model

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"kidsloop-stm-lambda/config"
	"kidsloop-stm-lambda/entity"
	"sync"
)

type IJsonWrite interface {
	Curriculums(ctx context.Context, curriculums []*entity.Curriculum) error
	Levels(ctx context.Context, levelMap map[string]*entity.Level) error
	Units(ctx context.Context, unitMap map[string]*entity.Unit) error
	LessonPlan(ctx context.Context, lessonPlanMap map[string]*entity.LessonPlan) error
}

var (
	_jsonWriter     IJsonWrite
	_jsonWriterOnce sync.Once
)

func GetJsonWriter(ctx context.Context) IJsonWrite {
	_jsonWriterOnce.Do(func() {
		_jsonWriter = &JsonS3Writer{
			svc:    s3.New(session.New()),
			bucket: aws.String(config.Get().DestinationS3.Bucket),
			prefix: "test",
		}
	})
	return _jsonWriter
}
