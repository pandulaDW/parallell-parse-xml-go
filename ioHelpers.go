package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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

func unzipFilesToDisk(unzippedContent []byte, filename string) error {
	err := ioutil.WriteFile(filename, unzippedContent, 0666)
	return err
}

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
