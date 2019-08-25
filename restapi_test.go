// -*- mode: go; coding: utf-8 -*-
// vi: set syntax=go :
// cSpell.language:en-GB
// cSpell:disable

package what3words

import (
	"net/http"
	"net/http/httptest"
	"testing"

	geojson "github.com/paulmach/go.geojson"
)

func TestConvertTo3waImpl_StatusOK(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"country": "GB",
			"square": {
				"southwest": {
					"lng": -0.195543,
					"lat": 51.520833
				},
				"northeast": {
					"lng": -0.195499,
					"lat": 51.52086
				}
			},
			"nearestPlace": "Bayswater, London",
			"coordinates": {
				"lng": -0.195521,
				"lat": 51.520847
			},
			"words": "filled.count.soap",
			"language": "en",
			"map": "https://w3w.co/filled.count.soap"
			}`))
	}))
	defer ts.Close()

	api := NewGeocoder("XXXXXXXX")
	err := api.SetBaseURL(ts.URL)
	if err != nil {
		t.Error(err)
		return
	}
	coords, err := NewCoordinates(51.520847, -0.195521)
	if err != nil {
		t.Error(err)
		return
	}
	c := make(chan *ConvertTo3waResponse)

	go retrieveConvertTo3waResponse(c, api, coords)

	resp := <-c
	if resp.Country != "GB" {
		t.Error("Country code mismatch")
	}
	if resp.Words != "filled.count.soap" {
		t.Error("Words mismatch")
	}
	if resp.Map != "https://w3w.co/filled.count.soap" {
		t.Error("Map mismatch")
	}
	if resp.NearestPlace != "Bayswater, London" {
		t.Error("Nearest place mismatch")
	}
	if resp.Language != "en" {
		t.Error("Language mismatch")
	}
	if resp.Coordinates.Latitude != 51.520847 || resp.Coordinates.Longitude != -0.195521 {
		t.Error("Coordinates mismatch")
	}
	if resp.Square.SouthWest.Latitude != 51.520833 ||
		resp.Square.SouthWest.Longitude != -0.195543 ||
		resp.Square.NorthEast.Latitude != 51.52086 ||
		resp.Square.NorthEast.Longitude != -0.195499 {
		t.Error("Square coordinates mismatch")
	}
}

func TestConvertTo3waImpl_GeoJSONStatusOK(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"features": [
				{
					"bbox": [
						-0.195543,
						51.520833,
						-0.195499,
						51.52086
					],
					"geometry": {
						"coordinates": [
							-0.195521,
							51.520847
						],
						"type": "Point"
					},
					"type": "Feature",
					"properties": {
						"country": "GB",
						"nearestPlace": "Bayswater, London",
						"words": "filled.count.soap",
						"language": "en",
						"map": "https://w3w.co/filled.count.soap"
					}
				}
			],
			"type": "FeatureCollection"
		}`))
	}))
	defer ts.Close()

	api := NewGeocoder("XXXXXXXX")
	api.SetFormatGeoJSON()
	err := api.SetBaseURL(ts.URL)
	if err != nil {
		t.Error(err)
		return
	}
	coords, err := NewCoordinates(51.520847, -0.195521)
	if err != nil {
		t.Error(err)
		return
	}
	c := make(chan *geojson.FeatureCollection)

	go retrieveConvertTo3waResponseGeoJSON(c, api, coords)

	fc := <-c

	if fc.Features[0].Properties["country"] != "GB" {
		t.Error("Country code mismatch")
	}
	if fc.Features[0].Properties["words"] != "filled.count.soap" {
		t.Error("Words mismatch")
	}
	if fc.Features[0].Properties["map"] != "https://w3w.co/filled.count.soap" {
		t.Error("Map mismatch")
	}
	if fc.Features[0].Properties["nearestPlace"] != "Bayswater, London" {
		t.Error("Nearest place mismatch")
	}
	if fc.Features[0].Properties["language"] != "en" {
		t.Error("Language mismatch")
	}
	if fc.Features[0].Geometry.Point[1] != 51.520847 || fc.Features[0].Geometry.Point[0] != -0.195521 {
		t.Error("Coordinates mismatch")
	}
	if fc.Features[0].BoundingBox[0] != -0.195543 ||
		fc.Features[0].BoundingBox[1] != 51.520833 ||
		fc.Features[0].BoundingBox[2] != -0.195499 ||
		fc.Features[0].BoundingBox[3] != 51.52086 {
		t.Error("Square coordinates mismatch")
	}
}
func TestConvertTo3waImpl_StatusBadRequest(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{
			"error": {
				"code": "InvalidKey",
				"message": "Authentication failed; invalid API key"
			}
		}`))
	}))
	defer ts.Close()

	api := NewGeocoder("BADKEY")
	err := api.SetBaseURL(ts.URL)
	if err != nil {
		t.Error(err)
		return
	}
	coords, err := NewCoordinates(51.520847, -0.195521)
	if err != nil {
		t.Error(err)
		return
	}
	c := make(chan error)
	go retrieveConvertToErrorResponse(c, api, coords)

	respErr := <-c

	if respErr.Error() != "InvalidKey Authentication failed; invalid API key" {
		t.Error(respErr)
	}
}

func TestConvertToCoordsImpl_StatusOK(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"country": "GB",
			"square": {
				"southwest": {
					"lng": -0.195543,
					"lat": 51.520833
				},
				"northeast": {
					"lng": -0.195499,
					"lat": 51.52086
				}
			},
			"nearestPlace": "Bayswater, London",
			"coordinates": {
				"lng": -0.195521,
				"lat": 51.520847
			},
			"words": "filled.count.soap",
			"language": "en",
			"map": "https://w3w.co/filled.count.soap"
			}`))
	}))
	defer ts.Close()

	api := NewGeocoder("XXXXXXXX")
	err := api.SetBaseURL(ts.URL)
	if err != nil {
		t.Error(err)
		return
	}
	c := make(chan *ConvertToCoordsResponse)

	go retrieveConvertToCoordsResponse(c, api, "filled.count.soap")

	resp := <-c
	if resp.Country != "GB" {
		t.Error("Country code mismatch")
	}
	if resp.Words != "filled.count.soap" {
		t.Error("Words mismatch")
	}
	if resp.Map != "https://w3w.co/filled.count.soap" {
		t.Error("Map mismatch")
	}
	if resp.NearestPlace != "Bayswater, London" {
		t.Error("Nearest place mismatch")
	}
	if resp.Language != "en" {
		t.Error("Language mismatch")
	}
	if resp.Coordinates.Latitude != 51.520847 || resp.Coordinates.Longitude != -0.195521 {
		t.Error("Coordinates mismatch")
	}
	if resp.Square.SouthWest.Latitude != 51.520833 ||
		resp.Square.SouthWest.Longitude != -0.195543 ||
		resp.Square.NorthEast.Latitude != 51.52086 ||
		resp.Square.NorthEast.Longitude != -0.195499 {
		t.Error("Square coordinates mismatch")
	}
}

