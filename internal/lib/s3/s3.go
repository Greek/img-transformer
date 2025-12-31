package s3

import (
	"context"
	"errors"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	s3aws "github.com/aws/aws-sdk-go-v2/service/s3"
	env "github.com/greek/img-transform/internal/lib/envloader"
	"github.com/greek/img-transform/internal/lib/logging"
)

type S3Client struct {
	ctx    context.Context
	rawS3  *s3aws.Client
	logger *slog.Logger
}

var ctx = context.Background()

func InitS3() *S3Client {
	logger := logging.BuildLogger("S3")
	cfg := env.GetEnv()

	awsCfg := aws.Config{
		Region: cfg.S3_REGION,
		Credentials: aws.NewCredentialsCache(
			aws.CredentialsProviderFunc(
				func(ctx context.Context) (aws.Credentials, error) {
					return aws.Credentials{
						AccessKeyID:     cfg.S3_ACCESS_KEY,
						SecretAccessKey: cfg.S3_SECRET_KEY,
					}, nil
				},
			),
		),
	}
	client := s3aws.NewFromConfig(awsCfg)

	return &S3Client{rawS3: client, ctx: ctx, logger: logger}
}

func (c S3Client) GetFile(bucket string, key string) (string, error) {
	input := &s3aws.HeadBucketInput{
		Bucket: &bucket,
	}

	_, err := c.rawS3.HeadBucket(c.ctx, input)
	if err != nil {
		errMsg := "failed to get bucket metadata"

		c.logger.Error(errMsg, slog.String("error", err.Error()), slog.String("bucket", bucket))
		return "", errors.New(errMsg)
	}

	return "mock data", nil
}
