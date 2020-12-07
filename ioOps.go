package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func createBufferedReader(filename string) (*bufio.Reader, error) {
	zipContent, err := ioutil.ReadFile(filename)
	// zipContent, err := downloadFileToMemory("https://leidata.gleif.org/api/v1/concatenated-files/rr/20201206/zip")

	if err != nil {
		return nil, err
	}

	content, err := unzipFilesInMemory(zipContent)
	if err != nil {
		return nil, err
	}

	bytesReader := bytes.NewReader(content)
	bufferedReader := bufio.NewReader(bytesReader)
	return bufferedReader, nil
}

func createGzipWriter(zipfilePath, filename string) (*gzip.Writer, *os.File) {
	zipFile, _ := os.OpenFile(zipfilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	zipWriter := gzip.NewWriter(zipFile)
	zipWriter.Name = filename
	return zipWriter, zipFile
}

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

	// read all the files from zip archive
	allContent := make([]byte, 0)
	for _, zipFile := range zipReader.File {
		fmt.Println("Reading file: ", zipFile.Name)
		unzippedFileBytes, err := readZipFile(zipFile)
		if err != nil {
			log.Println(err)
			continue
		}
		allContent = append(allContent, unzippedFileBytes...)
	}

	return allContent, nil
}

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
