package main

// CsvColNamesRR contains column names of the written csv file of RR file
type CsvColNamesRR struct {
	StartNodeID                  string
	StartNodeIDType              string
	EndNodeID                    string
	EndNodeIDType                string
	RelationshipType             string
	RelationshipPeriodStartDate1 string
	RelationshipPeriodEndDate1   string
	RelationshipPeriodType1      string
	RelationshipPeriodStartDate2 string
	RelationshipPeriodEndDate2   string
	RelationshipPeriodType2      string
	RelationshipStatus           string
	QualifierDimension1          string
	QualifierCategory1           string
	QualifierDimension2          string
	QualifierCategory2           string
	MeasurementMethod1           string
	QuantifierAmount1            string
	QuantifierUnits1             string
	MeasurementMethod2           string
	QuantifierAmount2            string
	QuantifierUnits2             string
	InitialRegistrationDate      string
	LastUpdateDate               string
	RegistrationStatus           string
	NextRenewalDate              string
	ManagingLOU                  string
	ValidationSources            string
	ValidationDocuments          string
	ValidationReference          string
}

// CsvColNamesRepex contains column names of the written csv file Repex file
type CsvColNamesRepex struct {
	LEI               string
	ExceptionCategory string
	ExceptionReason   string
}
