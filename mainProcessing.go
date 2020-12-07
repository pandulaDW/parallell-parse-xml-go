package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func concurrentProcessing(prefix, category string) {
	start := time.Now()
	ch := make(chan string)

	bufferedReader, err := createBufferedReader("data/20201203-gleif-concatenated-file-rr.xml.5fc8c1302bde7.zip")
	recordSets := readAndUnmarshalByStream(bufferedReader, 1000, ch, prefix, category)

	outFile, err := os.OpenFile("data/sampleWrite2.csv", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	count := 0

	header := createCsvHeader()
	writer.WriteString(header)
	writer.WriteByte('\n')

	for count < recordSets {
		recordSet := <-ch
		writer.WriteString(recordSet)
		count++

		if count%50 == 0 {
			writer.Flush()
		}
	}

	writer.Flush()
	fmt.Printf("%d concurrent parses with time taken: %v", recordSets, time.Since(start))
}
