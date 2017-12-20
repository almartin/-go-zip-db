package ZipDb

import (
	"errors"
	"log"
)

// Goroutine safe lookup function (reader-only). Returns Geolocation structure and any errors.
//
// type GeoLocation struct {
//	 Zip       string
//	 Latitude  float32
//	 Longitude float32
//	 Name      string
//	 State     string
// }

func Lookup(zip string) (*GeoLocation, error) {
	log.Printf("Entering Zip.Lookup with %s\n", zip)
	log.Printf("Found %+v\n", m[zip])
	i, ok := m[zip]
	if !ok {
		return nil, errors.New("Unable to find zip requested")
	}
	return &i, nil
}
