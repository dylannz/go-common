package nzAddress

import (
	"strconv"
	"strings"
)

// Address is a complete New Zealand address with all the required fields to
// format into a human-readable address string.
type Address struct {
	UnitIdentifier  string `json:"unit_identifier"`
	StreetNumber    int    `json:"street_number"`
	StreetAlpha     string `json:"street_alpha"`
	StreetName      string `json:"street_name"`
	StreetType      string `json:"street_type"`
	StreetDirection string `json:"street_direction"`
	Suburb          string `json:"suburb"`
	City            string `json:"city"`

	BuildingName string `json:"building_name"`
	Floor        string `json:"floor"`
	RDNumber     string `json:"rd_number"`
	Postcode     int    `json:"postcode"`
}

// Street returns the formatted street name + type and direction
func (a Address) Street() string {
	return strings.Trim(
		strings.Join(
			[]string{
				a.StreetName,
				a.StreetType,
				a.StreetDirection,
			},
			" ",
		),
		" ",
	)
}

func titleCase(s string) string {
	return strings.Title(strings.ToLower(s))
}

// Display formats an address into a valid display address
func (a Address) Display() string {
	address := []string{}

	if a.BuildingName != "" {
		unitBuildingName := titleCase(a.BuildingName)
		if a.UnitIdentifier != "" {
			unitBuildingName = strings.ToUpper(a.UnitIdentifier) + " " + unitBuildingName
		}
		address = append(address, unitBuildingName)
	}

	var identifierStreet string
	street := titleCase(a.Street())
	if a.StreetNumber != 0 {
		if a.UnitIdentifier != "" && a.BuildingName == "" {
			identifierStreet += strings.ToUpper(a.UnitIdentifier) + "/"
		}
		identifierStreet += strconv.Itoa(a.StreetNumber) + strings.ToUpper(a.StreetAlpha)
		if street != "" {
			identifierStreet += " "
		}
	}
	identifierStreet += street
	if identifierStreet != "" {
		address = append(address, identifierStreet)
	}

	if a.Suburb != "" {
		address = append(address, titleCase(a.Suburb))
	}
	if a.RDNumber != "" {
		address = append(address, "RD "+strings.ToUpper(a.RDNumber))
	}
	if a.City != "" {
		address = append(address, titleCase(a.City))
	}
	return strings.Trim(strings.Join(address, ", "), " ")
}

// DisplayWithPostcode returns a display address with postcode appended to the
// end.
func (a Address) DisplayWithPostcode() string {
	return a.Display() + " " + strconv.Itoa(a.Postcode)
}
