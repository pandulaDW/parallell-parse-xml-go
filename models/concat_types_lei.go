package models

// LEIData is the parent of LEIRecords
type LEIData struct {
	LEIRecords []LEIRecord `xml:"LEIRecords>LEIRecord"`
}

// LEIRecord is the parent of Relationship and Registration
type LEIRecord struct {
	LEI          string
	Entity       Entity
	Registration LEIRegistration
}

// Entity defines extracted fields of Entity records
type Entity struct {
	LegalName             string
	OtherEntityName       []string `xml:"OtherEntityNames>OtherEntityName"`
	LegalAddress          LEIAddress
	HeadquartersAddress   LEIAddress
	OtherAddresses        []LEIAddress `xml:"OtherAddresses>OtherAddress"`
	RegistrationAuthority RegistrationAuthority
	LegalJurisdiction     string
	LegalForm             LegalForm
	EntityStatus          string
}

// LEIRegistration defines the extracted fields of Registration records
type LEIRegistration struct {
	InitialRegistrationDate string
	LastUpdateDate          string
	RegistrationStatus      string
	NextRenewalDate         string
	ManagingLOU             string
	ValidationSources       string
	ValidationAuthority     ValidationAuthority
}

// LEIAddress is a subtype of Entity
type LEIAddress struct {
	FirstAddressLine      string
	AdditionalAddressLine string
	City                  string
	Region                string
	Country               string
	PostalCode            string
}

// RegistrationAuthority is a subtype of Entity
type RegistrationAuthority struct {
	RegistrationAuthorityID       string
	RegistrationAuthorityEntityID string
}

// LegalForm is a subtype of Entity
type LegalForm struct {
	EntityLegalFormCode string
	OtherLegalForm      string
}

// ValidationAuthority is a subtype of Entity
type ValidationAuthority struct {
	ValidationAuthorityID       string
	ValidationAuthorityEntityID string
}
