package storage

import (
    "context"
    "io"
    "log"
    "secureshare/internal/config"

    "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
)

type MinIOStorage struct {
    client *minio.Client
    bucket string
}

func NewMinIOClient(cfg *config.Config) *MinIOStorage {
    client, err := minio.New(cfg.MinIOEndpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(cfg.MinIOAccessKey, cfg.MinIOSecretKey, ""),
        Secure: false,
    })
    if err != nil {
        log.Fatal("Failed to connect to MinIO:", err)
    }

    ctx := context.Background()
    exists, err := client.BucketExists(ctx, cfg.MinIOBucket)
    if err != nil {
        log.Fatal("Error checking bucket:", err)
    }

    if !exists {
        err = client.MakeBucket(ctx, cfg.MinIOBucket, minio.MakeBucketOptions{})
        if err != nil {
            log.Fatal("Error creating bucket:", err)
        }
    }

    return &MinIOStorage{
        client: client,
        bucket: cfg.MinIOBucket,
    }
}

func (s *MinIOStorage) UploadFile(ctx context.Context, objectName string, reader io.Reader, size int64, contentType string) error {
    _, err := s.client.PutObject(ctx, s.bucket, objectName, reader, size, minio.PutObjectOptions{
        ContentType: contentType,
    })
    return err
}

func (s *MinIOStorage) DownloadFile(ctx context.Context, objectName string) (*minio.Object, error) {
    return s.client.GetObject(ctx, s.bucket, objectName, minio.GetObjectOptions{})
}

func (s *MinIOStorage) DeleteFile(ctx context.Context, objectName string) error {
    return s.client.RemoveObject(ctx, s.bucket, objectName, minio.RemoveObjectOptions{})
}