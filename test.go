package main

import (
	"encoding/xml"
	"github.com/pandulaDW/parallell-parse-xml-go/csv"
	"github.com/pandulaDW/parallell-parse-xml-go/models"
	"io/ioutil"
	"log"
	"strings"
)

func test() {
	content, err := ioutil.ReadFile("data/leiSample.xml")
	if err != nil {
		log.Fatal(err)
	}

	records := models.LEIData{}
	err = xml.Unmarshal(content, &records)
	if err != nil {
		log.Fatal(err)
	}

	model := models.GliefModel{Prefix: "lei"}
	header := csv.CreateCsvHeader(&model)
	sb := strings.Builder{}
	sb.WriteString(header + "\n")

	for _, record := range records.LEIRecords {
		sb.WriteString(csv.ConvertToCSVRowLEI(&record) + "\n")
	}

	_ = ioutil.WriteFile("data/leiCSV.csv", []byte(sb.String()), 0666)
}
