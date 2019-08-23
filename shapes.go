// -*- mode: go; coding: utf-8 -*-
// vi: set syntax=go :
// cSpell.language:en-GB
// cSpell:disable

package what3words

import (
	"bytes"
	"fmt"

	"github.com/juju/errors"
)

// Box defines an area bounded by two coordinates.
// South West and North East points of the Box
// define its extent.
// // Tags are used to map from the JSON response.
type Box struct {
	SouthWest *Coordinates `json:"southwest"`
	NorthEast *Coordinates `json:"northeast"`
}

// NewBox returns a Box region
func NewBox(sw *Coordinates, ne *Coordinates) (*Box, error) {
	latSpan := ne.Latitude - sw.Latitude
	lonSpan := ne.Longitude - sw.Longitude
	if latSpan < 0 {
		return nil, errors.New("Latitude span is < 0")
	}
	if lonSpan < 0 {
		return nil, errors.New("Longitudes span is < 0")
	}
	if latSpan > 180 {
		return nil, errors.New("Latitude span is > 180")
	}
	if lonSpan > 360 {
		return nil, errors.New("Longitude span is > 360")
	}
	return &Box{
		SouthWest: sw,
		NorthEast: ne,
	}, nil
}

// String returns a string suitable for a URL parameter.
func (box *Box) String() string {
	return fmt.Sprintf("%.13f,%.13f,%.13f,%.13f",
		box.SouthWest.Latitude, box.SouthWest.Longitude,
		box.NorthEast.Latitude, box.NorthEast.Longitude)
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

// String returns a string suitable for a URL parameter.
func (circle *Circle) String() string {
	return fmt.Sprintf("%.13f,%.13f,%.13f",
		circle.Centre.Latitude, circle.Centre.Longitude, circle.Radius)
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

// String returns a string suitable for a URL parameter.
func (polygon *PolyGon) String() string {
	var buffer bytes.Buffer
	var number = len(polygon.Path)
	for index, coord := range polygon.Path {
		buffer.WriteString(coord.String())
		if index < number-1 {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}
