// -*- mode: go; coding: utf-8 -*-
// vi: set syntax=go :
// cSpell.language:en-GB
// cSpell:disable

package what3words

import (
	"fmt"

	"github.com/juju/errors"
)

// NewAutoSuggestRequest instantiate a AutoSuggestRequest object.
func NewAutoSuggestRequest(input string) *AutoSuggestRequest {
	return &AutoSuggestRequest{
		Input:         input,
		NResults:      3,
		NFocusResults: 3,
		PreferLand:    false,
	}
}

// NewAutoSuggestResponse instantiate a AutoSuggestResponse object.
func NewAutoSuggestResponse() *AutoSuggestResponse {
	return &AutoSuggestResponse{}
}

// SetNResults sets the Number of results parameter.
func (ar *AutoSuggestRequest) SetNResults(number int) error {
	if number <= 0 || number > 100 {
		return errors.New("Number of results must be in range (0, 100]")
	}
	ar.NResults = number
	return nil
}

// SetFocus sets the focus point.
func (ar *AutoSuggestRequest) SetFocus(focus *Coordinates) {
	ar.Focus = focus
}

// SetNFocusResults sets the Number of focus results parameter.
func (ar *AutoSuggestRequest) SetNFocusResults(number int) error {
	if number <= 0 || number > ar.NResults {
		return errors.New("Number of results must be in range (0, Number-results]")
	}
	ar.NFocusResults = number
	return nil
}

// SetClipToCountry sets the list of countries.
func (ar *AutoSuggestRequest) SetClipToCountry(countries []string) {
	ar.ClipToCountry = countries
}

// SetClipToBoundingBox sets the cliping box.
func (ar *AutoSuggestRequest) SetClipToBoundingBox(box *Box) {
	ar.ClipToBoundingBox = box
}

// SetClipToCircle sets the clipping circle.
func (ar *AutoSuggestRequest) SetClipToCircle(circle *Circle) {
	ar.ClipToCircle = circle
}

// SetClipToPolyGon sets the clipping circle.
func (ar *AutoSuggestRequest) SetClipToPolyGon(polygon *PolyGon) {
	ar.ClipToPolyGon = polygon
}

// SetInputType sets the Input type.
func (ar *AutoSuggestRequest) SetInputType(input string) error {
	if input != vconHybrid && input != inputText && input != nmdpAsr {
		return errors.New(fmt.Sprintf("Inout type must be one of '%s' '%s' or '%s'",
			inputText, vconHybrid, nmdpAsr))
	}
	ar.InputType = input
	return nil
}

// InputTypeIsText sets the Input type.
func (ar *AutoSuggestRequest) InputTypeIsText() bool {
	return (ar.InputType == inputText)
}

// SetPreferLand sets the PeferLand type.
func (ar *AutoSuggestRequest) SetPreferLand(land bool) {
	ar.PreferLand = land
}
