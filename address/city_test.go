package address

import (
	"fmt"
	"testing"
)

func TestCityString(t *testing.T) {
	c := FakeCity()
	if c == nil {
		t.Errorf("Failed to generate a new city")
	}

	str, err := c.String()
	if err != nil {
		t.Errorf("Failed to convert city to a string.")
	}

	switch c.Template {
	case `{{- .Prefix }} {{.FirstName -}}{{ .Suffix -}}`:
		if fmt.Sprintf("%s %s%s", c.Prefix, c.FirstName, c.Suffix) != str {
			t.Errorf(`%s did not match template: {{- .Prefix }} {{.FirstName -}}{{ .Suffix -}}`, str)
		}
	case `{{- .Prefix }} {{ .FirstName -}}`:
		if fmt.Sprintf("%s %s", c.Prefix, c.FirstName) != str {
			t.Errorf(`%s did not match template: {{- .Prefix }} {{ .FirstName -}}`, str)
		}
	case `{{- .FirstName }}{{ .Suffix -}}`:
		if fmt.Sprintf("%s%s", c.FirstName, c.Suffix) != str {
			t.Errorf(`%s did not match template: {{- .FirstName }}{{ .Suffix -}}`, str)
		}
	case `{{- .LastName}}{{ .Suffix -}}`:
		if fmt.Sprintf("%s%s", c.LastName, c.Suffix) != str {
			t.Errorf(`%s did not match template: {{- .LastName}}{{ .Suffix -}}`, str)
		}
	default:
		t.Errorf("Failed to parse City to String")
	}
}
