package model

import (
	"context"
	"kidsloop-stm-lambda/config"
	"testing"
)

func TestKidsloopProvider_MapContents(t *testing.T) {
	ctx := context.Background()
	config.LoadEnvConfig(ctx)
	IDs := []string{
		"628da79e552ba3b9994c9200",
		"6257868a9456ed3fc792b775",
		"624419aace8e2cbaa66ca0f8",
	}
	result, err := mustKidsloopProvider(ctx).MapContents(ctx, IDs)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
