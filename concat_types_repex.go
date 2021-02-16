package main

// RepexData holds all the exception data
type RepexData struct {
	ReportingExceptions []Exception `xml:"ReportingExceptions>Exception"`
}

// Exception represents one exception record
type Exception struct {
	LEI               string
	ExceptionCategory string
	ExceptionReason   string
}
