package io

import (
	"bufio"
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pandulaDW/parallell-parse-xml-go/models"
	"io/ioutil"
	"os"
	"runtime"
)

// CreateBufferedReader creates a buffered reader based on the input processing stage
func CreateBufferedReader(model models.GliefModel, inStage models.InputStage) (*bufio.Reader, error) {
	switch inStage {
	case models.XMLFileRead:
		return createBufferedFileReader(model.XmlFileName)
	case models.ZipFileRead:
		return createBufferedZipFileReader(model.ZipFileName)
	case models.DownloadZipRead:
		return createBufferedDownloadReader(model.Url)
	case models.XMLDownloadAndRead:
		return createDownloadWriteAndReader(model.Url, model.XmlFileName)
	case models.XMLWriteAndRead:
		return createReadWriteAndReader(model.ZipFileName, model.XmlFileName)
	case models.S3XMLFileRead:
		return createS3fileReader(model)
	default:
		return nil, nil
	}
}

// Readers -----------------------------------------------------
func createBufferedFileReader(filename string) (*bufio.Reader, error) {
	file, err := os.Open(filename)
	return bufio.NewReader(file), err
}

func createBufferedZipFileReader(zipFile string) (*bufio.Reader, error) {
	zippedContent, err := ioutil.ReadFile(zipFile)
	if err != nil {
		return nil, err
	}
	content, err := unzipFilesInMemory(zippedContent)
	return bufio.NewReader(bytes.NewReader(content)), err
}

func createBufferedDownloadReader(url string) (*bufio.Reader, error) {
	zippedContent, err := downloadFileToMemory(url)
	if err != nil {
		return nil, err
	}
	content, err := unzipFilesInMemory(zippedContent)
	return bufio.NewReader(bytes.NewReader(content)), err
}

func createDownloadWriteAndReader(url, filename string) (*bufio.Reader, error) {
	zippedContent, err := downloadFileToMemory(url)
	if err != nil {
		return nil, err
	}
	content, err := unzipFilesInMemory(zippedContent)
	if err != nil {
		return nil, err
	}
	err = unzipFilesToDisk(content, filename)
	if err != nil {
		return nil, err
	}
	return createBufferedFileReader(filename)
}

func createReadWriteAndReader(zipFile, filename string) (*bufio.Reader, error) {
	zippedContent, err := ioutil.ReadFile(zipFile)
	if err != nil {
		return nil, err
	}
	content, err := unzipFilesInMemory(zippedContent)
	if err != nil {
		return nil, err
	}
	err = unzipFilesToDisk(content, filename)
	if err != nil {
		return nil, err
	}
	content = nil
	runtime.GC()
	return createBufferedFileReader(filename)
}

func createS3fileReader(model models.GliefModel) (*bufio.Reader, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(model.Bucket),
		Key:    aws.String(model.XmlFileName),
	}

	request, err := model.SVC.GetObject(input)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(request.Body)
	return reader, nil
}
