// -*- mode: go; coding: utf-8 -*-
// vi: set syntax=go :
// cSpell.language:en-GB
// cSpell:disable

package what3words

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/juju/errors"
)

const (
	convertTo3waPath         = "convert-to-3wa"
	convertToCoordinatesPath = "convert-to-coordinates"
	autoSuggestPath          = "autosuggest"
)

// ConvertTo3waImpl perform REST API request over HTTP.
func ConvertTo3waImpl(geo *Geocoder, coords *Coordinates) (*ConvertTo3waResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", geo.BaseURL().String(), nil)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to build GET request for ConvertTo3wa()")
	}
	req.Header.Add("Accept", "application/json")
	req.URL.Path = fmt.Sprintf("/%s/%s", geo.Version(), convertTo3waPath)
	q := req.URL.Query()
	q.Add("key", geo.APIKey())
	q.Add("coordinates", fmt.Sprintf("%.6f,%.6f", coords.Latitude, coords.Longitude))
	q.Add("format", geo.Format())
	q.Add("language", geo.Language())
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to complete GET request for ConvertTo3wa()")
	}

	if resp.StatusCode != 200 && resp.StatusCode != 400 && resp.StatusCode != 401 {
		return nil, errors.New(fmt.Sprintf("Status Code %d on GET request for ConvertTo3wa()", resp.StatusCode))
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to read Body in response for ConvertTo3wa()")
	}
	defer resp.Body.Close()

	c3wResp := NewConvertTo3waResponse()
	c3wErr := NewResponseError()

	if resp.StatusCode != 200 {
		err = json.Unmarshal(respBody, c3wErr)
		if err != nil {
			return nil, errors.Annotate(err, "Unable to unmarshal error response for ConvertTo3wa()")
		}
		return nil, errors.New(fmt.Sprintf("%s %s", c3wErr.Code(), c3wErr.Message()))
	}

	err = json.Unmarshal(respBody, c3wResp)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to unmarshal response for ConvertTo3wa()")
	}
	return c3wResp, nil
}

// ConvertToCoordsImpl perform REST API request over HPTTP.
func ConvertToCoordsImpl(geo *Geocoder, words string) (*ConvertToCoordsResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", geo.BaseURL().String(), nil)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to build GET request for ConvertToCoords()")
	}
	req.Header.Add("Accept", "application/json")
	req.URL.Path = fmt.Sprintf("/%s/%s", geo.Version(), convertToCoordinatesPath)
	q := req.URL.Query()
	q.Add("key", geo.APIKey())
	q.Add("words", words)
	q.Add("format", geo.Format())
	q.Add("language", geo.Language())
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to complete GET request for ConvertToCoords()")
	}

	if resp.StatusCode != 200 && resp.StatusCode != 400 && resp.StatusCode != 401 {
		return nil, errors.New(fmt.Sprintf("Status Code %d on GET request for ConvertToCoords()", resp.StatusCode))
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to read Body in response for ConvertToCoords()")
	}
	defer resp.Body.Close()

	coordsResp := NewConvertToCoordsResponse()
	coordsErr := NewResponseError()

	if resp.StatusCode != 200 {
		err = json.Unmarshal(respBody, coordsErr)
		if err != nil {
			return nil, errors.Annotate(err, "Unable to unmarshal error response for ConvertToCoords()")
		}
		return nil, errors.New(fmt.Sprintf("%s %s", coordsErr.Code(), coordsErr.Message()))
	}

	err = json.Unmarshal(respBody, coordsResp)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to unmarshal response for ConvertToCoords()")
	}
	return coordsResp, nil
}

// AutoSuggestImpl perform REST API request over HTTP.
func AutoSuggestImpl(geo *Geocoder, areq *AutoSuggestRequest) (*AutoSuggestResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", geo.BaseURL().String(), nil)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to build GET request for ConvertToCoords()")
	}
	req.Header.Add("Accept", "application/json")
	req.URL.Path = fmt.Sprintf("/%s/%s", geo.Version(), autoSuggestPath)

	q := req.URL.Query()
	q.Add("key", geo.APIKey())
	q.Add("format", geo.Format())
	q.Add("language", geo.Language())

	if areq.Input != "" {
		q.Add("input", areq.Input)
	}
	if areq.NResults > 0 {
		q.Add("n-results", fmt.Sprintf("%d", areq.NResults))
	}
	if areq.Focus != nil {
		q.Add("focus", fmt.Sprintf("%.6f,%.6f", areq.Focus.Latitude, areq.Focus.Longitude))
	}
	if areq.NFocusResults > 0 {
		q.Add("n-focus-results", fmt.Sprintf("%d", areq.NFocusResults))
	}
	if len(areq.ClipToCountry) > 0 {
		q.Add("clip-to-country", strings.Join(areq.ClipToCountry[:], ","))
	}
	if areq.ClipToBoundingBox != nil {
		q.Add("clip-to-bounding-box", fmt.Sprintf("%.6f,%.6f,%.6f,%.6f",
			areq.ClipToBoundingBox.SouthWest.Latitude,
			areq.ClipToBoundingBox.SouthWest.Longitude,
			areq.ClipToBoundingBox.NorthEast.Latitude,
			areq.ClipToBoundingBox.NorthEast.Longitude))
	}
	if areq.ClipToCircle != nil {
		q.Add("clip-to-circle", fmt.Sprintf("%.6f,%.6f,%.6f",
			areq.ClipToCircle.Centre.Latitude, areq.ClipToCircle.Centre.Longitude,
			areq.ClipToCircle.Radius))
	}
	if areq.ClipToPolyGon != nil {
		var buffer bytes.Buffer
		var number = len(areq.ClipToPolyGon.Path)
		for index, coord := range areq.ClipToPolyGon.Path {
			buffer.WriteString(fmt.Sprintf("%.6f,%6f", coord.Latitude, coord.Longitude))
			if index < number-1 {
				buffer.WriteString(",")
			}
		}
		q.Add("clip-to-polygon", buffer.String())
	}
	if areq.InputType != "" {
		q.Add("input-type", areq.InputType)
	}
	q.Add("prefer-land", strconv.FormatBool(areq.PreferLand))
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to complete GET request for AutoSuggest()")
	}

	if resp.StatusCode != 200 && resp.StatusCode != 400 && resp.StatusCode != 401 {
		return nil, errors.New(fmt.Sprintf("Status Code %d on GET request for AutoSuggest()", resp.StatusCode))
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to read Body in response for AutoSuggest()")
	}
	defer resp.Body.Close()

	autoResp := NewAutoSuggestResponse()
	autoErr := NewResponseError()

	if resp.StatusCode != 200 {
		err = json.Unmarshal(respBody, autoErr)
		if err != nil {
			return nil, errors.Annotate(err, "Unable to unmarshal error response for AutoSuggest()")
		}
		return nil, errors.New(fmt.Sprintf("%s %s", autoErr.Code(), autoErr.Message()))
	}

	err = json.Unmarshal(respBody, autoResp)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to unmarshal response for AutoSuggest()")
	}
	return autoResp, nil
}
