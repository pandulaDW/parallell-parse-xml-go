package io

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"runtime"
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

	// run the GC to clear zip file content
	runtime.GC()

	return buf.Bytes(), nil
}

func unzipFilesToDisk(unzippedContent []byte, filename string) error {
	err := ioutil.WriteFile(filename, unzippedContent, 0666)
	return err
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
