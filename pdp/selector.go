package pdp

import (
	"net/url"

	log "github.com/sirupsen/logrus"
)

// Selector provides a generic way to access external data may required
// by policies.
type Selector interface {
	// Enabled returns true for active selector. Disabled selector isn't
	// initialized and can't be used in policies.
	Enabled() bool
	// Scheme returns a name of URI scheme associated with selector.
	Scheme() string
	// Initialize is called for all registered and enabled selectors
	// by InitializeSelectors.
	Initialize()
	// SelectorFunc returns selector expression for given URI,
	// set of arguments and desired result type. The last two arguments
	// define the values selector to return in case of missing content value
	// or other errors
	SelectorFunc(*url.URL, []Expression, Type, Expression, Expression) (Expression, error)
}

var selectorMap = make(map[string]Selector)

// MakeSelector returns new selector for given uri with path as a set of
// arguments and desired result type.
func MakeSelector(uri *url.URL, path []Expression, t Type, def, err Expression) (Expression, error) {
	s := GetSelector(uri.Scheme)
	if s == nil {
		return nil, newUnsupportedSelectorSchemeError(uri)
	}
	if !s.Enabled() {
		return nil, newDisabledSelectorError(s)
	}
	return s.SelectorFunc(uri, path, t, def, err)
}

// RegisterSelector puts given selector to PDP's registry.
func RegisterSelector(s Selector) {
	selectorMap[s.Scheme()] = s
}

// GetSelector returns selector registered for given schema.
func GetSelector(scheme string) Selector {
	if s, ok := selectorMap[scheme]; ok {
		return s
	}

	return nil
}

// InitializeSelectors initializes all registered and enabled selectors.
func InitializeSelectors() {
	for _, s := range selectorMap {
		if s.Enabled() {
			log.WithField("scheme", s.Scheme()).Info("initializing selector")
			s.Initialize()
		}
	}
}
