package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

func readAndUnmarshalByStream(reader *bufio.Reader, recordsPerRoutine int, ch chan<- *string, prefix, category string) int {
	sb := strings.Builder{}
	recordCount := 0
	shouldAppend := false
	recordSets := 0
	sb.WriteString(fmt.Sprintf("<%v:%vData>\n", prefix, category))
	sb.WriteString(fmt.Sprintf("<%v:%vRecords>\n", prefix, category))

	// buffered reading and forking goroutines
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		if strings.Contains(line, fmt.Sprintf("<%v:%vRecord>\n", prefix, category)) || shouldAppend {
			shouldAppend = true
			sb.WriteString(line)

			if strings.Contains(line, fmt.Sprintf("</%v:%vRecord>\n", prefix, category)) {
				recordCount++

				if recordCount == recordsPerRoutine-1 {
					sb.WriteString(fmt.Sprintf("</%v:%vRecords>\n", prefix, category))
					sb.WriteString(fmt.Sprintf("</%v:%vData>\n", prefix, category))
					str := sb.String()
					go unmarshalRecords(&str, ch, category)
					sb.Reset()
					recordSets++
					sb.WriteString(fmt.Sprintf("<%v:%vData>\n", prefix, category))
					sb.WriteString(fmt.Sprintf("<%v:%vRecords>\n", prefix, category))
					recordCount = 0
				}
			}
		}
		if err == io.EOF {
			if sb.Len() > 0 {
				str := sb.String()
				go unmarshalRecords(&str, ch, category)
				recordSets++
			}
			return recordSets
		}
	}
}
