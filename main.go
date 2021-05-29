package main

import (
	"compress/gzip"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pandulaDW/parallell-parse-xml-go/io"
	"github.com/pandulaDW/parallell-parse-xml-go/models"
	"github.com/pandulaDW/parallell-parse-xml-go/processing"
	"log"
	"os"
	"path/filepath"
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
	zipFile, err := os.OpenFile(leiModel.GZipFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	zipWriter := gzip.NewWriter(zipFile)
	zipWriter.Name = filepath.Base(leiModel.CsvFileName)

	content := processing.ConcurrentProcessing(*leiModel, models.DownloadZipRead)
	fmt.Println("Finished processing relationship file")

	_, err = zipWriter.Write(content)
	if err != nil {
		log.Fatal(err)
	}
	err = zipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Finished writing the zip file")

	err = io.WriteFileToS3(sess, leiModel.GZipFileName)
	if err != nil {
		fmt.Println(err)
	}

	io.PrintMemUsage()
}
