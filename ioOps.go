package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"os"
)

// creates the buffered reader based on the processing stage
func createBufferedReader(model GliefModel, method string) (*bufio.Reader, error) {
	switch method {
	case "XMLFileRead":
		return createBufferedFileReader(model.xmlFileName)
	case "ZipFileRead":
		return createBufferedZipFileReader(model.xmlFileName)
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

// Writers ------------------------------------------------------
func createGzipWriter(model GliefModel) *bufio.Writer {
	zipFile, _ := os.OpenFile(model.GZipFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	zipWriter := gzip.NewWriter(zipFile)
	zipWriter.Name = model.csvFileName
	bufferedWriter := bufio.NewWriter(zipWriter)
	return bufferedWriter
}
