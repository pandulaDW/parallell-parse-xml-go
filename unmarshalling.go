package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

// ExampleUnMarshal function
func ExampleUnMarshal() {
	type Email struct {
		Where  string `xml:"where,attr"`
		Addr   string
		Domain string
	}

	type Address struct {
		City, State string
	}

	type Result struct {
		XMLName xml.Name `xml:"Person"`
		Name    string   `xml:"FullName"`
		Phone   string
		Email   []Email
		Groups  []string `xml:"Group>Value"`
		Address
	}

	v := Result{Name: "none", Phone: "none"}

	content, err := ioutil.ReadFile("readSample.xml")
	if err != nil {
		log.Fatal(err)
	}

	if err := xml.Unmarshal(content, &v); err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("XMLName: %#v\n", v.XMLName)
	fmt.Printf("Name: %q\n", v.Name)
	fmt.Printf("Phone: %q\n", v.Phone)
	fmt.Printf("Email: %v\n", v.Email)
	fmt.Printf("Groups: %v\n", v.Groups)
	fmt.Printf("Address: %v\n", v.Address.State)
}
