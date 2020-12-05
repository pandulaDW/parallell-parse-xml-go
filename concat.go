package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"strings"
)

func unmarshalRecords(content string, ch chan<- string) {
	records := RelationshipData{}

	if err := xml.Unmarshal([]byte(content), &records); err != nil {
		fmt.Println(err)
	}

	sb := strings.Builder{}
	for _, record := range records.RelationshipRecords {
		row := convertToCSVRow(&record)
		sb.WriteString(row)
		sb.WriteByte('\n')
	}
	ch <- sb.String()
}

func readAndUnmarshalByStream(reader *bufio.Reader, recordsPerRoutine int, ch chan<- string) int {
	sb := strings.Builder{}
	recordCount := 0
	shouldAppend := false
	recordSets := 0
	sb.WriteString("<rr:RelationshipData>\n")
	sb.WriteString("<rr:RelationshipRecords>\n")

	// buffered reading and forking goroutines
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		if strings.Contains(line, "<rr:RelationshipRecord>") || shouldAppend {
			shouldAppend = true
			sb.WriteString(line)

			if strings.Contains(line, "</rr:RelationshipRecord>") {
				recordCount++

				if recordCount == recordsPerRoutine-1 {
					sb.WriteString("</rr:RelationshipRecords>")
					sb.WriteString("</rr:RelationshipData>")
					go unmarshalRecords(sb.String(), ch)
					sb.Reset()
					recordSets++
					sb.WriteString("<rr:RelationshipData>\n")
					sb.WriteString("<rr:RelationshipRecords>\n")
					recordCount = 0
				}
			}
		}
		if err == io.EOF {
			if sb.Len() > 0 {
				go unmarshalRecords(sb.String(), ch)
				recordSets++
			}
			return recordSets
		}
	}
}
