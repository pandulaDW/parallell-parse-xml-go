package main

// GliefModel includes information needed for Glief file processing
// Based on the processing stage, (testFile, testRar, testAWS) different
// properties will be used
type GliefModel struct {
	prefix            string
	category          string
	recordsPerRoutine int
	xmlFileName       string
	archiveFileName   string
	GZipFileName      string
	csvFileName       string
	url               string
}

// GliefFileOps includes operations that should be performed for Glief files
type GliefFileOps interface {
	create() GliefModel
	unmarshal(content *string) *string
}

// InputStage represents the input method of the processing pipeline
type InputStage string

const (
	// XMLFileRead reads the xml file directly from the disk
	XMLFileRead InputStage = "XMLFileRead"
	// ZipFileRead reads the zip file directly from the disk
	ZipFileRead InputStage = "ZipFileRead"
	// DownloadZipRead downloads the zip file and process it
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
