package minio

import (
	"bytes"
	"context"
	"mime/multipart"
	"path/filepath"

	"github.com/aditya3232/gatewatchApp-services.git/connection"
	"github.com/aditya3232/gatewatchApp-services.git/helper"
	"github.com/aditya3232/gatewatchApp-services.git/log"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

var MinioClient *minio.Client

// upload file to minio server using PutObject
func UploadFileFromPutObject(bucketName string, objectName string, imageBytes []byte) (minio.UploadInfo, error) {
	imageBytes, err := helper.CompressImageBytes(imageBytes)
	if err != nil {
		return minio.UploadInfo{}, err
	}

	n, err := connection.MinioGatewatch().PutObject(context.TODO(), bucketName, objectName, bytes.NewReader(imageBytes), int64(len(imageBytes)), minio.PutObjectOptions{
		ContentType: "image/jpeg/jpg", // Sesuaikan tipe konten sesuai dengan gambar yang diunduh
	})

	if err != nil {
		return minio.UploadInfo{}, err
	}

	return n, nil
}

func UploadFile(bucketName string, objectName string, filePath string) (minio.UploadInfo, error) {
	n, err := MinioClient.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		log.Error(err)
		return minio.UploadInfo{}, err
	}

	return n, nil
}

// function to upload file from multipart.FileHeader
func UploadFileFromHeader(bucketName string, objectName string, fileHeader *multipart.FileHeader) (minio.UploadInfo, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return minio.UploadInfo{}, err
	}
	defer file.Close()

	// upload file to minio server using PutObject
	n, err := MinioClient.PutObject(context.Background(), bucketName, objectName, file, fileHeader.Size, minio.PutObjectOptions{ContentType: fileHeader.Header.Get("Content-Type")})

	if err != nil {
		return minio.UploadInfo{}, err
	}

	return n, nil
}

// function to upload base64 file to minio server using gin context
func UploadBase64(bucketName string, objectName string, base64 string) (minio.UploadInfo, error) {
	image, _ := helper.Base64ToImage(base64)

	// get mime type from image path
	mimeType := helper.GetMimeType(image)

	// check if object name is empty, if empty then use name from image
	if objectName == "" {
		objectName = filepath.Base(image)
	}

	n, err := MinioClient.FPutObject(context.Background(), bucketName, objectName, image, minio.PutObjectOptions{ContentType: mimeType})

	if err != nil {
		return minio.UploadInfo{}, err
	}

	// remove image from temp folder
	helper.RemoveFile(image)

	return n, nil
}

// function to upload file to minio server using gin context
func UploadFileWithContext(c *gin.Context, bucketName string, objectName string) (minio.UploadInfo, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return minio.UploadInfo{}, err
	}

	buffer, err := file.Open()
	if err != nil {
		return minio.UploadInfo{}, err
	}
	defer buffer.Close()

	// upload file to minio server using PutObject
	n, err := MinioClient.PutObject(context.Background(), bucketName, objectName, buffer, file.Size, minio.PutObjectOptions{ContentType: file.Header.Get("Content-Type")})

	if err != nil {
		return minio.UploadInfo{}, err
	}

	return n, nil
}
