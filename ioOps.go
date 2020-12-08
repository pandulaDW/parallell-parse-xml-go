package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"os"
)

// creates the buffered reader based on the input processing stage
func createBufferedReader(model GliefModel, inStage InputStage) (*bufio.Reader, error) {
	switch inStage {
	case XMLFileRead:
		return createBufferedFileReader(model.xmlFileName)
	case ZipFileRead:
		return createBufferedZipFileReader(model.archiveFileName)
	case DownloadZipRead:
		return createBufferedDownloadReader(model.url)
	default:
		return nil, nil
	}
}

// creates the buffered writer based on the output processing stage
func createBufferedWriter(model GliefModel, outStage OutputStage) (*bufio.Writer, error) {
	switch outStage {
	case CSVFileWrite:
		return createFileWriter(model.csvFileName)
	case ZipFileWrite:
		return createGzipWriter(model)
	default:
		return nil, nil
	}
}

// Readers -----------------------------------------------------
func createFileReader(filename string) ([]byte, error) {
	content, err := ioutil.ReadFile(filename)
	return content, err
}

func createBufferedFileReader(filename string) (*bufio.Reader, error) {
	file, err := os.Open(filename)
	return bufio.NewReader(file), err
}

func createZipFileReader(zipFile string) ([]byte, error) {
	zippedContent, err := ioutil.ReadFile(zipFile)
	if err != nil {
		return nil, err
	}
	return unzipFilesInMemory(zippedContent)
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

// Writers ------------------------------------------------------
func createFileWriter(filename string) (*bufio.Writer, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return bufio.NewWriter(file), err
}

func createGzipWriter(model GliefModel) (*bufio.Writer, error) {
	zipFile, err := os.OpenFile(model.GZipFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	zipWriter := gzip.NewWriter(zipFile)
	zipWriter.Name = model.csvFileName
	bufferedWriter := bufio.NewWriter(zipWriter)
	return bufferedWriter, nil
}
