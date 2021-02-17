package csv

// CsvColNamesRR contains column names of the written csv file of RR file
type ColNamesRR struct {
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

// CsvColNamesLEI contains column names of the written csv file LEI data file
type ColNamesLEI struct {
	LEI                                      string
	LegalName                                string
	OtherEntityName1                         string
	LegalAddressFirstAddressLine             string
	LegalAddressAdditionalAddressLine        string
	LegalAddressCity                         string
	LegalAddressRegion                       string
	LegalAddressCountry                      string
	LegalAddressPostalCode                   string
	HeadquartersAddressFirstAddressLine      string
	HeadquartersAddressAdditionalAddressLine string
	HeadquartersAddressCity                  string
	HeadquartersAddressRegion                string
	HeadquartersAddressCountry               string
	HeadquartersAddressPostalCode            string
	OtherAddresses1FirstAddressLine          string
	OtherAddresses1AdditionalAddressLine     string
	OtherAddresses1City                      string
	OtherAddresses1Region                    string
	OtherAddresses1Country                   string
	OtherAddresses1PostalCode                string
	RegistrationAuthorityID                  string
	RegistrationAuthorityEntityID            string
	LegalJurisdiction                        string
	EntityLegalFormCode                      string
	OtherLegalForm                           string
	EntityStatus                             string
	InitialRegistrationDate                  string
	LastUpdateDate                           string
	RegistrationStatus                       string
	NextRenewalDate                          string
	ManagingLOU                              string
	ValidationSources                        string
	ValidationAuthorityID                    string
	ValidationAuthorityEntityID              string
}

// CsvColNamesRepex contains column names of the written csv file Repex file
type ColNamesRepex struct {
	LEI               string
	ExceptionCategory string
	ExceptionReason   string
}
