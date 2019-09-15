// -*- mode: go; coding: utf-8 -*-
// vi: set syntax=go :
// cSpell.language:en-GB
// cSpell:disable

package what3words

import (
	"net/url"

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

// GeocoderInterface interface to object
type GeocoderInterface interface {
	SetAPIKey(apiKey string)
	APIKey() string
	BaseURL() *url.URL
	SetBaseURL(urlString string) error
	Format() string
	SetFormat(format string) error
	SetFormatJSON()
	SetFormatGeoJSON()
	IsFormatJSON() bool
	IsFormatGeoJSON() bool
	Language() string
	SetLanguage(language string)
	Version() string
	ConvertTo3wa(coord *Coordinates) (*ConvertTo3waResponse, error)
	ConvertTo3waGeoJSON(coord *Coordinates) (*geojson.FeatureCollection, error)
	ConvertToCoords(words string) (*ConvertToCoordsResponse, error)
	ConvertToCoordsGeoJSON(words string) (*geojson.FeatureCollection, error)
	AutoSuggest(request *AutoSuggestRequest) (*AutoSuggestResponse, error)
	GridSection(box *Box) (*GridSectionResponse, error)
	GridSectionGeoJSON(box *Box) (*geojson.FeatureCollection, error)
	AvailableLanguages() (*LanguagesResponse, error)
}

// Verify this object implement the necessary interfaces
var _ GeocoderInterface = &Geocoder{}
