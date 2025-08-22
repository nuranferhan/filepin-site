# FilePin - Modern Secure File Sharing Application

<div align="center">
 <img src="https://img.shields.io/badge/FilePin-Go%20%2B%20React-00ADD8?style=for-the-badge" alt="FilePin" />
 <img src="https://img.shields.io/badge/License-MIT-06B6D4?style=for-the-badge" alt="License" />
 <img src="https://img.shields.io/badge/React-18.0+-61DAFB?style=for-the-badge&logo=react" alt="React" />
 <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go" alt="Go" />
 <img src="https://img.shields.io/badge/MinIO-C72E49?style=for-the-badge&logo=minio" alt="MinIO" />
</div>

## Project Overview

<div align="center">
  <img width="85%" alt="Image" src="https://github.com/user-attachments/assets/67bcc5b4-2916-4b38-b9a9-32d6f9489aeb" />
</div>

**FilePin** is a comprehensive secure file sharing web application built with Go and React, representing a modern approach to secure data transfer and file management. The application embodies the philosophy of "Secure by Design, Simple by Nature" through its robust encryption architecture, intuitive user interface and enterprise grade security features.

The development process followed carefully planned security first iterations, progressing from encryption foundations to advanced features like time limited sharing and download restrictions. The result is a production ready application that showcases modern secure web development capabilities while solving real world data sharing challenges.

## Key Features & Functionality

### File Management Capabilities
Core functionality centers around secure file operations with comprehensive upload, download and lifecycle management features.

<div align="center">
  <img width="50%" alt="Image" src="https://github.com/user-attachments/assets/84055582-1137-4bb3-93a3-b8ffdb187064" />
</div>

**Comprehensive File Features:**
- **Secure File Upload** - Multi format file support with encryption and metadata tracking
- **Flexible Download Options** - Public link generation with embedded access controls
- **File Lifecycle Management** - Automated cleanup, expiration handling and storage optimization
- **Access Analytics** - Detailed download tracking, usage statistics and audit logs
- **File Integrity Verification** - Checksum validation and corruption detection
  
### Security First Design
The application provides enterprise grade security features with end to end encryption, secure key management and comprehensive access controls. All files undergo encryption before storage using industry standard cryptographic algorithms.

<div align="center">
  <img width="50%" alt="Image" src="https://github.com/user-attachments/assets/cc51c103-556a-4de3-9a63-31101f8d211d" />
</div>

**Advanced Security Features:**
- **End to End Encryption** - AES-256-GCM encryption with authenticated encryption
- **Secure Key Management** - Environment based encryption key storage
- **Time Limited Sharing** - Configurable automatic file expiration and secure cleanup
- **Download Restrictions** - Granular maximum download limits with usage tracking
- **One Time Downloads** - Self-destructing file links with automatic cleanup
- **Metadata Protection** - Encrypted file metadata storage with integrity verification

### Advanced Sharing Controls
**Intelligent Sharing System:**
- **Time-based Expiration** - Configurable file lifetime with granular time controls
- **Download Limiting** - Precise download count restrictions per file
- **One-time Access** - Single use download links with immediate cleanup
- **Public Link Generation** - Secure shareable URLs with embedded access tokens
- **Access Logging** - Comprehensive download tracking and security audit trails

## Technical Highlights

### Encryption & Security Implementation
The application implements multiple layers of security to ensure comprehensive data protection and regulatory compliance.
**Security Implementation Details:**
- **AES-256-GCM Encryption** - Industry standard symmetric encryption with authenticated encryption
- **Secure Random Key Generation** - Cryptographically secure key management
- **Input Validation & Sanitization** - Comprehensive prevention of injection attacks and XSS
- **CORS Policy Configuration** - Secure cross-origin request handling
- **Request Rate Limiting** - Advanced DDoS protection and abuse prevention

### Performance Optimization
**Performance Enhancement Features:**
- **Redis Caching Layer** - High performance metadata retrieval with intelligent cache invalidation
- **MinIO Object Storage** - Scalable and high-performance distributed file storage
- **Gin Framework Optimization** - High throughput HTTP server with efficient routing
- **Efficient File Streaming** - Memory optimized download performance for large files
- **Background Processing** - Non-blocking file operations with asynchronous processing

### Scalability & Infrastructure
**Infrastructure & Scalability Features:**
- **Docker Containerization** - Consistent deployment across environments
- **Docker Compose Orchestration** - Multi service application management
- **Microservices Architecture** - Independent service scaling and deployment flexibility
- **Cloud Ready Design** - Platform agnostic design suitable for major cloud providers
- **Health Check Endpoints** - Comprehensive system monitoring integration

## Technical Architecture

### Backend Implementation
The server side architecture leverages **Go 1.24** with **Gin Web Framework** to create a high performance RESTful API. File storage utilizes **MinIO** S3 compatible object storage for scalable file management, integrated with **Redis** for fast metadata caching. The encryption system implements industry standard **AES-256-GCM** for comprehensive file protection.

**Core Backend Technologies:**
- **Go 1.24** - Latest Go with modern language features and performance optimizations
- **Gin Web Framework 1.9.1** - High performance HTTP web framework with middleware support
- **MinIO Client 7.0.63** - Enterprise grade S3 compatible object storage integration
<div align="center">
  <img width="50%" alt="Image" src="https://github.com/user-attachments/assets/ea890ee1-15fe-47c1-acd8-c0ecbf8f4448" />
</div>

- **Redis Client 9.11.0** - High performance in memory caching and metadata storage
- **Clean Architecture** - Layered design pattern: handlers → services → repository → storage
- **Comprehensive Middleware** - CORS, logging, recovery and request validation

