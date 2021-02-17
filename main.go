package main

import (
	"fmt"
	"github.com/pandulaDW/parallell-parse-xml-go/models"
	"github.com/pandulaDW/parallell-parse-xml-go/processing"
	"strings"
)

func main() {
	fmt.Println("Processing started...")
	processingInServer()
	fmt.Println(strings.Repeat("-", 50))
}

func processingInServer() {
	//rrModel := models.CreateRelationshipModel()
	leiModel := models.CreateLEIModel()
	leiModel.ZipFileName = "data/20201202-gleif-concatenated-file-lei2.xml.5fc7579cab4ee.zip"
	//repexModel := models.CreateReportingExceptionModel()

	//processing.ConcurrentProcessing(*rrModel, models.XMLDownloadAndRead, models.CSVFileWrite)
	//fmt.Println("Finished processing relationship file")

	processing.ConcurrentProcessing(*leiModel, models.ZipFileRead, models.CSVFileWrite)
	fmt.Println("Finished processing lei file")

	//processing.ConcurrentProcessing(*repexModel, models.XMLDownloadAndRead, models.CSVFileWrite)
	//fmt.Println("Finished processing repex file")
}
