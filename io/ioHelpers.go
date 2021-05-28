package io

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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

func unzipFilesInMemory(zipContent []byte, fileSize int) ([]byte, error) {
	zipReader, err := zip.NewReader(bytes.NewReader(zipContent), int64(len(zipContent)))
	if err != nil {
		return nil, err
	}

	zipFile := zipReader.File[0]
	fmt.Println("Reading file: ", zipFile.Name)

	f, err := zipFile.Open()
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, fileSize))
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
