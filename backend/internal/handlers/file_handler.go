package handlers

import (
    "io"
    "net/http"
    "secureshare/internal/models"
    "secureshare/internal/services"
    "strconv"

    "github.com/gin-gonic/gin"
)

type FileHandler struct {
    fileService *services.FileService
}

func NewFileHandler(fileService *services.FileService) *FileHandler {
    return &FileHandler{fileService: fileService}
}

func (h *FileHandler) UploadFile(c *gin.Context) {
    file, header, err := c.Request.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
        return
    }
    defer file.Close()

    data, err := io.ReadAll(file)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
        return
    }

    expiryHours, _ := strconv.Atoi(c.PostForm("expiry_hours"))
    if expiryHours == 0 {
        expiryHours = 24
    }

    maxDownloads, _ := strconv.Atoi(c.PostForm("max_downloads"))
    isOneTime := c.PostForm("is_one_time") == "true"

    req := &models.UploadRequest{
        ExpiryHours:  expiryHours,
        MaxDownloads: maxDownloads,
        IsOneTime:    isOneTime,
    }

    response, err := h.fileService.UploadFile(c.Request.Context(), header.Filename, data, req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, response)
}

func (h *FileHandler) DownloadFile(c *gin.Context) {
    fileID := c.Param("id")
    
    data, metadata, err := h.fileService.DownloadFile(c.Request.Context(), fileID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.Header("Content-Disposition", "attachment; filename="+metadata.OriginalName)
    c.Header("Content-Type", metadata.ContentType)
    c.Data(http.StatusOK, metadata.ContentType, data)
}

func (h *FileHandler) ListFiles(c *gin.Context) {
    files, err := h.fileService.ListFiles(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"files": files})
}

func (h *FileHandler) DeleteFile(c *gin.Context) {
    fileID := c.Param("id")
    
    err := h.fileService.DeleteFile(c.Request.Context(), fileID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}