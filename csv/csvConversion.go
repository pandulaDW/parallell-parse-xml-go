package csv

import (
	"fmt"
	"github.com/pandulaDW/parallell-parse-xml-go/models"
	"reflect"
	"strings"
)

// CreateCsvHeader will create a header based on the file type
func CreateCsvHeader(model *models.GliefModel) string {
	var colNames interface{}
	switch {
	case model.Prefix == "rr":
		colNames = ColNamesRR{}
	case model.Prefix == "lei":
		colNames = ColNamesLEI{}
	case model.Prefix == "repex":
		colNames = ColNamesRepex{}
	}
	colNamesType := reflect.ValueOf(colNames).Type()
	sb := make([]string, colNamesType.NumField())

	for i := 0; i < colNamesType.NumField(); i++ {
		sb[i] = colNamesType.Field(i).Name
	}

	return strings.Join(sb, ",")
}

// ConvertToCSVRowRR creates a row based on a relationship record
func ConvertToCSVRowRR(r *models.RelationshipRecord) string {
	rowContent := make([]string, 10)
	rowContent[0] = fmt.Sprintf(`"%v","%v","%v","%v","%v"`,
		r.Relationship.StartNodeID,
		r.Relationship.StartNodeIDType,
		r.Relationship.EndNodeID,
		r.Relationship.EndNodeIDType,
		r.Relationship.RelationshipType,
	)

	for i := 0; i < 2; i++ {
		if len(r.Relationship.RelationshipPeriods) > i {
			rowContent[1+i] = fmt.Sprintf(`"%v","%v","%v"`,
				r.Relationship.RelationshipPeriods[i].StartDate,
				r.Relationship.RelationshipPeriods[i].EndDate,
				r.Relationship.RelationshipPeriods[i].PeriodType,
			)
		} else {
			rowContent[1+i] = fmt.Sprintf(`"","",""`)
		}
	}

	rowContent[3] = fmt.Sprintf(`"%v"`, r.Relationship.RelationshipStatus)

	for i := 0; i < 2; i++ {
		if len(r.Relationship.RelationshipQualifiers) > i {
			rowContent[4+i] = fmt.Sprintf(`"%v","%v"`,
				r.Relationship.RelationshipQualifiers[i].QualifierDimension,
				r.Relationship.RelationshipQualifiers[i].QualifierCategory,
			)
		} else {
			rowContent[4+i] = fmt.Sprintf(`"",""`)
		}
	}

	for i := 0; i < 2; i++ {
		if len(r.Relationship.RelationshipQuantifiers) > i {
			rowContent[6+i] = fmt.Sprintf(`"%v","%v","%v"`,
				r.Relationship.RelationshipQuantifiers[i].MeasurementMethod,
				r.Relationship.RelationshipQuantifiers[i].QuantifierAmount,
				r.Relationship.RelationshipQuantifiers[i].QuantifierUnits,
			)
		} else {
			rowContent[6+i] = fmt.Sprintf(`"","",""`)
		}
	}

	rowContent[8] = fmt.Sprintf(`"%v","%v","%v","%v","%v","%v","%v","%v"`,
		r.Registration.InitialRegistrationDate,
		r.Registration.LastUpdateDate,
		r.Registration.RegistrationStatus,
		r.Registration.NextRenewalDate,
		r.Registration.ManagingLOU,
		r.Registration.ValidationSources,
		r.Registration.ValidationDocuments,
		replaceDoubleQuotes(r.Registration.ValidationReference),
	)

	return strings.Join(rowContent, ",")
}

// ConvertToCSVRowLEI creates a row based on a lei record
func ConvertToCSVRowLEI(lei *models.LEIRecord) string {
	rowContent := make([]string, 10)
	rowContent[0] = fmt.Sprintf(`"%v","%v"`,
		lei.LEI,
		lei.Entity.LegalName,
	)

	if len(lei.Entity.OtherEntityName) > 0 {
		rowContent[1] = fmt.Sprintf(`"%v","%v"`, lei.Entity.OtherEntityName[0])
	} else {
		rowContent[1] = fmt.Sprintf(`""`)
	}

	rowContent[2] = fmt.Sprintf(`"%v","%v","%v","%v","%v","%v","%v","%v","%v","%v","%v","%v"`,
		lei.Entity.LegalAddress.FirstAddressLine,
		lei.Entity.LegalAddress.AdditionalAddressLine,
		lei.Entity.LegalAddress.City,
		lei.Entity.LegalAddress.Region,
		lei.Entity.LegalAddress.Country,
		lei.Entity.LegalAddress.PostalCode,
		lei.Entity.HeadquartersAddress.FirstAddressLine,
		lei.Entity.HeadquartersAddress.AdditionalAddressLine,
		lei.Entity.HeadquartersAddress.City,
		lei.Entity.HeadquartersAddress.Region,
		lei.Entity.HeadquartersAddress.Country,
		lei.Entity.HeadquartersAddress.PostalCode,
	)

	if len(lei.Entity.OtherAddresses) > 0 {
		rowContent[3] = fmt.Sprintf(`"%v","%v","%v","%v","%v","%v"`,
			lei.Entity.OtherAddresses[0].FirstAddressLine,
			lei.Entity.OtherAddresses[0].AdditionalAddressLine,
			lei.Entity.OtherAddresses[0].City,
			lei.Entity.OtherAddresses[0].Region,
			lei.Entity.OtherAddresses[0].Country,
			lei.Entity.OtherAddresses[0].PostalCode,
		)
	} else {
		rowContent[3] = fmt.Sprintf(`"","","","","",""`)
	}

	rowContent[4] = fmt.Sprintf(`"%v","%v","%v","%v","%v","%v","%v","%v","%v","%v","%v","%v","%v","%v"`,
		lei.Entity.RegistrationAuthority.RegistrationAuthorityID,
		lei.Entity.RegistrationAuthority.RegistrationAuthorityEntityID,
		lei.Entity.LegalJurisdiction,
		lei.Entity.LegalForm.EntityLegalFormCode,
		lei.Entity.LegalForm.OtherLegalForm,
		lei.Entity.EntityStatus,
		lei.Registration.InitialRegistrationDate,
		lei.Registration.LastUpdateDate,
		lei.Registration.RegistrationStatus,
		lei.Registration.NextRenewalDate,
		lei.Registration.ManagingLOU,
		lei.Registration.ValidationSources,
		lei.Registration.ValidationAuthority.ValidationAuthorityID,
		lei.Registration.ValidationAuthority.ValidationAuthorityEntityID,
	)

	return strings.Join(rowContent, ",")
}

// ConvertToCSVRowRepex creates a row based on a repex record
func ConvertToCSVRowRepex(repex *models.Exception) string {
	row := fmt.Sprintf(`"%v","%v","%v"`,
		repex.LEI,
		repex.ExceptionCategory,
		repex.ExceptionReason,
	)
	return row
}

func replaceDoubleQuotes(s string) string {
	return strings.ReplaceAll(s, "\"", "'")
}
