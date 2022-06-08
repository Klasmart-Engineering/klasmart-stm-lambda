package main

import (
	"context"
	"fmt"
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

var invokeCount = 0
var myObjects []*s3.Object

func init() {
	svc := s3.New(session.New())
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String("kidsloop-alpha-stm-data-intent-turtle"),
	}
	result, _ := svc.ListObjectsV2(input)
	myObjects = result.Contents
}

func LambdaHandler() (int, error) {
	invokeCount = invokeCount + 1
	fmt.Print(myObjects)
	return invokeCount, nil
}
