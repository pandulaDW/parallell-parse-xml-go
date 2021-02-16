package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

func readAndUnmarshalByStream(reader *bufio.Reader, ch chan<- *string, model GliefModel) int {
	prefix, category, recordsPerRoutine := model.prefix, model.category, model.recordsPerRoutine
	sb := strings.Builder{}
	recordCount := 0
	shouldAppend := false
	recordSets := 0
	if category != "Exception" {
		sb.WriteString(fmt.Sprintf("<%v:%vData>\n", prefix, category))
		sb.WriteString(fmt.Sprintf("<%v:%vRecords>\n", prefix, category))
	} else {
		sb.WriteString("<repex:ReportingExceptionData>")
		sb.WriteString("<repex:ReportingExceptions>")
	}

	// buffered reading and forking goroutines
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		if isRecordStart(line, &model, shouldAppend) {
			shouldAppend = true
			sb.WriteString(line)

			if isRecordStart(line, &model, shouldAppend) {
				recordCount++

				if recordCount == recordsPerRoutine-1 {
					if category != "Exception" {
						sb.WriteString(fmt.Sprintf("</%v:%vRecords>\n", prefix, category))
						sb.WriteString(fmt.Sprintf("</%v:%vData>\n", prefix, category))
					} else {
						sb.WriteString("</repex:ReportingExceptions>")
						sb.WriteString("</repex:ReportingExceptionData>")
					}
					str := sb.String()
					go unmarshalRecords(&str, ch, category)
					sb.Reset()
					recordSets++
					if category != "Exception" {
						sb.WriteString(fmt.Sprintf("<%v:%vData>\n", prefix, category))
						sb.WriteString(fmt.Sprintf("<%v:%vRecords>\n", prefix, category))
					} else {
						sb.WriteString("<repex:ReportingExceptionData>")
						sb.WriteString("<repex:ReportingExceptions>")
					}
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

// isRecordStart checks if the line is a start of a record
func isRecordStart(line string, model *GliefModel, shouldAppend bool) bool {
	var recordStart bool
	prefix, category := model.prefix, model.category
	if category != "Exception" {
		recordStart = strings.Contains(line, fmt.Sprintf("<%v:%vRecord>\n", prefix, category)) || shouldAppend
	} else {
		recordStart = strings.Contains(line, "<repex:Exception>") || shouldAppend
	}
	return recordStart
}
