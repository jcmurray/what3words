// -*- mode: go; coding: utf-8 -*-
// vi: set syntax=go :
// cSpell.language:en-GB
// cSpell:disable

package what3words

import (
	"fmt"
	"net/url"

	"github.com/juju/errors"
	geojson "github.com/paulmach/go.geojson"
)

const (
	baseURL         = "https://api.what3words.com"
	formatJSON      = "json"
	formatGeoJSON   = "geojson"
	defaultLanguage = "en"
	defaultVersion  = "v3"
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
		format:   formatJSON,
		language: defaultLanguage,
		version:  defaultVersion,
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

// SetBaseURL set the base URL for a Geocoder.
func (geocoder *Geocoder) SetBaseURL(urlString string) error {
	newURL, err := url.Parse(urlString)
	if err != nil {
		return err
	}
	geocoder.baseURL = newURL
	return nil
}

// Format get the Format from a Geocoder.
func (geocoder *Geocoder) Format() string {
	return geocoder.format
}

// SetFormat set the Format for a Geocoder.
func (geocoder *Geocoder) SetFormat(format string) error {
	if format != formatJSON && format != formatGeoJSON {
		return errors.New(fmt.Sprintf("Format must be one of '%s' or 'g%s'", formatJSON, formatGeoJSON))
	}
	geocoder.format = format
	return nil
}

// SetFormatJSON set format as JSON
func (geocoder *Geocoder) SetFormatJSON() {
	geocoder.SetFormat(formatJSON)
}

// SetFormatGeoJSON set format as GeoJSON
func (geocoder *Geocoder) SetFormatGeoJSON() {
	geocoder.SetFormat(formatGeoJSON)
}

// IsFormatJSON test the Format for a Geocoder.
func (geocoder *Geocoder) IsFormatJSON() bool {
	return (geocoder.Format() == formatJSON)
}

// IsFormatGeoJSON test the Format for a Geocoder.
func (geocoder *Geocoder) IsFormatGeoJSON() bool {
	return (geocoder.Format() == formatGeoJSON)
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
	geocoder.SetFormatJSON()
	resp, err := ConvertTo3waImpl(geocoder, coord)
	return resp.(*ConvertTo3waResponse), err
}

// ConvertTo3waGeoJSON converts a set of coordinates to a 3 word address.
func (geocoder *Geocoder) ConvertTo3waGeoJSON(coord *Coordinates) (*geojson.FeatureCollection, error) {
	geocoder.SetFormatGeoJSON()
	resp, err := ConvertTo3waImpl(geocoder, coord)
	return resp.(*geojson.FeatureCollection), err
}

// ConvertToCoords converts a 3 word address to coordinates.
func (geocoder *Geocoder) ConvertToCoords(words string) (*ConvertToCoordsResponse, error) {
	geocoder.SetFormatJSON()
	resp, err := ConvertToCoordsImpl(geocoder, words)
	return resp.(*ConvertToCoordsResponse), err
}

// ConvertToCoordsGeoJSON converts a 3 word address to coordinates.
func (geocoder *Geocoder) ConvertToCoordsGeoJSON(words string) (*geojson.FeatureCollection, error) {
	geocoder.SetFormatGeoJSON()
	resp, err := ConvertToCoordsImpl(geocoder, words)
	return resp.(*geojson.FeatureCollection), err
}

// AutoSuggest suggests a list of 3 word addresses from an imprecise input.
func (geocoder *Geocoder) AutoSuggest(request *AutoSuggestRequest) (*AutoSuggestResponse, error) {
	if !request.InputTypeIsText() {
		if geocoder.Language() == "" {
			return nil, errors.New("Non-text input types must have language specified")
		}
	}
	geocoder.SetFormatJSON()
	return AutoSuggestImpl(geocoder, request)
}

// GridSection returns a list of lines from a box.
func (geocoder *Geocoder) GridSection(box *Box) (*GridSectionResponse, error) {
	geocoder.SetFormatJSON()
	resp, err := GridSectionImpl(geocoder, box)
	return resp.(*GridSectionResponse), err
}

// GridSectionGeoJSON returns a list of lines from a box.
func (geocoder *Geocoder) GridSectionGeoJSON(box *Box) (*geojson.FeatureCollection, error) {
	geocoder.SetFormatGeoJSON()
	resp, err := GridSectionImpl(geocoder, box)
	return resp.(*geojson.FeatureCollection), err
}

// AvailableLanguages returns a list of available languages.
func (geocoder *Geocoder) AvailableLanguages() (*LanguagesResponse, error) {
	geocoder.SetFormatJSON()
	return AvailableLanguagesImpl(geocoder)
}
