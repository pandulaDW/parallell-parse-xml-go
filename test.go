package main

import (
	"encoding/xml"
	"github.com/pandulaDW/parallell-parse-xml-go/csv"
	"io/ioutil"
	"log"
	"strings"
)

func test() {
	content, err := ioutil.ReadFile("data/leiSample.xml")
	if err != nil {
		log.Fatal(err)
	}

	records := LEIData{}
	err = xml.Unmarshal(content, &records)
	if err != nil {
		log.Fatal(err)
	}

	model := GliefModel{prefix: "lei"}
	header := csv.createCsvHeader(&model)
	sb := strings.Builder{}
	sb.WriteString(header + "\n")

	for _, record := range records.LEIRecords {
		sb.WriteString(csv.convertToCSVRowLEI(&record) + "\n")
	}

	_ = ioutil.WriteFile("data/leiCSV.csv", []byte(sb.String()), 0666)
}