### Frontend Development
The client side application utilizes **React.js 18.2** with modern functional components and hooks API for efficient state management. **Axios** provides robust HTTP communication with comprehensive error handling. **React Dropzone** enables professional drag and drop file upload functionality.

**Frontend Technology Stack:**
- **React 18.2.0** - Latest React with concurrent features and modern component architecture
- **React Hooks** - Advanced state management with useState, useEffect and custom hooks
- **Axios 1.5.0** - Enterprise grade HTTP client with request/response interceptors
- **React Dropzone 14.2.3** - Professional drag and drop file upload with validation
- **Responsive CSS** - Modern CSS Grid and Flexbox with mobile first design
- **Comprehensive Error Handling** - User-friendly error messages and loading states

## API Documentation
<div align="center">
  <img width="60%" alt="Image" src="https://github.com/user-attachments/assets/b13ec928-7cf6-4843-8197-e161221ddfae" />
</div>

### Core Endpoints

#### Upload File
```http
POST /api/v1/upload
Content-Type: multipart/form-data

Form Parameters:
- file: File to upload (required)
- expiry_hours: File expiration time in hours (default: 24)
- max_downloads: Maximum download count (default: 0 = unlimited)
- is_one_time: Delete file after first download (default: false)

Response:
{
  "file_id": "550e8400-e29b-41d4-a716-446655440000",
  "share_link": "http://localhost:8080/api/v1/download/550e8400-e29b-41d4-a716-446655440000",
  "expires_at": "2024-01-03T12:00:00Z"
}
```

#### Download File
```http
GET /api/v1/download/:id
```

#### List Files
```http
GET /api/v1/files

Response:
{
  "files": [{
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "original_name": "document.pdf", 
    "size": 2048576,
    "content_type": "application/pdf",
    "upload_time": "2024-01-01T10:00:00Z",
    "expires_at": "2024-01-02T10:00:00Z",
    "download_count": 3,
    "max_downloads": 5
  }]
}
```

#### Health Check
```http
GET /health

Response:
{
  "status": "ok",
  "service": "filepin-backend"
}
```


## Project Structure
```
filepin/
├── backend/                      # Go backend application
│   ├── cmd/server/main.go        # Application entry point
│   ├── internal/                 # Internal application packages
│   │   ├── config/               # Configuration management
│   │   ├── crypto/               # AES-256-GCM encryption implementation
│   │   ├── handlers/             # HTTP request handlers
│   │   ├── middleware/           # CORS middleware configuration
│   │   ├── models/               # File data models and structures
│   │   ├── repository/           # Data access layer
│   │   ├── services/             # Business logic implementation
│   │   └── storage/              # MinIO storage integration
│   └── go.mod                    # Go module dependencies
├── frontend/                     # React frontend application
│   ├── src/
│   │   ├── components/           # React components
│   │   │   ├── FileUpload.jsx    # File upload interface
│   │   │   ├── FileDownload.jsx  # File download interface
│   │   │   └── FileList.jsx      # File management interface
│   │   ├── services/api.js       # API communication layer
│   │   └── App.jsx               # Main React component
│   └── package.json              # Frontend dependencies
├── docker-compose.yml            # Multi container deployment
├── Dockerfile.backend            # Backend container configuration
└── Dockerfile.frontend           # Frontend container configuration
```

## Installation & Setup

### Prerequisites
- Docker & Docker Compose
- Go 1.21+ (for development)
- Node.js 16+ (for development)

### Quick Start with Docker (Recommended)
```bash
# Clone the repository
git clone <repository-url>
cd filepin

# Start all services
docker-compose up -d

# Access the application
# Frontend: http://localhost:3000
# Backend API: http://localhost:8080  
# MinIO Console: http://localhost:9001
```

### Development Setup

#### Backend Development
```bash
cd backend
go mod download

# Set environment variables (required)
export SERVER_PORT=8080
export MINIO_ENDPOINT=localhost:9000
export MINIO_ACCESS_KEY=minioadmin
export MINIO_SECRET_KEY=minioadmin
export MINIO_BUCKET=filepin
export REDIS_ADDR=localhost:6379
export ENCRYPTION_KEY=your-32-char-encryption-key-here

go run cmd/server/main.go
```

#### Frontend Development
```bash
cd frontend
npm install
npm start
```

## Security Considerations

### Data Protection
- All files are encrypted using AES-256-GCM before storage with authenticated encryption
- Encryption keys are managed through environment variables
- File metadata is cached securely in Redis with TTL expiration
- Automatic cleanup removes expired files and associated metadata

### Access Control
- Time based file expiration prevents indefinite access to shared files
- Download limits provide granular control over file distribution
- One time download links provide maximum security for sensitive documents
- Public links contain embedded access tokens for secure sharing

## Deployment & Production

### Docker Production Deployment
```bash
# Production environment setup
export NODE_ENV=production
export GIN_MODE=release

# Deploy with production configuration
docker-compose up -d
```

### Security Hardening
- Enable HTTPS with SSL/TLS certificates
- Configure firewall rules for service access
- Set up reverse proxy with rate limiting
- Implement monitoring and alerting systems
- Regular security updates and dependency management

## Conclusion

FilePin represents a comprehensive demonstration of modern secure application development, showcasing advanced proficiency in Go and React technologies while implementing industry standard security practices. The project exemplifies the ability to deliver production quality applications that balance robust security with exceptional user experience.

The application demonstrates expertise in full stack development, security engineering and modern DevOps practices, suitable for enterprise environments requiring secure data handling and scalable architecture design.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

**Technology Stack:** Go • React.js • MinIO • Redis • Docker • AES-256-GCM Encryption



