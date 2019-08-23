// -*- mode: go; coding: utf-8 -*-
// vi: set syntax=go :
// cSpell.language:en-GB
// cSpell:disable

package what3words

// Response generic response from REST API.
// Tags are used to map from the JSON response.
type Response struct {
	Country      string       `json:"country"`
	Square       *Box         `json:"square"`
	NearestPlace string       `json:"nearestPlace"`
	Coordinates  *Coordinates `json:"coordinates"`
	Words        string       `json:"words"`
	Language     string       `json:"language"`
	Map          string       `json:"map"`
}
