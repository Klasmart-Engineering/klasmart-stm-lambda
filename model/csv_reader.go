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

type ICSVRead interface {
	Curriculums(ctx context.Context) ([]*entity.CSVCurriculum, error)
	Levels(ctx context.Context) ([]*entity.CSVLevel, error)
	Units(ctx context.Context) ([]*entity.CSVUnit, error)
	LevelUnitRelation(ctx context.Context) ([]*entity.CSVLevelUnitRelation, error)
	UnitLessonPlanRelation(ctx context.Context) ([]*entity.CSVUnitLessonPlanRelation, error)
}

var (
	_csvReader     ICSVRead
	_csvReaderOnce sync.Once
)

func GetCSVReader(ctx context.Context) ICSVRead {
	_csvReaderOnce.Do(func() {
		//_csvReader = &CSVLocalReader{}
		_csvReader = &CSVS3Reader{
			svc:    s3.New(session.New()),
			bucket: aws.String(config.Get().SourceS3.Bucket),
		}
	})
	return _csvReader
}
