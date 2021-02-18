package main

import (
	"fmt"
	"github.com/pandulaDW/parallell-parse-xml-go/io"
	"github.com/pandulaDW/parallell-parse-xml-go/models"
	"github.com/pandulaDW/parallell-parse-xml-go/processing"
	"strings"
)

func main() {
	fmt.Println("Processing started...")
	filenames := io.GetZipFileNames("zip_files")
	processingInServer(filenames)
	fmt.Println(strings.Repeat("-", 50))
}

func processingInServer(filenames map[string]string) {
	rrModel := models.CreateRelationshipModel()
	rrModel.ZipFileName = filenames["rr"]

	leiModel := models.CreateLEIModel()
	leiModel.ZipFileName = filenames["lei"]

	repexModel := models.CreateReportingExceptionModel()
	repexModel.ZipFileName = filenames["repex"]

	processing.ConcurrentProcessing(*rrModel, models.XMLDownloadAndRead, models.CSVFileWrite)
	fmt.Println("Finished processing relationship file")

	processing.ConcurrentProcessing(*leiModel, models.ZipFileRead, models.CSVFileWrite)
	fmt.Println("Finished processing lei file")

	processing.ConcurrentProcessing(*repexModel, models.XMLDownloadAndRead, models.CSVFileWrite)
	fmt.Println("Finished processing repex file")
}
