package config

import (
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    ServerPort      string
    MinIOEndpoint   string
    MinIOAccessKey  string
    MinIOSecretKey  string
    MinIOBucket     string
    RedisAddr       string
    RedisPassword   string
    EncryptionKey   string
}

func Load() *Config {
    godotenv.Load()

    return &Config{
        ServerPort:      getEnv("SERVER_PORT", "8080"),
        MinIOEndpoint:   getEnv("MINIO_ENDPOINT", "localhost:9000"),
        MinIOAccessKey:  getEnv("MINIO_ACCESS_KEY", "minioadmin"),
        MinIOSecretKey:  getEnv("MINIO_SECRET_KEY", "minioadmin"),
        MinIOBucket:     getEnv("MINIO_BUCKET", "secureshare"),
        RedisAddr:       getEnv("REDIS_ADDR", "localhost:6379"),
        RedisPassword:   getEnv("REDIS_PASSWORD", ""),
        EncryptionKey:   getEnv("ENCRYPTION_KEY", "your-32-char-encryption-key-here"),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}