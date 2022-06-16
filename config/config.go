package config

import (
	"context"
	"github.com/KL-Engineering/common-log/log"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"os"
)

type LocalSourceConfig struct {
	UseLocalSource bool   `json:"use_local_source"`
	CSVDir         string `json:"local_csv_dir"`
	JSONDir        string `json:"local_json_dir"`
}

type S3Config struct {
	Bucket string `json:"bucket"`
	Region string `json:"region"`
	Prefix string `json:"prefix"`
}

type CloudFrontConfig struct {
	Endpoint       string `json:"endpoint"`
	DistributionID string `json:"distribution_id"`
}

type Config struct {
	LocalSource   LocalSourceConfig `json:"local_source"`
	SourceS3      S3Config          `json:"source_s3"`
	DestinationS3 S3Config          `json:"destination_s3"`

	CmsEndpoint  string      `json:"cms_endpoint"`
	CmsAccessKey interface{} `json:"cms_access_key"`
	//CloudFrontEndpoint string           `json:"cloud_front_endpoint"`
	CloudFront  CloudFrontConfig `json:"cloud_front"`
	H5pEndpoint string           `json:"h5p_endpoint"`
}

var config = &Config{}

func LoadEnvConfig(ctx context.Context) {
	config.LocalSource.UseLocalSource = os.Getenv("use_local_source") == "true"
	config.LocalSource.CSVDir = os.Getenv("local_csv_dir")
	config.LocalSource.JSONDir = os.Getenv("local_json_dir")
	config.SourceS3.Region = os.Getenv("source_bucket_region")
	config.SourceS3.Bucket = os.Getenv("source_bucket")
	config.DestinationS3.Region = os.Getenv("destination_bucket_region")
	config.DestinationS3.Bucket = os.Getenv("destination_bucket")
	config.DestinationS3.Prefix = os.Getenv("destination_bucket_prefix")
	if config.DestinationS3.Prefix == "" {
		config.DestinationS3.Prefix = "test"
	}
	config.CmsEndpoint = os.Getenv("cms_endpoint")
	config.CloudFront.Endpoint = os.Getenv("cloud_front_endpoint")
	config.CloudFront.DistributionID = os.Getenv("cloud_front_distribution_id")
	config.H5pEndpoint = os.Getenv("h5p_endpoint")

	privateKeyData, err := ioutil.ReadFile(os.Getenv("stm_private_key_path"))
	if err != nil {
		log.Panic(ctx, "reade private key file", log.Err(err))
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		log.Panic(ctx, "parse private key", log.Err(err))
	}
	config.CmsAccessKey = privateKey

	config.validate(ctx)

	log.Info(ctx, "load environment config", log.Any("config", config))
}

func (c Config) validate(ctx context.Context) {
	if c.LocalSource.UseLocalSource {
		if c.LocalSource.JSONDir == "" {
			log.Panic(ctx, "need json dir", log.String("dir", c.LocalSource.JSONDir))
		}
		if c.LocalSource.CSVDir == "" {
			log.Panic(ctx, "need csv dir", log.String("dir", c.LocalSource.CSVDir))
		}
		return
	}

	if c.SourceS3.Bucket == "" || c.SourceS3.Region == "" {
		log.Panic(ctx, "source s3 not correct config", log.Any("source_s3", c.SourceS3))
	}
	if c.DestinationS3.Bucket == "" || c.DestinationS3.Region == "" {
		log.Panic(ctx, "destination s3 not correct config", log.Any("destination_s3", c.DestinationS3))
	}
	if c.CmsEndpoint == "" {
		log.Panic(ctx, "cms endpoint not correct config", log.String("cmd_endpoint", c.CmsEndpoint))
	}
	if c.CloudFront.DistributionID == "" {
		log.Panic(ctx, "cloud_front not correct config", log.Any("cloud_front", c.CloudFront))
	}
}
func Get() *Config {
	return config
}
