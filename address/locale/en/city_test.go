package enData

import "testing"

func TestCityPrefixes(t *testing.T) {
	possiblePrefixes := [7]string{
		"North",
		"East",
		"West",
		"South",
		"New",
		"Lake",
		"Port",
	}

	prefix := CityPrefix()
	validPrefix := false

	for _, possible := range possiblePrefixes {
		if prefix == possible {
			validPrefix = true
		}
	}

	if !validPrefix {
		t.Errorf("TestCityPrefixes() | Fakr/address/locale/en/city_test.go:  Invalid city prefix found.")
	}
}

func TestCitySuffixes(t *testing.T) {
	possibleSuffixes := [19]string{
		"town",
		"ton",
		"land",
		"ville",
		"berg",
		"burgh",
		"borough",
		"bury",
		"view",
		"port",
		"mouth",
		"stad",
		"furt",
		"chester",
		"mouth",
		"fort",
		"haven",
		"side",
		"shire",
	}

	suffix := CitySuffix()
	validSuffix := false

	for _, possible := range possibleSuffixes {
		if suffix == possible {
			validSuffix = true
		}
	}

	if !validSuffix {
		t.Errorf("TestCitySuffixes() | Fakr/address/locale/en/city_test.go:  Invalid city suffix found.")
	}
}

func TestCityTemplate(t *testing.T) {
	possibleTemplates := [4]string{
		`{{- .Prefix }} {{.FirstName -}}{{ .Suffix -}}`,
		`{{- .Prefix }} {{ .FirstName -}}`,
		`{{- .FirstName }}{{ .Suffix -}}`,
		`{{- .LastName}}{{ .Suffix -}}`,
	}

	tmpl := CityTemplate()
	validTemplate := false

	for _, possible := range possibleTemplates {
		if tmpl == possible {
			validTemplate = true
		}
	}

	if !validTemplate {
		t.Errorf("address/locale/en/city_test.go: Invalid city template found.")
	}
}
