package processing

import (
	"bufio"
	"fmt"
	"github.com/pandulaDW/parallell-parse-xml-go/models"
	"io"
	"log"
	"strings"
)

func readAndUnmarshalByStream(reader *bufio.Reader, ch chan<- *string, model models.GliefModel) int {
	prefix, category, recordsPerRoutine := model.Prefix, model.Category, model.RecordsPerRoutine
	sb := strings.Builder{}
	recordCount := 0
	shouldAppend := false
	recordSets := 0
	if category != "Exception" {
		sb.WriteString(fmt.Sprintf("<%v:%vData>\n", prefix, category))
		sb.WriteString(fmt.Sprintf("<%v:%vRecords>\n", prefix, category))
	} else {
		sb.WriteString("<repex:ReportingExceptionData>\n")
		sb.WriteString("<repex:ReportingExceptions>\n")
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

			if isRecordEnd(line, &model) {
				recordCount++

				if recordCount == recordsPerRoutine-1 {
					if category != "Exception" {
						sb.WriteString(fmt.Sprintf("</%v:%vRecords>\n", prefix, category))
						sb.WriteString(fmt.Sprintf("</%v:%vData>\n", prefix, category))
					} else {
						sb.WriteString("</repex:ReportingExceptions>\n")
						sb.WriteString("</repex:ReportingExceptionData>\n")
					}
					str := sb.String()
					go unmarshalRecords(&str, ch, category)
					sb.Reset()
					recordSets++
					if category != "Exception" {
						sb.WriteString(fmt.Sprintf("<%v:%vData>\n", prefix, category))
						sb.WriteString(fmt.Sprintf("<%v:%vRecords>\n", prefix, category))
					} else {
						sb.WriteString("<repex:ReportingExceptionData>\n")
						sb.WriteString("<repex:ReportingExceptions>\n")
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
func isRecordStart(line string, model *models.GliefModel, shouldAppend bool) bool {
	var recordStart bool
	prefix, category := model.Prefix, model.Category
	if category != "Exception" {
		recordStart = strings.Contains(line, fmt.Sprintf("<%v:%vRecord>\n", prefix, category)) || shouldAppend
	} else {
		recordStart = strings.Contains(line, "<repex:Exception>\n") || shouldAppend
	}
	return recordStart
}

// isRecordEnd checks if the line is an end of a record
func isRecordEnd(line string, model *models.GliefModel) bool {
	var recordEnd bool
	prefix, category := model.Prefix, model.Category
	if category != "Exception" {
		recordEnd = strings.Contains(line, fmt.Sprintf("</%v:%vRecord>\n", prefix, category))
	} else {
		recordEnd = strings.Contains(line, "</repex:Exception>\n")
	}
	return recordEnd
}
