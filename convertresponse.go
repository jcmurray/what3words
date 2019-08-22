// -*- mode: go; coding: utf-8 -*-
// vi: set syntax=go :
// cSpell.language:en-GB
// cSpell:disable

package what3words

// ConvertTo3waResponse response from REST API.
// Tags are used to map from the JSON response.
type ConvertTo3waResponse struct {
	Country      string       `json:"country"`
	Square       *Square      `json:"square"`
	NearestPlace string       `json:"nearestPlace"`
	Coordinates  *Coordinates `json:"coordinates"`
	Words        string       `json:"words"`
	Language     string       `json:"language"`
	Map          string       `json:"map"`
}

// ConvertToCoordsResponse response from REST API.
// Tags are used to map from the JSON response.
type ConvertToCoordsResponse struct {
	Country      string       `json:"country"`
	Square       *Square      `json:"square"`
	NearestPlace string       `json:"nearestPlace"`
	Coordinates  *Coordinates `json:"coordinates"`
	Words        string       `json:"words"`
	Language     string       `json:"language"`
	Map          string       `json:"map"`
}

// Square defines an area bounded by two coordinates.
// South West and North East points of the square
// define its extent.
// // Tags are used to map from the JSON response.
type Square struct {
	SouthWest *Coordinates `json:"southwest"`
	NorthEast *Coordinates `json:"northeast"`
}

// NewConvertTo3waResponse instantiate a ConvertTo3waResponse object.
func NewConvertTo3waResponse() *ConvertTo3waResponse {
	return &ConvertTo3waResponse{}
}

// NewConvertToCoordsResponse instantiate a ConvertToCoordsResponse object.
func NewConvertToCoordsResponse() *ConvertToCoordsResponse {
	return &ConvertToCoordsResponse{}
}
