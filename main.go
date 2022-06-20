package main

import (
	"context"
	"github.com/KL-Engineering/common-log/log"
	"github.com/KL-Engineering/tracecontext"
	"github.com/aws/aws-lambda-go/lambda"
	"kidsloop-stm-lambda/config"
	"kidsloop-stm-lambda/model"
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

func LambdaHandler(ctx context.Context) error {
	//result, err := model.GetCSVReader(ctx).Curriculums(ctx)
	err := model.GetBuilder(ctx).Build(ctx, nil, nil)
	if err != nil {
		log.Error(ctx, "json build", log.Err(err))
		return err
	}
	return nil
}