func TestConvertToCoordinatesImpl_GeoJSONStatusOK(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"features": [
				{
					"bbox": [
						-0.195543,
						51.520833,
						-0.195499,
						51.52086
					],
					"geometry": {
						"coordinates": [
							-0.195521,
							51.520847
						],
						"type": "Point"
					},
					"type": "Feature",
					"properties": {
						"country": "GB",
						"nearestPlace": "Bayswater, London",
						"words": "filled.count.soap",
						"language": "en",
						"map": "https://w3w.co/filled.count.soap"
					}
				}
			],
			"type": "FeatureCollection"
		}`))
	}))
	defer ts.Close()

	api := NewGeocoder("XXXXXXXX")
	api.SetFormatGeoJSON()
	err := api.SetBaseURL(ts.URL)
	if err != nil {
		t.Error(err)
		return
	}
	c := make(chan *geojson.FeatureCollection)

	go retrieveConvertToCoordinatesResponseGeoJSON(c, api, "filled.count.soap")

	fc := <-c

	if fc.Features[0].Properties["country"] != "GB" {
		t.Error("Country code mismatch")
	}
	if fc.Features[0].Properties["words"] != "filled.count.soap" {
		t.Error("Words mismatch")
	}
	if fc.Features[0].Properties["map"] != "https://w3w.co/filled.count.soap" {
		t.Error("Map mismatch")
	}
	if fc.Features[0].Properties["nearestPlace"] != "Bayswater, London" {
		t.Error("Nearest place mismatch")
	}
	if fc.Features[0].Properties["language"] != "en" {
		t.Error("Language mismatch")
	}
	if fc.Features[0].Geometry.Point[1] != 51.520847 || fc.Features[0].Geometry.Point[0] != -0.195521 {
		t.Error("Coordinates mismatch")
	}
	if fc.Features[0].BoundingBox[0] != -0.195543 ||
		fc.Features[0].BoundingBox[1] != 51.520833 ||
		fc.Features[0].BoundingBox[2] != -0.195499 ||
		fc.Features[0].BoundingBox[3] != 51.52086 {
		t.Error("Square coordinates mismatch")
	}
}

func retrieveConvertTo3waResponse(c chan<- *ConvertTo3waResponse, api *Geocoder, coords *Coordinates) {
	resp, _ := ConvertTo3waImpl(api, coords)
	c <- resp.(*ConvertTo3waResponse)
}

func retrieveConvertTo3waResponseGeoJSON(c chan<- *geojson.FeatureCollection, api *Geocoder, coords *Coordinates) {
	resp, _ := ConvertTo3waImpl(api, coords)
	c <- resp.(*geojson.FeatureCollection)
}

func retrieveConvertToErrorResponse(c chan<- error, api *Geocoder, coords *Coordinates) {
	_, err := ConvertTo3waImpl(api, coords)
	c <- err
}

func retrieveConvertToCoordsResponse(c chan<- *ConvertToCoordsResponse, api *Geocoder, words string) {
	resp, _ := ConvertToCoordsImpl(api, words)
	c <- resp.(*ConvertToCoordsResponse)
}

func retrieveConvertToCoordinatesResponseGeoJSON(c chan<- *geojson.FeatureCollection, api *Geocoder, words string) {
	resp, _ := ConvertToCoordsImpl(api, words)
	c <- resp.(*geojson.FeatureCollection)
}
