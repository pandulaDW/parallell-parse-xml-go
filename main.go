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
	filenames := io.GetZipFileNames("data")
	processingInServer(filenames)
	fmt.Println(strings.Repeat("-", 50))
}

func processingInServer(filenames map[string]string) {
	rrModel := models.CreateRelationshipModel()
	rrModel.ZipFileName = "./data/" + filenames["rr"]

	leiModel := models.CreateLEIModel()
	leiModel.ZipFileName = "./data/" + filenames["lei"]

	repexModel := models.CreateReportingExceptionModel()
	repexModel.ZipFileName = "./zip_files/" + filenames["repex"]

	processing.ConcurrentProcessing(*rrModel, models.ZipFileRead, models.CSVFileWrite)
	fmt.Println("Finished processing relationship file")

	//processing.ConcurrentProcessing(*leiModel, models.ZipFileRead, models.CSVFileWrite)
	//fmt.Println("Finished processing lei file")

	//processing.ConcurrentProcessing(*repexModel, models.ZipFileRead, models.CSVFileWrite)
	//fmt.Println("Finished processing repex file")
}
