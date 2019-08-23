// -*- mode: go; coding: utf-8 -*-
// vi: set syntax=go :
// cSpell.language:en-GB
// cSpell:disable

package what3words

import (
	"net/url"
)

const (
	baseURL = "https://api.what3words.com"
)

// Geocoder struct with information needed to make HTTP API calls.
type Geocoder struct {
	apiKey   string
	baseURL  *url.URL
	format   string
	language string
	version  string
}

// NewGeocoder return a new Geocoder.
// Default values ensure W3W API V3 is used, json is returned
// and English is the default language.
func NewGeocoder(apiKey string) *Geocoder {
	url, _ := url.Parse(baseURL)
	return &Geocoder{
		apiKey:   apiKey,
		baseURL:  url,
		format:   "json",
		language: "en",
		version:  "v3",
	}
}

// SetAPIKey set the API Key in a Geocoder.
func (geocoder *Geocoder) SetAPIKey(apiKey string) {
	geocoder.apiKey = apiKey
}

// APIKey get the API Key from a Geocoder.
func (geocoder *Geocoder) APIKey() string {
	return geocoder.apiKey
}

// BaseURL get the base URL from a Geocoder.
func (geocoder *Geocoder) BaseURL() *url.URL {
	return geocoder.baseURL
}

// Format get the Format from a Geocoder.
func (geocoder *Geocoder) Format() string {
	return geocoder.format
}

// Language get the language from a Geocoder.
func (geocoder *Geocoder) Language() string {
	return geocoder.language
}

// SetLanguage set the language in a Geocoder.
func (geocoder *Geocoder) SetLanguage(language string) {
	geocoder.language = language
}

// Version get the API version from a Geocoder.
func (geocoder *Geocoder) Version() string {
	return geocoder.version
}

// ConvertTo3wa converts a set of coordinates to a 3 word address.
func (geocoder *Geocoder) ConvertTo3wa(coord *Coordinates) (*ConvertTo3waResponse, error) {
	return ConvertTo3waImpl(geocoder, coord)
}

// ConvertToCoords converts a 3 word address to coordinates.
func (geocoder *Geocoder) ConvertToCoords(words string) (*ConvertToCoordsResponse, error) {
	return ConvertToCoordsImpl(geocoder, words)
}

// AutoSuggest suggests a list of 3 word addresses from an imprecise input.
func (geocoder *Geocoder) AutoSuggest(request *AutoSuggestRequest) (*AutoSuggestResponse, error) {
	return AutoSuggestImpl(geocoder, request)
}

// GridSection returns a list of lines from a box.
func (geocoder *Geocoder) GridSection(box *Box) (*GridSectionResponse, error) {
	return GridSectionImpl(geocoder, box)
}

// AvailableLanguages returns a list of available languages.
func (geocoder *Geocoder) AvailableLanguages() (*LanguagesResponse, error) {
	return AvailableLanguagesImpl(geocoder)
}
