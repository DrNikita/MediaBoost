package minio

import (
	"bit/internal/config"
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type minioStorage struct {
	client *minio.Client
	cfg    config.MinioConfig
}

func NewMinioStorage(cfg config.MinioConfig) *minioStorage {
	return &minioStorage{
		cfg: cfg,
	}
}

func (ms *minioStorage) Connect(ctx context.Context) error {
	client, err := minio.New(fmt.Sprintf("%s:%s", ms.cfg.Host, ms.cfg.Port), &minio.Options{
		Creds:  credentials.NewStaticV4(ms.cfg.RootUser, ms.cfg.RootPassword, ""),
		Secure: ms.cfg.UseSSL,
	})
	if err != nil {
		return err
	}

	ms.client = client

	exists, err := ms.client.BucketExists(ctx, ms.cfg.BucketName)
	if err != nil {
		return err
	}
	if !exists {
		err = ms.client.MakeBucket(ctx, ms.cfg.BucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}

func (ms *minioStorage) SaveBit(ctx context.Context, data io.Reader, dataLen int64, fileName string) error {
	_, err := ms.client.PutObject(ctx, ms.cfg.BucketName, "/bit/{userId}/{fileName}", data, dataLen, minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (ms *minioStorage) SaveBitAsChild(ctx context.Context, data io.Reader, dataLen int64, fileName string, parentBit int64) error {
	_, err := ms.client.PutObject(ctx, ms.cfg.BucketName, "/bit/{parentBitId}/{userId}/{fileName}", data, dataLen, minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}
