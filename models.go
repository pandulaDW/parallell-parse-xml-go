package main

import (
	"fmt"
	"time"
)

// GliefModel includes information needed for Glief file processing
// Based on the processing stage, (testFile, testRar, testAWS) different
// properties will be used
type GliefModel struct {
	prefix            string
	category          string
	recordsPerRoutine int
	xmlFileName       string
	zipFileName       string
	GZipFileName      string
	csvFileName       string
	url               string
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

// returns the model definition for Relationship file type with sensible defaults
func createRelationshipModel() *GliefModel {
	rrModel := GliefModel{prefix: "rr", category: "Relationship"}
	rrModel.recordsPerRoutine = 1000
	rrModel.csvFileName = "data/rrFile.csv"
	rrModel.url = createUrl("rr")
	return &rrModel
}

// returns the model definition for LEI file type with sensible defaults
func createLEIModel() *GliefModel {
	leiModel := GliefModel{prefix: "lei", category: "LEI"}
	leiModel.recordsPerRoutine = 2000
	leiModel.csvFileName = "data/leiFile.csv"
	leiModel.xmlFileName = "leiXML.xml"
	leiModel.url = createUrl("lei2") // url is defined as this
	return &leiModel
}

// returns the model definition for Reporting exception file type with sensible defaults
func createReportingExceptionModel() *GliefModel {
	repexModel := GliefModel{prefix: "repex", category: "Exception"}
	repexModel.recordsPerRoutine = 1000
	repexModel.csvFileName = "data/repexFile.csv"
	repexModel.url = createUrl("repex")
	return &repexModel
}
