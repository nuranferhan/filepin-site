package services

import (
    "bytes"
    "context"
    "errors"
    "fmt"
    "io"
    "secureshare/internal/crypto"
    "secureshare/internal/models"
    "secureshare/internal/repository"
    "secureshare/internal/storage"
    "time"

    "github.com/google/uuid"
)

type FileService struct {
    storage    *storage.MinIOStorage
    repository *repository.FileRepository
    encryption *crypto.AESEncryption
}

func NewFileService(storage *storage.MinIOStorage, repo *repository.FileRepository) *FileService {
    encryption := crypto.NewAESEncryption("your-secret-encryption-key-32-chars")
    return &FileService{
        storage:    storage,
        repository: repo,
        encryption: encryption,
    }
}

func (s *FileService) UploadFile(ctx context.Context, filename string, data []byte, req *models.UploadRequest) (*models.UploadResponse, error) {
    fileID := uuid.New().String()
    
    encryptedData, err := s.encryption.Encrypt(data)
    if err != nil {
        return nil, fmt.Errorf("encryption failed: %w", err)
    }

    objectName := fmt.Sprintf("encrypted_%s", fileID)
    reader := bytes.NewReader(encryptedData)
    
    err = s.storage.UploadFile(ctx, objectName, reader, int64(len(encryptedData)), "application/octet-stream")
    if err != nil {
        return nil, fmt.Errorf("upload failed: %w", err)
    }

    expiresAt := time.Now().Add(time.Duration(req.ExpiryHours) * time.Hour)
    metadata := &models.FileMetadata{
        ID:           fileID,
        OriginalName: filename,
        Size:         int64(len(data)),
        ContentType:  "application/octet-stream",
        UploadTime:   time.Now(),
        ExpiresAt:    expiresAt,
        MaxDownloads: req.MaxDownloads,
        IsOneTime:    req.IsOneTime,
    }

    err = s.repository.SaveFileMetadata(ctx, metadata)
    if err != nil {
        return nil, fmt.Errorf("metadata save failed: %w", err)
    }

    shareLink := fmt.Sprintf("/download/%s", fileID)
    return &models.UploadResponse{
        FileID:    fileID,
        ShareLink: shareLink,
        ExpiresAt: expiresAt.Format(time.RFC3339),
    }, nil
}

func (s *FileService) DownloadFile(ctx context.Context, fileID string) ([]byte, *models.FileMetadata, error) {
    metadata, err := s.repository.GetFileMetadata(ctx, fileID)
    if err != nil {
        return nil, nil, errors.New("file not found")
    }

    if time.Now().After(metadata.ExpiresAt) {
        return nil, nil, errors.New("file expired")
    }

    if metadata.MaxDownloads > 0 && metadata.DownloadCount >= metadata.MaxDownloads {
        return nil, nil, errors.New("download limit exceeded")
    }

    objectName := fmt.Sprintf("encrypted_%s", fileID)
    obj, err := s.storage.DownloadFile(ctx, objectName)
    if err != nil {
        return nil, nil, fmt.Errorf("download failed: %w", err)
    }
    defer obj.Close()

    encryptedData, err := io.ReadAll(obj)
    if err != nil {
        return nil, nil, fmt.Errorf("read failed: %w", err)
    }

    decryptedData, err := s.encryption.Decrypt(encryptedData)
    if err != nil {
        return nil, nil, fmt.Errorf("decryption failed: %w", err)
    }

    metadata.DownloadCount++
    s.repository.UpdateFileMetadata(ctx, metadata)

    if metadata.IsOneTime {
        s.DeleteFile(ctx, fileID)
    }

    return decryptedData, metadata, nil
}

func (s *FileService) DeleteFile(ctx context.Context, fileID string) error {
    objectName := fmt.Sprintf("encrypted_%s", fileID)
    s.storage.DeleteFile(ctx, objectName)
    return s.repository.DeleteFileMetadata(ctx, fileID)
}

func (s *FileService) ListFiles(ctx context.Context) ([]*models.FileMetadata, error) {
    return s.repository.ListFiles(ctx)
}