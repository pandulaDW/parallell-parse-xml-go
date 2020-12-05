package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func concurrent() {
	start := time.Now()
	ch := make(chan string)

	// bufferedReader, err := createBufferedReader("data/20201203-gleif-concatenated-file-rr.xml.5fc8c1302bde7.zip")
	inFile, err := os.Open("data/20201202-gleif-concatenated-file-lei2.xml")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	bufferedReader := bufio.NewReader(inFile)
	recordSets := readAndUnmarshalByStreamLEI(bufferedReader, 1000, ch)

	outFile, err := os.OpenFile("data/sampleWrite.csv", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	count := 0

	// header := createCsvHeader()
	// sb.WriteString(header)
	// sb.WriteByte('\n')

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

func main() {
	concurrent()
}
