// -*- mode: go; coding: utf-8 -*-
// vi: set syntax=go :
// cSpell.language:en-GB
// cSpell:disable

package what3words

// ConvertTo3waResponse response from REST API.
// Tags are used to map from the JSON response.
type ConvertTo3waResponse struct {
	Response
}

// ConvertToCoordsResponse response from REST API.
// Tags are used to map from the JSON response.
type ConvertToCoordsResponse struct {
	Response
}

// NewConvertTo3waResponse instantiate a ConvertTo3waResponse object.
func NewConvertTo3waResponse() *ConvertTo3waResponse {
	return &ConvertTo3waResponse{}
}

// NewConvertToCoordsResponse instantiate a ConvertToCoordsResponse object.
func NewConvertToCoordsResponse() *ConvertToCoordsResponse {
	return &ConvertToCoordsResponse{}
}
