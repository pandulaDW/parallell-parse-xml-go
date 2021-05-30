package csv

import (
	"encoding/csv"
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
	rowContent := make([]string, 0, 30)
	rowContent = append(rowContent, r.Relationship.StartNodeID, r.Relationship.StartNodeIDType,
		r.Relationship.EndNodeID, r.Relationship.EndNodeIDType, r.Relationship.RelationshipType)

	for i := 0; i < 2; i++ {
		if len(r.Relationship.RelationshipPeriods) > i {
			rowContent = append(rowContent, r.Relationship.RelationshipPeriods[i].StartDate,
				r.Relationship.RelationshipPeriods[i].EndDate, r.Relationship.RelationshipPeriods[i].PeriodType)
		} else {
			rowContent = append(rowContent, "", "", "")
		}
	}

	rowContent = append(rowContent, r.Relationship.RelationshipStatus)

	for i := 0; i < 2; i++ {
		if len(r.Relationship.RelationshipQualifiers) > i {
			rowContent = append(rowContent, r.Relationship.RelationshipQualifiers[i].QualifierDimension,
				r.Relationship.RelationshipQualifiers[i].QualifierCategory)
		} else {
			rowContent = append(rowContent, "", "")
		}
	}

	for i := 0; i < 2; i++ {
		if len(r.Relationship.RelationshipQuantifiers) > i {
			rowContent = append(rowContent,
				r.Relationship.RelationshipQuantifiers[i].MeasurementMethod,
				fmt.Sprintf("%v", r.Relationship.RelationshipQuantifiers[i].QuantifierAmount),
				r.Relationship.RelationshipQuantifiers[i].QuantifierUnits,
			)
		} else {
			rowContent = append(rowContent, "", "", "")
		}
	}

	rowContent = append(rowContent, r.Registration.InitialRegistrationDate, r.Registration.LastUpdateDate,
		r.Registration.RegistrationStatus, r.Registration.NextRenewalDate, r.Registration.ManagingLOU,
		r.Registration.ValidationSources, r.Registration.ValidationDocuments, r.Registration.ValidationReference,
	)

	sb := strings.Builder{}
	writer := csv.NewWriter(&sb)
	_ = writer.Write(rowContent)
	writer.Flush()

	return sb.String()
}

// ConvertToCSVRowLEI creates a row based on a lei record
func ConvertToCSVRowLEI(lei *models.LEIRecord) string {
	rowContent := make([]string, 0, 40)
	rowContent = append(rowContent, lei.LEI, lei.Entity.LegalName)

	if len(lei.Entity.OtherEntityName) > 0 {
		rowContent = append(rowContent, lei.Entity.OtherEntityName[0])
	} else {
		rowContent = append(rowContent, "")
	}

	rowContent = append(rowContent, lei.Entity.LegalAddress.FirstAddressLine,
		lei.Entity.LegalAddress.AdditionalAddressLine, lei.Entity.LegalAddress.City,
		lei.Entity.LegalAddress.Region, lei.Entity.LegalAddress.Country, lei.Entity.LegalAddress.PostalCode,
		lei.Entity.HeadquartersAddress.FirstAddressLine, lei.Entity.HeadquartersAddress.AdditionalAddressLine,
		lei.Entity.HeadquartersAddress.City, lei.Entity.HeadquartersAddress.Region,
		lei.Entity.HeadquartersAddress.Country, lei.Entity.HeadquartersAddress.PostalCode,
	)

	if len(lei.Entity.OtherAddresses) > 0 {
		rowContent = append(rowContent, lei.Entity.OtherAddresses[0].FirstAddressLine,
			lei.Entity.OtherAddresses[0].AdditionalAddressLine, lei.Entity.OtherAddresses[0].City,
			lei.Entity.OtherAddresses[0].Region, lei.Entity.OtherAddresses[0].Country,
			lei.Entity.OtherAddresses[0].PostalCode)
	} else {
		rowContent = append(rowContent, "", "", "", "", "", "")
	}

	rowContent = append(rowContent,
		lei.Entity.RegistrationAuthority.RegistrationAuthorityID,
		lei.Entity.RegistrationAuthority.RegistrationAuthorityEntityID, lei.Entity.LegalJurisdiction,
		lei.Entity.LegalForm.EntityLegalFormCode, lei.Entity.LegalForm.OtherLegalForm, lei.Entity.EntityStatus,
		lei.Registration.InitialRegistrationDate, lei.Registration.LastUpdateDate, lei.Registration.RegistrationStatus,
		lei.Registration.NextRenewalDate, lei.Registration.ManagingLOU, lei.Registration.ValidationSources,
		lei.Registration.ValidationAuthority.ValidationAuthorityID,
		lei.Registration.ValidationAuthority.ValidationAuthorityEntityID)

	sb := strings.Builder{}
	writer := csv.NewWriter(&sb)
	_ = writer.Write(rowContent)
	writer.Flush()

	return sb.String()
}

// ConvertToCSVRowRepex creates a row based on a repex record
func ConvertToCSVRowRepex(repex *models.Exception) string {
	row := []string{repex.LEI, repex.ExceptionCategory, repex.ExceptionReason}

	sb := strings.Builder{}
	writer := csv.NewWriter(&sb)
	_ = writer.Write(row)
	writer.Flush()

	return sb.String()
}
