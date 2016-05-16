package enData

import (
	"math/rand"
	"time"
)

// CityPrefix returns a single random city prefix value from the array.
// Possible values include North, East, West, South, New, Lake, and Port.
func CityPrefix() string {
	rand.Seed(time.Now().Unix())
	cityPrefixes := [7]string{
		"North",
		"East",
		"West",
		"South",
		"New",
		"Lake",
		"Port",
	}

	n := rand.Int() % len(cityPrefixes)
	return cityPrefixes[n]
}

// CitySuffix returns a single random city suffix value from the array.
// Possible values include town, ton, land, ville, berg, burgh, borough, bury,
// view, port, mouth, stad, furt, chester, mouth, fort, haven, side, and shire.
func CitySuffix() string {
	rand.Seed(time.Now().Unix())
	citySuffixes := [19]string{
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

	n := rand.Int() % len(citySuffixes)
	return citySuffixes[n]
}

// CityTemplate returns a single string which can be parsed as a template
// such as text/template or html/template. Responses will be in one of the
// following formats:
// 1. {{- .Prefix }} {{.FirstName -}}{{ .Suffix -}}
// 2. {{- .Prefix }} {{ .FirstName -}}
// 3. {{- .FirstName }}{{ .Suffix -}}
// 4. {{- .LastName}}{{ .Suffix -}}
func CityTemplate() string {
	rand.Seed(time.Now().Unix())
	cityTemplates := [4]string{
		`{{- .Prefix }} {{.FirstName -}}{{ .Suffix -}}`,
		`{{- .Prefix }} {{ .FirstName -}}`,
		`{{- .FirstName }}{{ .Suffix -}}`,
		`{{- .LastName}}{{ .Suffix -}}`,
	}

	n := rand.Int() % len(cityTemplates)
	return cityTemplates[n]
}
