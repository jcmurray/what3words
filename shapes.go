// -*- mode: go; coding: utf-8 -*-
// vi: set syntax=go :
// cSpell.language:en-GB
// cSpell:disable

package what3words

import "github.com/juju/errors"

// Square defines an area bounded by two coordinates.
// South West and North East points of the square
// define its extent.
// // Tags are used to map from the JSON response.
type Square struct {
	SouthWest *Coordinates `json:"southwest"`
	NorthEast *Coordinates `json:"northeast"`
}

// NewSquare returns a square region
func NewSquare(sw *Coordinates, ne *Coordinates) (*Square, error) {
	if sw.Latitude > ne.Latitude {
		return nil, errors.New("Latitudes are in wrong order")
	}
	if sw.Longitude > ne.Longitude {
		return nil, errors.New("Longitudes are in wrong order")
	}
	return &Square{
		SouthWest: sw,
		NorthEast: ne,
	}, nil
}

// Circle defines an area bounded by a coordinate
// and a radius.
// Tags are used to map from the JSON response.
type Circle struct {
	Centre *Coordinates `json:"centre"`
	Radius float64      `json:"radius"`
}

// NewCircle returns a circular region
func NewCircle(centre *Coordinates, radius float64) (*Circle, error) {
	if radius <= 0.0 {
		return nil, errors.New("Radius is not positive")
	}
	return &Circle{
		Centre: centre,
		Radius: radius,
	}, nil
}

// PolyGon defines an area bounded by a slice of coordinates.
// Tags are used to map from the JSON response.
type PolyGon struct {
	Path []*Coordinates `json:"path"`
}

// NewPolyGon returns a polygon region
func NewPolyGon(path []*Coordinates) *PolyGon {
	return &PolyGon{
		Path: path,
	}
}
