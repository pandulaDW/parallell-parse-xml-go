package processing

import (
	"encoding/xml"
	"fmt"
	"github.com/pandulaDW/parallell-parse-xml-go/csv"
	"github.com/pandulaDW/parallell-parse-xml-go/models"
	"strings"
)

func unmarshalRR(content *string) *string {
	records := models.RelationshipData{}
	if err := xml.Unmarshal([]byte(*content), &records); err != nil {
		fmt.Println(err)
	}
	sb := strings.Builder{}
	for _, record := range records.RelationshipRecords {
		row := csv.ConvertToCSVRowRR(&record)
		sb.WriteString(row)
	}
	str := sb.String()
	return &str
}

func unmarshalLEI(content *string) *string {
	records := models.LEIData{}
	if err := xml.Unmarshal([]byte(*content), &records); err != nil {
		fmt.Println(err)
	}
	sb := strings.Builder{}
	for _, record := range records.LEIRecords {
		row := csv.ConvertToCSVRowLEI(&record)
		sb.WriteString(row)
		sb.WriteByte('\n')
	}
	str := sb.String()
	return &str
}

func unmarshalRepex(content *string) *string {
	records := models.ReportingExceptionData{}
	if err := xml.Unmarshal([]byte(*content), &records); err != nil {
		fmt.Println(err)
	}
	sb := strings.Builder{}
	for _, record := range records.ReportingExceptions {
		row := csv.ConvertToCSVRowRepex(&record)
		sb.WriteString(row)
		sb.WriteByte('\n')
	}
	str := sb.String()
	return &str
}

func unmarshalRecords(content *string, ch chan<- *string, category string) {
	var rows *string
	switch {
	case category == "Relationship":
		rows = unmarshalRR(content)
	case category == "LEI":
		rows = unmarshalLEI(content)
	case category == "Exception":
		rows = unmarshalRepex(content)
	}
	ch <- rows
}
