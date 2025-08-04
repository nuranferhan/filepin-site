package models

import (
    "time"
)

type FileMetadata struct {
    ID           string    `json:"id"`
    OriginalName string    `json:"original_name"`
    Size         int64     `json:"size"`
    ContentType  string    `json:"content_type"`
    UploadTime   time.Time `json:"upload_time"`
    ExpiresAt    time.Time `json:"expires_at"`
    DownloadCount int      `json:"download_count"`
    MaxDownloads  int      `json:"max_downloads"`
    IsOneTime     bool     `json:"is_one_time"`
    EncryptionKey string   `json:"-"`
}

type UploadRequest struct {
    ExpiryHours   int  `json:"expiry_hours"`
    MaxDownloads  int  `json:"max_downloads"`
    IsOneTime     bool `json:"is_one_time"`
}

type UploadResponse struct {
    FileID     string `json:"file_id"`
    ShareLink  string `json:"share_link"`
    ExpiresAt  string `json:"expires_at"`
}