package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func test() {
	start := time.Now()
	content, _ := ioutil.ReadFile("data/20201202-gleif-concatenated-file-lei2.xml")
	records := LEIData{}

	if err := xml.Unmarshal(content, &records); err != nil {
		log.Println(err)
	}

	f, err := os.OpenFile("data/sampleWrite2.csv", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	writer := bufio.NewWriter(f)

	for i, record := range records.LEIRecords {
		recordStr := fmt.Sprintf("%v -> %v\n", record.LEI, record.Entity.LegalName)
		writer.WriteString(recordStr)

		if i%50 == 0 {
			writer.Flush()
		}
	}

	// write any remaining data to the underlying writer
	writer.Flush()

	fmt.Println(len(records.LEIRecords), " records were parsed")
	fmt.Println("time taken: ", time.Since(start))
}
