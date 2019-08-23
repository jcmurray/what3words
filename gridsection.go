// -*- mode: go; coding: utf-8 -*-
// vi: set syntax=go :
// cSpell.language:en-GB
// cSpell:disable

package what3words

// GridSectionResponse response from REST API.
// Tags are used to map from the JSON response.
type GridSectionResponse struct {
	Lines []*Line `json:"lines"`
}

// Line object.
// Tags are used to map from the JSON response.
type Line struct {
	Start *Coordinates `json:"start"`
	End   *Coordinates `json:"end"`
}

// NewGridSectionResponse instantiate a GridSectionResponse object.
func NewGridSectionResponse() *GridSectionResponse {
	return &GridSectionResponse{}
}
