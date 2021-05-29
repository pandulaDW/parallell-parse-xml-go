package io

import (
	"compress/gzip"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pandulaDW/parallell-parse-xml-go/models"
	"os"
	"path/filepath"
)

func CreateZipFileWriter(model models.GliefModel) (*gzip.Writer, error) {
	zipFile, err := os.OpenFile(model.GZipFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	zipWriter := gzip.NewWriter(zipFile)
	zipWriter.Name = filepath.Base(model.CsvFileName)
	return zipWriter, nil
}

func WriteFileToS3(model models.GliefModel) error {
	uploader := s3manager.NewUploader(model.Session)
	f, err := os.Open(model.GZipFileName)
	if err != nil {
		return err
	}

	// Upload the file to S3.
	fmt.Println("Uploading file to S3...")
	key := "zip/" + filepath.Base(model.GZipFileName)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(model.Bucket),
		Key:    aws.String(key),
		Body:   f,
	})

	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
	fmt.Printf("File uploaded to, %s\n", result.Location)

	return nil
}
