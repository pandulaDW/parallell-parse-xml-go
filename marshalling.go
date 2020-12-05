package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

// Address type
type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

// Person type
type Person struct {
	XMLName   xml.Name `xml:"person"`
	ID        int      `xml:"id,attr"`
	Code      string   `xml:"code,attr"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Age       int      `xml:"age"`
	Height    float32  `xml:"height,omitempty"`
	Married   bool
	Address
	Comment string `xml:",comment"`
}

// ExampleMarshalIndent type
func ExampleMarshalIndent() {
	v := &Person{ID: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Code = "12WP"
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}

	output, err := xml.MarshalIndent(v, "", "	")
	// output, err := xml.Marshal(v)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
}

// ExampleEncoder type
func ExampleEncoder() {
	v := &Person{ID: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Code = "12WP"
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}

	file, err := os.OpenFile("sample.xml", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	enc := xml.NewEncoder(file)
	enc.Indent("", "	")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
