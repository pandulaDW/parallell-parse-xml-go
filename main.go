package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Processing started...")
	processingInServer()
	fmt.Println(strings.Repeat("-", 50))
}

func processingInServer() {
	rrModel := createRelationshipModel()
	leiModel := createLEIModel()
	repexModel := createReportingExceptionModel()

	concurrentProcessing(*rrModel, XMLDownloadAndRead, CSVFileWrite)
	fmt.Println("Finished processing relationship file")

	concurrentProcessing(*leiModel, XMLDownloadAndRead, CSVFileWrite)
	fmt.Println("Finished processing lei file")

	concurrentProcessing(*repexModel, XMLDownloadAndRead, CSVFileWrite)
	fmt.Println("Finished processing repex file")
}
