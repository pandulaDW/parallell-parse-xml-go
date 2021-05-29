package processing

import (
	"compress/gzip"
	"fmt"
	"github.com/pandulaDW/parallell-parse-xml-go/csv"
	"github.com/pandulaDW/parallell-parse-xml-go/io"
	"github.com/pandulaDW/parallell-parse-xml-go/models"
	"log"
	"os"
	"path/filepath"
	"time"
)

// ConcurrentProcessing function acts as the main function to process a given file
func ConcurrentProcessing(model models.GliefModel, inStage models.InputStage) {
	start := time.Now()
	ch := make(chan *string)

	bufferedReader, err := io.CreateBufferedReader(model, inStage)
	if err != nil {
		log.Fatal(err)
	}

	zipFile, err := os.OpenFile(model.GZipFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	zipWriter := gzip.NewWriter(zipFile)
	zipWriter.Name = filepath.Base(model.CsvFileName)

	recordSets := readAndUnmarshalByStream(bufferedReader, ch, model)

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
		log.Fatal(err)
	}

	fmt.Printf("%d concurrent parses with time taken: %v\n", recordSets, time.Since(start))
}
