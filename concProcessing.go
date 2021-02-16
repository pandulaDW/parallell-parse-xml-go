package main

import (
	"fmt"
	"log"
	"time"
)

func concurrentProcessing(model GliefModel, inStage InputStage, outStage OutputStage) {
	start := time.Now()
	ch := make(chan *string)

	bufferedReader, err := createBufferedReader(model, inStage)
	if err != nil {
		log.Fatal(err)
	}

	recordSets := readAndUnmarshalByStream(bufferedReader, ch, model)
	bufferedWriter, err := createBufferedWriter(model, outStage)
	if err != nil {
		log.Fatal(err)
	}

	header := createCsvHeader(&model)
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
	fmt.Printf("%d concurrent parses with time taken: %v", recordSets, time.Since(start))
}
