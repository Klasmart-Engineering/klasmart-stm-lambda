package config

import (
	"context"
	"github.com/KL-Engineering/common-log/log"
	"os"
)

type S3Config struct {
	Bucket string `json:"bucket"`
	Region string `json:"region"`
}

type Config struct {
	SourceS3           S3Config `json:"source_s3"`
	DestinationS3      S3Config `json:"destination_s3"`
	CmsEndpoint        string   `json:"cms_endpoint"`
	CloudFrontEndpoint string   `json:"cloud_front_endpoint"`
	H5pEndpoint        string   `json:"h5p_endpoint"`
}

var config = &Config{}

func LoadEnvConfig(ctx context.Context) {
	config.SourceS3.Region = os.Getenv("source_bucket_region")
	config.SourceS3.Bucket = os.Getenv("source_bucket")
	config.DestinationS3.Region = os.Getenv("destination_bucket_region")
	config.DestinationS3.Bucket = os.Getenv("destination_bucket")
	config.CmsEndpoint = os.Getenv("cms_endpoint")
	config.CloudFrontEndpoint = os.Getenv("cloud_front_endpoint")
	config.H5pEndpoint = os.Getenv("h5p_endpoint")

	log.Info(ctx, "load environment config", log.Any("config", config))
}

func Get() *Config {
	return config
}
