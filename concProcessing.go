package main

import (
	"fmt"
	"log"
	"time"
)

func concurrentProcessing(model GliefModel, method string) {
	start := time.Now()
	ch := make(chan *string)

	bufferedReader, err := createBufferedReader(model, method)
	if err != nil {
		log.Fatal(err)
	}

	recordSets := readAndUnmarshalByStream(bufferedReader, ch, model)
	bufferedWriter := createGzipWriter(model)

	header := createCsvHeader()
	bufferedWriter.WriteString(header)
	bufferedWriter.WriteByte('\n')

	count := 0
	for count < recordSets {
		recordSet := <-ch
		bufferedWriter.WriteString(*recordSet)
		count++

		if count%50 == 0 {
			bufferedWriter.Flush()
		}
	}

	bufferedWriter.Flush()
	fmt.Printf("%d concurrent parses with time taken: %v", recordSets, time.Since(start))
}
