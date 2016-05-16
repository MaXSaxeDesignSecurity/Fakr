package address

import (
	"bytes"
	"html/template"

	"github.com/kkirsche/Fakr/address/locale/en"
)

// City represents the information you would
type City struct {
	Prefix    string
	Suffix    string
	FirstName string
	LastName  string
	Template  string
}

// FakeCity generates information necessary for the fake city
func FakeCity() *City {
	return &City{
		Prefix:   enData.CityPrefix(),
		Suffix:   enData.CitySuffix(),
		Template: enData.CityTemplate(),
	}
}

func (c *City) String() (string, error) {
	t, err := template.New("city").Parse(c.Template)
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer([]byte{})
	err = t.Execute(buf, c)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
