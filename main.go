package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func concurrent() {
	start := time.Now()
	ch := make(chan string)
	bufferedReader, err := createBufferedReader("data/20201203-gleif-concatenated-file-rr.xml.5fc8c1302bde7.zip")
	if err != nil {
		panic(err)
	}

	recordSets := readAndUnmarshalByStream(bufferedReader, 500, ch)

	sb := strings.Builder{}
	count := 0

	header := createCsvHeader()
	sb.WriteString(header)
	sb.WriteByte('\n')

	for count < recordSets {
		recordSet := <-ch
		sb.WriteString(recordSet)
		count++
	}

	ioutil.WriteFile("data/sampleWrite.csv", []byte(sb.String()), 0666)
	fmt.Printf("%d concurrent parses with time taken: %v", recordSets, time.Since(start))
}

func main() {
	test()
}
