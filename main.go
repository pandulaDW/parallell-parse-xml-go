package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pandulaDW/parallell-parse-xml-go/models"
	"github.com/pandulaDW/parallell-parse-xml-go/processing"
	"os"
	"runtime"
)

func HandleRequest() (string, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_DEFAULT_REGION")),
	}))
	_ = sess

	fmt.Println("Processing started...")
	err := processingInServer()
	if err != nil {
		return "Unsuccessful", err
	}

	return "Successful", nil
}

func processingInServer() error {
	rrModel := models.CreateRelationshipModel()
	err := processing.ConcurrentProcessing(*rrModel, models.XMLFileRead)
	if err != nil {
		return err
	}
	runtime.GC()
	fmt.Println("Finished processing relationship file")

	leiModel := models.CreateLEIModel()
	err = processing.ConcurrentProcessing(*leiModel, models.XMLFileRead)
	if err != nil {
		return err
	}
	runtime.GC()
	fmt.Println("Finished processing lei file")

	repexModel := models.CreateReportingExceptionModel()
	err = processing.ConcurrentProcessing(*repexModel, models.XMLFileRead)
	if err != nil {
		return err
	}
	runtime.GC()
	fmt.Println("Finished processing relationship file")

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
