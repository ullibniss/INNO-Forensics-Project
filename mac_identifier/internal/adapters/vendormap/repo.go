package vendormap

import (
	"bufio"
	"fmt"
	"mac_identifier/internal/applications/mac"
	"os"
	"strings"
)

type VendorMapRepo struct {
	vendors map[string]string
}

func NewVendorMapRepo(path string) (*VendorMapRepo, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open vendor map: %w", err)
	}
	defer f.Close()

	vendors := make(map[string]string, 24000)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "\t", 2)
		if len(parts) == 2 {
			vendors[strings.ToUpper(parts[0])] = parts[1]
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read vendor map: %w", err)
	}
	return &VendorMapRepo{vendors: vendors}, nil
}

func (r *VendorMapRepo) FindVendor(prefix string) (string, error) {
	vendor, ok := r.vendors[strings.ToUpper(prefix)]
	if !ok {
		return "", mac.ErrNotFound
	}
	return vendor, nil
}
