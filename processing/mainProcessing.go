package processing

import (
	"fmt"
	"github.com/pandulaDW/parallell-parse-xml-go/csv"
	"github.com/pandulaDW/parallell-parse-xml-go/io"
	"github.com/pandulaDW/parallell-parse-xml-go/models"
	"log"
	"time"
)

// ConcurrentProcessing function acts as the main function to process a given file
func ConcurrentProcessing(model models.GliefModel, inStage models.InputStage, outStage models.OutputStage) {
	start := time.Now()
	ch := make(chan *string)

	bufferedReader, err := io.CreateBufferedReader(model, inStage)
	if err != nil {
		log.Fatal(err)
	}

	recordSets := readAndUnmarshalByStream(bufferedReader, ch, model)
	bufferedWriter, err := io.CreateBufferedWriter(model, outStage)
	if err != nil {
		log.Fatal(err)
	}

	header := csv.CreateCsvHeader(&model)
	_, _ = bufferedWriter.WriteString(header)
	_ = bufferedWriter.WriteByte('\n')

	count := 0
	for count < recordSets {
		recordSet := <-ch
		_, _ = bufferedWriter.WriteString(*recordSet)
		count++

		if count%50 == 0 {
			_ = bufferedWriter.Flush()
		}
	}

	_ = bufferedWriter.Flush()
	fmt.Printf("%d concurrent parses with time taken: %v\n", recordSets, time.Since(start))
}
