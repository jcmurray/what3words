// -*- mode: go; coding: utf-8 -*-
// vi: set syntax=go :
// cSpell.language:en-GB
// cSpell:disable

package what3words

// ShapeInterface common operations on shapes
type ShapeInterface interface {
	String() string
}

// Box defines an area bounded by two coordinates.
// South West and North East points of the Box
// define its extent.
// // Tags are used to map from the JSON response.
type Box struct {
	SouthWest *Coordinates `json:"southwest"`
	NorthEast *Coordinates `json:"northeast"`
}

// Circle defines an area bounded by a coordinate
// and a radius.
// Tags are used to map from the JSON response.
type Circle struct {
	Centre *Coordinates `json:"centre"`
	Radius float64      `json:"radius"`
}

// PolyGon defines an area bounded by a slice of coordinates.
// Tags are used to map from the JSON response.
type PolyGon struct {
	Path []*Coordinates `json:"path"`
}

// Verify these objects implement the necessary interfaces
var _ ShapeInterface = &Box{}
var _ ShapeInterface = &Circle{}
var _ ShapeInterface = &PolyGon{}
