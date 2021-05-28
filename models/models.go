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
	FileSize          int
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

// OutputStage represents the output method of the processing pipeline
type OutputStage string

const (
	// CSVFileWrite write the csv file to disk
	CSVFileWrite OutputStage = "CSVFileWrite"
	// ZipFileWrite writes the zip file to disk
	ZipFileWrite OutputStage = "ZipFileWrite"
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
	rrModel.CsvFileName = "data/rrFile.csv"
	rrModel.Url = createUrl("rr")
	rrModel.FileSize = 2 * 1024 * 1024 * 1024
	return &rrModel
}

// CreateLEIModel returns the model definition for LEI file type with sensible defaults
func CreateLEIModel() *GliefModel {
	leiModel := GliefModel{Prefix: "lei", Category: "LEI"}
	leiModel.RecordsPerRoutine = 2000
	leiModel.CsvFileName = "data/leiFile.csv"
	leiModel.XmlFileName = "leiXML.xml"
	leiModel.Url = createUrl("lei2") // url is defined as this
	leiModel.FileSize = 4.5 * 1024 * 1024 * 1024
	return &leiModel
}

// CreateReportingExceptionModel returns the model definition for Reporting exception file type with sensible defaults
func CreateReportingExceptionModel() *GliefModel {
	repexModel := GliefModel{Prefix: "repex", Category: "Exception"}
	repexModel.RecordsPerRoutine = 1000
	repexModel.CsvFileName = "data/repexFile.csv"
	repexModel.Url = createUrl("repex")
	repexModel.FileSize = 2 * 1024 * 1024 * 1024
	return &repexModel
}
