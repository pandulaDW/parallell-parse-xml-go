package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pandulaDW/parallell-parse-xml-go/processing"
)

func HandleRequest() (string, error) {
	fmt.Println("Processing started...")
	err := processFiles()
	if err != nil {
		return "Unsuccessful", err
	}
	return "Successful", nil
}

func processFiles() error {
	err := processing.ProcessFile("rr")
	if err != nil {
		return err
	}
	fmt.Println("Finished processing relationship file")

	err = processing.ProcessFile("lei")
	if err != nil {
		return err
	}
	fmt.Println("Finished processing lei file")

	err = processing.ProcessFile("repex")
	if err != nil {
		return err
	}
	fmt.Println("Finished processing relationship file")

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
