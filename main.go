package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pandulaDW/parallell-parse-xml-go/io"
	"github.com/pandulaDW/parallell-parse-xml-go/models"
	"github.com/pandulaDW/parallell-parse-xml-go/processing"
	"os"
)

func HandleRequest() (string, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_DEFAULT_REGION")),
	}))

	fmt.Println("Processing started...")
	processingInServer(sess)

	return "Successful", nil
}

func main() {
	lambda.Start(HandleRequest)
}

func processingInServer(sess *session.Session) {
	leiModel := models.CreateLEIModel()
	leiModel.XmlFileName = "data/20201202-gleif-concatenated-file-lei2.xml"

	processing.ConcurrentProcessing(*leiModel, models.XMLFileRead)
	fmt.Println("Finished processing relationship file")

	fmt.Println("Finished writing the zip file")

	io.PrintMemUsage()
}
