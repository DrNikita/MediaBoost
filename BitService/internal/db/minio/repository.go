package minio

import (
	"bit/internal/config"
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type minioStorage struct {
	client *minio.Client
}

func NewMinioStorage() *minioStorage {
	return &minioStorage{}
}

func (ms *minioStorage) InitMinioStorage(ctx context.Context, cfg config.MinioConfig) error {
	client, err := minio.New(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port), &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.RootUser, cfg.RootPassword, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		return err
	}

	ms.client = client

	exists, err := ms.client.BucketExists(ctx, cfg.BucketName)
	if err != nil {
		return err
	}
	if !exists {
		err := ms.client.MakeBucket(ctx, cfg.BucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}

func (ms *minioStorage) SaveBit() error {
	return nil
}
