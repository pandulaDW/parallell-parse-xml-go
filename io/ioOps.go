package io

import (
	"bufio"
	"bytes"
	"compress/gzip"
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
		return createBufferedZipFileReader(model.ZipFileName, model.FileSize)
	case models.DownloadZipRead:
		return createBufferedDownloadReader(model.Url, model.FileSize)
	case models.XMLDownloadAndRead:
		return createDownloadWriteAndReader(model.Url, model.XmlFileName, model.FileSize)
	case models.XMLWriteAndRead:
		return createReadWriteAndReader(model.ZipFileName, model.XmlFileName, model.FileSize)
	default:
		return nil, nil
	}
}

// CreateBufferedWriter creates a buffered writer based on the output processing stage
func CreateBufferedWriter(model models.GliefModel, outStage models.OutputStage) (*bufio.Writer, error) {
	switch outStage {
	case models.CSVFileWrite:
		return createFileWriter(model.CsvFileName)
	case models.ZipFileWrite:
		return createGzipWriter(model)
	default:
		return nil, nil
	}
}

// Readers -----------------------------------------------------
func createBufferedFileReader(filename string) (*bufio.Reader, error) {
	file, err := os.Open(filename)
	return bufio.NewReader(file), err
}

func createBufferedZipFileReader(zipFile string, fileSize int) (*bufio.Reader, error) {
	zippedContent, err := ioutil.ReadFile(zipFile)
	if err != nil {
		return nil, err
	}
	content, err := unzipFilesInMemory(zippedContent, fileSize)
	return bufio.NewReader(bytes.NewReader(content)), err
}

func createBufferedDownloadReader(url string, fileSize int) (*bufio.Reader, error) {
	zippedContent, err := downloadFileToMemory(url)
	if err != nil {
		return nil, err
	}
	content, err := unzipFilesInMemory(zippedContent, fileSize)
	return bufio.NewReader(bytes.NewReader(content)), err
}

func createDownloadWriteAndReader(url, filename string, fileSize int) (*bufio.Reader, error) {
	zippedContent, err := downloadFileToMemory(url)
	if err != nil {
		return nil, err
	}
	content, err := unzipFilesInMemory(zippedContent, fileSize)
	if err != nil {
		return nil, err
	}
	err = unzipFilesToDisk(content, filename)
	if err != nil {
		return nil, err
	}
	return createBufferedFileReader(filename)
}

func createReadWriteAndReader(zipFile, filename string, fileSize int) (*bufio.Reader, error) {
	zippedContent, err := ioutil.ReadFile(zipFile)
	if err != nil {
		return nil, err
	}
	content, err := unzipFilesInMemory(zippedContent, fileSize)
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

// Writers ------------------------------------------------------
func createFileWriter(filename string) (*bufio.Writer, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return bufio.NewWriter(file), err
}

func createGzipWriter(model models.GliefModel) (*bufio.Writer, error) {
	zipFile, err := os.OpenFile(model.GZipFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	zipWriter := gzip.NewWriter(zipFile)
	zipWriter.Name = model.CsvFileName
	bufferedWriter := bufio.NewWriter(zipWriter)
	return bufferedWriter, nil
}
