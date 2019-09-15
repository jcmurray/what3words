// -*- mode: go; coding: utf-8 -*-
// vi: set syntax=go :
// cSpell.language:en-GB
// cSpell:disable

package what3words

const (
	inputText  = "text"
	vconHybrid = "vocon-hybrid"
	nmdpAsr    = "nmdp-asr"
)

// AutoSuggestRequest response from REST API.
// Tags are used to map from the JSON response.
type AutoSuggestRequest struct {
	Input             string
	NResults          int
	Focus             *Coordinates
	NFocusResults     int
	ClipToCountry     []string
	ClipToBoundingBox *Box
	ClipToCircle      *Circle
	ClipToPolyGon     *PolyGon
	InputType         string
	PreferLand        bool
}

// AutoSuggestRequestInterface interface definiton
type AutoSuggestRequestInterface interface {
	SetNResults(number int) error
	SetFocus(focus *Coordinates)
	SetNFocusResults(number int) error
	SetClipToCountry(countries []string)
	SetClipToBoundingBox(box *Box)
	SetClipToCircle(circle *Circle)
	SetClipToPolyGon(polygon *PolyGon)
	SetInputType(input string) error
	InputTypeIsText() bool
	SetPreferLand(land bool)
}

// AutoSuggestResponse response from REST API.
// Tags are used to map from the JSON response.
type AutoSuggestResponse struct {
	Suggestions []*Suggestion `json:"suggestions"`
}

// Suggestion response from REST API.
// Tags are used to map from the JSON response.
type Suggestion struct {
	Country           string  `json:"country"`
	NearestPlace      string  `json:"nearestPlace"`
	Words             string  `json:"words"`
	DistanceToFocusKm float64 `json:"distanceToFocusKm"`
	Rank              int     `json:"rank"`
	Language          string  `json:"language"`
}

// Verify these objects implement the necessary interfaces
var _ AutoSuggestRequestInterface = &AutoSuggestRequest{}
