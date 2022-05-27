package main

import (
	"context"
	"github.com/KL-Engineering/common-log/log"
	"github.com/KL-Engineering/tracecontext"
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
	initLogger()
	log.Info(ctx, ">>>>>>>>>> stm build start >>>>>>>>>>>>")
	log.Info(ctx, "<<<<<<<<<< stm build ended <<<<<<<<<<<<")
}
