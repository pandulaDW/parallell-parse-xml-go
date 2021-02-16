package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

func test() {
	content, err := ioutil.ReadFile("data/repexSample.xml")
	if err != nil {
		log.Fatal(err)
	}

	records := ReportingExceptionData{}
	err = xml.Unmarshal(content, &records)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(records)
}
