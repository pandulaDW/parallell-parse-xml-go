package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

func unmarshalRR(content *string) *string {
	records := RelationshipData{}
	if err := xml.Unmarshal([]byte(*content), &records); err != nil {
		fmt.Println(err)
	}
	sb := strings.Builder{}
	for _, record := range records.RelationshipRecords {
		row := convertToCSVRow(&record)
		sb.WriteString(row)
		sb.WriteByte('\n')
	}
	str := sb.String()
	return &str
}

func unmarshalLEI(content *string) *string {
	records := LEIData{}
	if err := xml.Unmarshal([]byte(*content), &records); err != nil {
		fmt.Println(err)
	}
	sb := strings.Builder{}
	for _, record := range records.LEIRecords {
		row := fmt.Sprintf("%v -> %v", record.LEI, record.Entity.LegalName)
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
	}
	ch <- rows
}
