package s3

import (
	"context"
	"errors"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	s3aws "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go"

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

func (c S3Client) GetFile(bucket string, key string) (s3aws.GetObjectOutput, error) {
	logger := logging.BuildLogger("GetFile")

	_, err := c.rawS3.HeadBucket(c.ctx, &s3aws.HeadBucketInput{
		Bucket: &bucket,
	})

	if err != nil {
		var ogerr smithy.APIError
		errMsg := ""
		if errors.As(err, &ogerr) {
			if ogerr.ErrorCode() == "NotFound" {
				errMsg = "Bucket not found"
			}
			logger.Error(errMsg, slog.Any("code", ogerr.ErrorCode()))
		} else {
			logger.Error("Unknown error", slog.Any("error", err))
		}

		return s3aws.GetObjectOutput{}, errors.New(errMsg)
	}

	data, err := c.rawS3.GetObject(ctx, &s3aws.GetObjectInput{
		Key:    &key,
		Bucket: &bucket,
	})
	if err != nil {
		var ogerr smithy.APIError
		errMsg := ""
		if errors.As(err, &ogerr) {
			if ogerr.ErrorCode() == "NoSuchKey" {
				errMsg = "File not found"
			}
			logger.Error(errMsg, slog.Any("code", ogerr.ErrorCode()))
		} else {
			logger.Error("Unknown error", slog.Any("error", err))
		}

		return s3aws.GetObjectOutput{}, errors.New(errMsg)
	}

	return *data, nil
}
