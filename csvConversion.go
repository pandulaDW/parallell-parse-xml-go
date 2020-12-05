package main

import (
	"fmt"
	"reflect"
	"strings"
)

func createCsvHeader() string {
	colNames := CsvColNames{}
	colNamesType := reflect.ValueOf(colNames).Type()
	sb := make([]string, colNamesType.NumField())

	for i := 0; i < colNamesType.NumField(); i++ {
		sb[i] = colNamesType.Field(i).Name
	}

	return strings.Join(sb, ",")
}

func convertToCSVRow(r *RelationshipRecord) string {
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

func replaceDoubleQuotes(s string) string {
	return strings.ReplaceAll(s, "\"", "'")
}
