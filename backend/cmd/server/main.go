package main

import (
    "log"
    "secureshare/internal/config"
    "secureshare/internal/handlers"
    "secureshare/internal/middleware"
    "secureshare/internal/repository"
    "secureshare/internal/services"
    "secureshare/internal/storage"

    "github.com/gin-gonic/gin"
)

func main() {
    cfg := config.Load()

    minioClient := storage.NewMinIOClient(cfg)
    fileRepo := repository.NewFileRepository()
    fileService := services.NewFileService(minioClient, fileRepo)
    fileHandler := handlers.NewFileHandler(fileService)

    r := gin.Default()
    r.Use(middleware.CORSMiddleware())

    api := r.Group("/api/v1")
    {
        api.POST("/upload", fileHandler.UploadFile)
        api.GET("/download/:id", fileHandler.DownloadFile)
        api.GET("/files", fileHandler.ListFiles)
        api.DELETE("/files/:id", fileHandler.DeleteFile)
    }

    log.Printf("Server starting on port %s", cfg.ServerPort)
    r.Run(":" + cfg.ServerPort)
}