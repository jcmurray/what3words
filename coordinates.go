// -*- mode: go; coding: utf-8 -*-
// vi: set syntax=go :
// cSpell.language:en-GB
// cSpell:disable

package what3words

import (
	"fmt"

	"github.com/juju/errors"
)

// Coordinates struct with latitude and longitude.
// Tags are used to map from the JSON response.
type Coordinates struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

// NewCoordinates return a new Coordinates object.
// constructed from latitude and longitude values.
// It verifies range of latitude and longitude are correct.
func NewCoordinates(lat float64, lon float64) (*Coordinates, error) {

	if lat < -90.0 || lat > 90.0 {
		return nil, errors.New("Latitude must be in range [-90, 90]")
	}

	if lon < -180.0 {
		return nil, errors.New("Longitude must be in > -180")
	}
	return &Coordinates{
		Latitude:  lat,
		Longitude: lon,
	}, nil
}

// SetLat sets the Latitude in a Coordinates object.
// verifying range of latitude is correct.
func (coord *Coordinates) SetLat(lat float64) error {
	if lat < -90.0 || lat > 90.0 {
		return errors.New("Latitude must be in range [-90, 90]")
	}
	coord.Latitude = lat
	return nil
}

// SetLon sets the Longitude in a Coordinates object.
// verifying range of longitude is correct.
func (coord *Coordinates) SetLon(lon float64) error {
	if lon < -180.0 {
		return errors.New("Longitude must be > -180")
	}
	coord.Longitude = lon
	return nil
}

// String returns a string suitable for a URL parameter.
func (coord *Coordinates) String() string {
	return fmt.Sprintf("%.13f,%6f", coord.Latitude, coord.Longitude)
}
