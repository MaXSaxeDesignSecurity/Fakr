package address

import "github.com/kkirsche/Fakr/locale"

// Address represents the information needed by the address package to determine
// what locale / language information should be used when generating data.
type Address struct {
	locale string
}

// New creates a new Address structure.
func New(locale string) (*Address, error) {
	a := &Address{}
	err := a.RegisterLocale(locale)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// RegisterLocale is used to register the locale in which you are located. The
// locale listings can be seen in Fakr/locale.List()
func (a *Address) RegisterLocale(posixLocale string) error {
	if locale.Verify(posixLocale) {
		a.locale = posixLocale
		return nil
	}

	return nil
}

// Locale returns the currently registered locale of the Address structure.
func (a *Address) Locale() string {
	return a.locale
}
