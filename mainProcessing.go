package main

import (
	"bufio"
	"fmt"
	"log"
	"time"
)

func concurrentProcessing(prefix, category string) {
	start := time.Now()
	ch := make(chan *string)

	bufferedReader, err := createBufferedReader("data/20201202-gleif-concatenated-file-lei2.xml.5fc7579cab4ee.zip")
	if err != nil {
		log.Fatal(err)
	}

	recordSets := readAndUnmarshalByStream(bufferedReader, 4000, ch, prefix, category)

	zipWriter, zipFile := createGzipWriter("file.zip", "testfile.csv")
	bufferedWriter := bufio.NewWriter(zipWriter)

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
	zipWriter.Close()
	zipFile.Close()
	fmt.Printf("%d concurrent parses with time taken: %v", recordSets, time.Since(start))
}
