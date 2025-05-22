package storage

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Storage interface {
	UploadFile(file *multipart.FileHeader, path string) (string, error)
	UploadFileRename(file *multipart.FileHeader, path string, name *string) (string, error)
}

type s3StorageImpl struct {
	client     *s3.S3
	bucketName string
	endpoint   string
}

func NewS3Storage(bucketName, region, endpoint, accessKey, secretKey string) S3Storage {
	if accessKey == "" || secretKey == "" || region == "" || bucketName == "" {
		panic("AWS credentials not found. Please set AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION, and AWS_BUCKET_NAME")
	}

	session := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			accessKey,
			secretKey,
			"",
		),
		Endpoint: &endpoint,
	}))

	client := s3.New(session)
	return &s3StorageImpl{
		client:     client,
		bucketName: bucketName,
		endpoint:   endpoint,
	}
}

// UploadFile implements S3Storage.
func (s *s3StorageImpl) UploadFile(file *multipart.FileHeader, path string) (string, error) {
	if file == nil {
		return "", fmt.Errorf("file is nil")
	}

	// Buka file
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer src.Close()

	// Tentukan path file
	var filePath string

	filePath = fmt.Sprintf("/%s/%s", path, file.Filename)

	// Clean path (hapus double slash)
	filePath = strings.Replace(filePath, "//", "/", -1)

	// Upload ke S3
	_, err = s.client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(filePath),
		Body:   src,
		ACL:    aws.String("public-read"),
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload to S3: %v", err)
	}

	url := fmt.Sprintf("%s/%s%s", s.endpoint, s.bucketName, filePath)

	return url, nil
}

// UploadFileRename implements S3Storage.
func (s *s3StorageImpl) UploadFileRename(file *multipart.FileHeader, path string, name *string) (string, error) {
	if file == nil {
		return "", fmt.Errorf("file is nil")
	}

	// Buka file
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer src.Close()

	// Tentukan path file
	var filePath string

	// Generate nama file baru
	var fileName string
	if name != nil && *name != "" {
		fileName = *name
	} else {
		fileName = strconv.FormatInt(time.Now().Unix(), 10)
	}

	// Dapatkan ekstensi file
	fileExt := filepath.Ext(file.Filename)

	filePath = fmt.Sprintf("/%s/%s%s", path, fileName, fileExt)

	// Clean path (hapus double slash)
	filePath = strings.Replace(filePath, "//", "/", -1)

	// Upload ke S3
	_, err = s.client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(filePath),
		Body:   src,
		ACL:    aws.String("public-read"),
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload to S3: %v", err)
	}

	url := fmt.Sprintf("%s/%s%s", s.endpoint, s.bucketName, filePath)

	return url, nil
}

// // UploadFile implements S3Storage.
// func (s *s3StorageImpl) UploadFile(file multipart.File, path string) (string, error) {
// 	uploader := s3manager.NewUploader(s.session)
// 	key := now
// 	contentType := fileHeader.Header.Get("Content-Type")

// 	_, err := uploader.Upload(&s3manager.UploadInput{
// 		Bucket:      aws.String(bucketName),
// 		Key:         aws.String(key),
// 		Body:        file,
// 		ContentType: &contentType,
// 	})

// 	if err != nil {
// 		return "", err
// 	}

// 	url := fmt.Sprintf("%s/%s/%s", s.endpoint, bucketName, key)
// 	return url, nil
// }
