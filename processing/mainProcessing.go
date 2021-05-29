package processing

import (
	"fmt"
	"github.com/pandulaDW/parallell-parse-xml-go/csv"
	"github.com/pandulaDW/parallell-parse-xml-go/io"
	"github.com/pandulaDW/parallell-parse-xml-go/models"
	"runtime"
	"time"
)

// ConcurrentProcessing function acts as the main function to process a given file
func ConcurrentProcessing(model models.GliefModel, inStage models.InputStage) error {
	start := time.Now()
	ch := make(chan *string)

	bufferedReader, err := io.CreateBufferedReader(model, inStage)
	if err != nil {
		return err
	}

	recordSets := readAndUnmarshalByStream(bufferedReader, ch, model)
	zipWriter, err := io.CreateZipFileWriter(model)
	if err != nil {
		return err
	}

	header := csv.CreateCsvHeader(&model)
	_, _ = zipWriter.Write([]byte(header))
	_, _ = zipWriter.Write([]byte("\n"))

	count := 0
	for count < recordSets {
		recordSet := <-ch
		_, _ = zipWriter.Write([]byte(*recordSet))
		count++

		if count%50 == 0 {
			_ = zipWriter.Flush()
		}
	}
	_ = zipWriter.Flush()
	err = zipWriter.Close()
	if err != nil {
		return err
	}

	fmt.Printf("%d concurrent parses with time taken: %v\n", recordSets, time.Since(start))
	return nil
}

// ProcessFile will go through the full flow for a given type
func ProcessFile(_type string) error {
	var model *models.GliefModel
	switch _type {
	case "rr":
		model = models.CreateRelationshipModel()
	case "lei":
		model = models.CreateLEIModel()
	default:
		model = models.CreateReportingExceptionModel()
	}

	err := ConcurrentProcessing(*model, models.S3XMLFileRead)
	if err != nil {
		return err
	}
	runtime.GC()

	err = io.WriteFileToS3(*model)
	if err != nil {
		return err
	}

	return nil
}
