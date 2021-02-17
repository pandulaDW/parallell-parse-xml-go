package models

// RelationshipData is the parent of RelationshipRecord
type RelationshipData struct {
	RelationshipRecords []RelationshipRecord `xml:"RelationshipRecords>RelationshipRecord"`
}

// RelationshipRecord is the parent of Relationship and Registration
type RelationshipRecord struct {
	Relationship Relationship
	Registration Registration
}

// Relationship defines extracted fields of Relationship records
type Relationship struct {
	StartNodeID             string `xml:"StartNode>NodeID"`
	StartNodeIDType         string `xml:"StartNode>NodeIDType"`
	EndNodeID               string `xml:"EndNode>NodeID"`
	EndNodeIDType           string `xml:"EndNode>NodeIDType"`
	RelationshipType        string
	RelationshipPeriods     []RelationshipPeriod `xml:"RelationshipPeriods>RelationshipPeriod"`
	RelationshipStatus      string
	RelationshipQualifiers  []RelationshipQualifier  `xml:"RelationshipQualifiers>RelationshipQualifier"`
	RelationshipQuantifiers []RelationshipQuantifier `xml:"RelationshipQuantifiers>RelationshipQuantifier"`
}

// Registration defines the extracted fields of Registration records
type Registration struct {
	InitialRegistrationDate string
	LastUpdateDate          string
	RegistrationStatus      string
	NextRenewalDate         string
	ManagingLOU             string
	ValidationSources       string
	ValidationDocuments     string
	ValidationReference     string
}

// RelationshipPeriod is a sub type of Relationship
type RelationshipPeriod struct {
	StartDate  string
	EndDate    string
	PeriodType string
}

// RelationshipQualifier is a sub type of Relationship
type RelationshipQualifier struct {
	QualifierDimension string
	QualifierCategory  string
}

// RelationshipQuantifier is a sub type of Relationship
type RelationshipQuantifier struct {
	MeasurementMethod string
	QuantifierAmount  float32
	QuantifierUnits   string
}
