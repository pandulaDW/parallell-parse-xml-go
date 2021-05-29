package io

import (
	"compress/gzip"
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
