// -*- mode: go; coding: utf-8 -*-
// vi: set syntax=go :
// cSpell.language:en-GB
// cSpell:disable

package what3words

// Coordinates struct with latitude and longitude.
// Tags are used to map from the JSON response.
type Coordinates struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

// CoordinatesInterface interface definition
type CoordinatesInterface interface {
	SetLat(lat float64) error
	SetLon(lon float64) error
	String() string
}

// Verify this object implement the necessary interfaces
var _ CoordinatesInterface = &Coordinates{}
