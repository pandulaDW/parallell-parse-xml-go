package models

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
	"time"
)

// common variables
var bucket string
var sess *session.Session
var svc *s3.S3

func init() {
	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_DEFAULT_REGION")),
	}))

	svc = s3.New(sess)
	bucket = "lambda-test-go-upload"
}

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
	Bucket            string
	SVC               *s3.S3
	Session           *session.Session
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
	// S3XMLFileRead reads xml file from S3 line by line
	S3XMLFileRead InputStage = "S3XMLFileRead"
)

func createUrl(prefix string) string {
	day := time.Now().AddDate(0, 0, -1).Format("20060102")
	url := fmt.Sprintf("https://leidata.gleif.org/api/v1/concatenated-files/%s/%s/zip", prefix, day)
	return url
}

func addDefaultParams(model *GliefModel) {
	model.Bucket = bucket
	model.SVC = svc
	model.Session = sess
}

// CreateRelationshipModel returns the model definition for Relationship file type with sensible defaults
func CreateRelationshipModel() *GliefModel {
	rrModel := GliefModel{Prefix: "rr", Category: "Relationship"}
	addDefaultParams(&rrModel)
	rrModel.RecordsPerRoutine = 1000
	rrModel.CsvFileName = "csv/rrFile.csv"
	rrModel.GZipFileName = "/tmp/rrFile.zip"
	rrModel.Url = createUrl("rr")
	rrModel.XmlFileName = "xml/rr.xml"
	return &rrModel
}

// CreateLEIModel returns the model definition for LEI file type with sensible defaults
func CreateLEIModel() *GliefModel {
	leiModel := GliefModel{Prefix: "lei", Category: "LEI"}
	addDefaultParams(&leiModel)
	leiModel.RecordsPerRoutine = 2000
	leiModel.CsvFileName = "csv/leiFile.csv"
	leiModel.GZipFileName = "/tmp/leiFile.zip"
	leiModel.XmlFileName = "xml/lei2.xml"
	leiModel.Url = createUrl("lei2") // url is defined as this
	leiModel.Bucket = bucket
	leiModel.SVC = svc
	return &leiModel
}

// CreateReportingExceptionModel returns the model definition for Reporting exception file type with sensible defaults
func CreateReportingExceptionModel() *GliefModel {
	repexModel := GliefModel{Prefix: "repex", Category: "Exception"}
	addDefaultParams(&repexModel)
	repexModel.RecordsPerRoutine = 1000
	repexModel.CsvFileName = "csv/repexFile.csv"
	repexModel.GZipFileName = "/tmp/repexFile.zip"
	repexModel.Url = createUrl("repex")
	repexModel.XmlFileName = "xml/repex.xml"
	repexModel.Bucket = bucket
	repexModel.SVC = svc
	return &repexModel
}
