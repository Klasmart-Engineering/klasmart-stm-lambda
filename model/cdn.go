package model

import (
	"context"
	"github.com/KL-Engineering/common-log/log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"kidsloop-stm-lambda/config"
	"sync"
)

type IContentDeliveryNetwork interface {
	RefreshAll(ctx context.Context) error
}

type AWSCloudFront struct {
	svc            *cloudfront.CloudFront
	distributionID string
}

func (cloudFront *AWSCloudFront) RefreshAll(ctx context.Context) error {
	input := &cloudfront.CreateInvalidationInput{
		DistributionId: aws.String(cloudFront.distributionID),
		InvalidationBatch: &cloudfront.InvalidationBatch{
			Paths: &cloudfront.Paths{
				Items: []*string{aws.String("/*")},
			},
		},
	}

	result, err := cloudFront.svc.CreateInvalidation(input)
	if err != nil {
		log.Error(ctx, "create invalidation", log.Any("distribution_id", cloudFront.distributionID), log.Any("input", input))
		return err
	}

	log.Info(ctx, "refresh", log.Any("result", result))
	return nil
}

var (
	_cdn     IContentDeliveryNetwork
	_cdnOnce sync.Once
)

func GetContentDeliveryNetwork(ctx context.Context) IContentDeliveryNetwork {
	_cdnOnce.Do(func() {
		_cdn = &AWSCloudFront{
			svc:            cloudfront.New(session.New()),
			distributionID: config.Get().CloudFront.DistributionID,
		}
	})
	return _cdn
}
