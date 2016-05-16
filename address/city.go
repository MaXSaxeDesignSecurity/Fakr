package address

import "github.com/kkirsche/Fakr/address/locale/en"

// City represents the information you would
type City struct {
	Prefix    string
	Suffix    string
	FirstName string
	LastName  string
	Template  string
}

// FakeCity generates a fake city from the formats and
func FakeCity() *City {
	return &City{
		Prefix:   enData.CityPrefix(),
		Suffix:   enData.CitySuffix(),
		Template: enData.CityTemplate(),
	}
}
