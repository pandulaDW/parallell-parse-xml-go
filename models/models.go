package models

import (
	"fmt"
	"time"
)

// GliefModel includes information needed for Glief file processing
// Based on the processing stage, (testFile, testRar, testAWS) different
// properties will be used
type GliefModel struct {
	Prefix            string
	Category          string
	RecordsPerRoutine int
	XmlFileName       string
	ZipFileName       string
	GZipFileName      string
	CsvFileName       string
	Url               string
}

// InputStage represents the input method of the processing pipeline
type InputStage string

const (
	// XMLFileRead reads the xml file directly from the disk
	XMLFileRead InputStage = "XMLFileRead"
	// XMLDownloadAndRead downloads, writes the zip content to a file and then stream read from that file
	XMLDownloadAndRead InputStage = "XMLDownloadAndRead"
	// XMLWriteAndRead reads, writes the zip content to a file and then stream read from that file
	XMLWriteAndRead InputStage = "XMLWriteAndRead"
	// ZipFileRead reads the zip file directly from the disk
	ZipFileRead InputStage = "ZipFileRead"
	// DownloadZipRead downloads the zip file and process it in memory
	DownloadZipRead InputStage = "DownloadZipRead"
)

func createUrl(prefix string) string {
	day := time.Now().AddDate(0, 0, -1).Format("20060102")
	url := fmt.Sprintf("https://leidata.gleif.org/api/v1/concatenated-files/%s/%s/zip", prefix, day)
	return url
}

// CreateRelationshipModel returns the model definition for Relationship file type with sensible defaults
func CreateRelationshipModel() *GliefModel {
	rrModel := GliefModel{Prefix: "rr", Category: "Relationship"}
	rrModel.RecordsPerRoutine = 1000
	rrModel.CsvFileName = "csv/rrFile.csv"
	rrModel.GZipFileName = "/tmp/rrFile.zip"
	rrModel.Url = createUrl("rr")
	return &rrModel
}

// CreateLEIModel returns the model definition for LEI file type with sensible defaults
func CreateLEIModel() *GliefModel {
	leiModel := GliefModel{Prefix: "lei", Category: "LEI"}
	leiModel.RecordsPerRoutine = 2000
	leiModel.CsvFileName = "csv/leiFile.csv"
	leiModel.GZipFileName = "/tmp/leiFile.zip"
	leiModel.XmlFileName = "leiXML.xml"
	leiModel.Url = createUrl("lei2") // url is defined as this
	return &leiModel
}

// CreateReportingExceptionModel returns the model definition for Reporting exception file type with sensible defaults
func CreateReportingExceptionModel() *GliefModel {
	repexModel := GliefModel{Prefix: "repex", Category: "Exception"}
	repexModel.RecordsPerRoutine = 1000
	repexModel.CsvFileName = "csv/repexFile.csv"
	repexModel.GZipFileName = "/tmp/repexFile.zip"
	repexModel.Url = createUrl("repex")
	return &repexModel
}
