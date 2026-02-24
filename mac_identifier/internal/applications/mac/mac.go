package mac

import (
	"errors"
	"strings"
)

var ErrNotFound = errors.New("vendor not found")

type Service interface {
	Lookup(mac string) (*LookupResult, error)
}

type service struct {
	Repository Repository
}

func NewService(r Repository) Service {
	return service{r}
}

// Lookup finds the vendor for a MAC address.
// Accepts formats: AA:BB:CC:DD:EE:FF, AA-BB-CC-DD-EE-FF, AABBCCDDEEFF
func (s service) Lookup(mac string) (*LookupResult, error) {
	normalized := strings.ToUpper(strings.NewReplacer(":", "", "-", "").Replace(mac))
	if len(normalized) < 6 {
		return nil, errors.New("invalid MAC address")
	}
	prefix := normalized[:6]

	vendor, err := s.Repository.FindVendor(prefix)
	if err != nil {
		return nil, err
	}

	return &LookupResult{MAC: mac, Vendor: vendor}, nil
}
