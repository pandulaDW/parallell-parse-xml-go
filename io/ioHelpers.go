package io

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func downloadFileToMemory(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buffer := new(bytes.Buffer)
	_, err = io.Copy(buffer, resp.Body)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(buffer)
}

func unzipFilesInMemory(zipContent []byte) ([]byte, error) {
	zipReader, err := zip.NewReader(bytes.NewReader(zipContent), int64(len(zipContent)))
	if err != nil {
		return nil, err
	}

	zipFile := zipReader.File[0]
	fmt.Println("Reading file:", zipFile.Name)

	f, err := zipFile.Open()
	if err != nil {
		return nil, err
	}

	bufferSize := zipFile.FileInfo().Size()
	buf := bytes.NewBuffer(make([]byte, 0, bufferSize))
	_, err = io.Copy(buf, f)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func unzipFilesToDisk(unzippedContent []byte, filename string) error {
	err := ioutil.WriteFile(filename, unzippedContent, 0666)
	return err
}

//GetZipFileNames will return the files names dictionary
func GetZipFileNames(path string) map[string]string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	filenames := make(map[string]string)
	for _, file := range files {
		if strings.Contains(file.Name(), "rr.xml") {
			filenames["rr"] = file.Name()
		}
		if strings.Contains(file.Name(), "repex.xml") {
			filenames["repex"] = file.Name()
		}
		if strings.Contains(file.Name(), "lei2.xml") {
			filenames["lei"] = file.Name()
		}
	}

	return filenames
}

// PrintMemUsage prints allocation and GC information
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func WriteFileToS3(sess *session.Session, fileName string) error {
	bucketName := "lambda-test-go-upload"
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// Upload the file to S3.
	fmt.Println("Uploading file to S3...")
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filepath.Base(fileName)),
		Body:   f,
	})

	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
	fmt.Printf("File uploaded to, %s\n", result.Location)

	return nil
}
