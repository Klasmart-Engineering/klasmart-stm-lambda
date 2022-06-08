package main

import (
	"context"
	"github.com/KL-Engineering/common-log/log"
	"github.com/KL-Engineering/tracecontext"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"kidsloop-stm-lambda/config"
)

func initLogger() {
	logger := log.New(log.WithDynamicFields(func(ctx context.Context) (fields []log.Field) {
		badaCtx, ok := tracecontext.GetTraceContext(ctx)
		if !ok {
			return
		}

		if badaCtx.CurrTid != "" {
			fields = append(fields, log.String("currTid", badaCtx.CurrTid))
		}

		if badaCtx.PrevTid != "" {
			fields = append(fields, log.String("prevTid", badaCtx.PrevTid))
		}

		if badaCtx.EntryTid != "" {
			fields = append(fields, log.String("entryTid", badaCtx.EntryTid))
		}

		return
	}))
	log.ReplaceGlobals(logger)
}

func main() {
	ctx := context.Background()
	config.LoadEnvConfig(ctx)
	initLogger()
	log.Info(ctx, ">>>>>>>>>> stm build start >>>>>>>>>>>>")
	lambda.Start(LambdaHandler)
	log.Info(ctx, "<<<<<<<<<< stm build ended <<<<<<<<<<<<")
}

func LambdaHandler(ctx context.Context) (int, error) {
	var invokeCount = 0
	var myObjects []*s3.Object

	svc := s3.New(session.New())
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(config.Get().SourceS3.Bucket),
	}
	result, err := svc.ListObjectsV2(input)
	if err != nil {
		log.Error(ctx, "list objects", log.Err(err))
		return 0, err
	}
	myObjects = result.Contents
	invokeCount = invokeCount + 1
	log.Info(ctx, "lambda handler",
		log.Any("objects", myObjects),
		log.Int("count", invokeCount))
	return invokeCount, nil
}
