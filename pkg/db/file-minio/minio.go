package file_minio

import (
	"context"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type MnConfig struct {
	Endpoint       string `yaml:"MINIO_ENDPOINT" env:"MINIO_ENDPOINT" env-default:"localhost:9000"`
	AccessKey      string `yaml:"MINIO_ACCESS_KEY" env:"MINIO_ACCESS_KEY" env-default:"minioadmin"`
	SecretKey      string `yaml:"MINIO_SECRET_KEY" env:"MINIO_SECRET_KEY" env-default:"minioadmin"`
	Region         string `yaml:"MINIO_REGION" env:"MINIO_REGION" env-default:"us-east-1"`
	UseSSL         bool   `yaml:"MINIO_USE_SSL" env:"MINIO_USE_SSL" env-default:"false"`
	ForcePathStyle bool   `yaml:"MINIO_FORCE_PATH_STYLE" env:"MINIO_FORCE_PATH_STYLE" env-default:"true"`
}

type Client struct {
	client     *s3.S3
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
}

type FileObject struct {
	Content      []byte
	ContentType  string
	Size         int64
	LastModified string
}

func New(cfg MnConfig) (*Client, error) {
	sess, err := session.NewSession(&aws.Config{
		Endpoint:         aws.String(cfg.Endpoint),
		Credentials:      credentials.NewStaticCredentials(cfg.AccessKey, cfg.SecretKey, ""),
		Region:           aws.String(cfg.Region),
		DisableSSL:       aws.Bool(!cfg.UseSSL),
		S3ForcePathStyle: aws.Bool(cfg.ForcePathStyle),
	})
	if err != nil {
		return nil, err
	}

	return &Client{
		client:     s3.New(sess),
		uploader:   s3manager.NewUploader(sess),
		downloader: s3manager.NewDownloader(sess),
	}, nil
}

func (c *Client) GetFile(ctx context.Context, bucket, key string) ([]byte, error) {
	buf := aws.NewWriteAtBuffer([]byte{})
	_, err := c.downloader.DownloadWithContext(ctx, buf, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *Client) GetFileWithMeta(ctx context.Context, bucket, key string) (*FileObject, error) {
	output, err := c.client.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	defer output.Body.Close()

	content, err := io.ReadAll(output.Body)
	if err != nil {
		return nil, err
	}

	contentType := ""
	if output.ContentType != nil {
		contentType = *output.ContentType
	}

	lastModified := ""
	if output.LastModified != nil {
		lastModified = output.LastModified.Format(time.RFC3339)
	}

	size := int64(0)
	if output.ContentLength != nil {
		size = *output.ContentLength
	}

	return &FileObject{
		Content:      content,
		ContentType:  contentType,
		Size:         size,
		LastModified: lastModified,
	}, nil
}
