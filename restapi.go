// -*- mode: go; coding: utf-8 -*-
// vi: set syntax=go :
// cSpell.language:en-GB
// cSpell:disable

package what3words

import (
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
	gridSectionPath          = "grid-section"
	availableLanguagesPath   = "available-languages"
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
	q.Add("coordinates", coords.String())
	q.Add("format", geo.Format())
	q.Add("language", geo.Language())
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to complete GET request for ConvertTo3wa()")
	}

	if httpError(resp) {
		return nil, errors.New(fmt.Sprintf("Status '%s' on GET request for ConvertTo3wa()", resp.Status))
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to read Body in response for ConvertTo3wa()")
	}
	defer resp.Body.Close()

	c3wResp := NewConvertTo3waResponse()
	c3wErr := NewResponseError()

	if appError(resp) {
		err = json.Unmarshal(respBody, c3wErr)
		if err != nil {
			return nil, errors.Annotate(err, fmt.Sprintf("Status '%s' response for ConvertTo3wa()", resp.Status))
		}
		return nil, c3wErr.AsError()
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

	if httpError(resp) {
		return nil, errors.New(fmt.Sprintf("Status Code %d on GET request for ConvertToCoords()", resp.StatusCode))
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to read Body in response for ConvertToCoords()")
	}
	defer resp.Body.Close()

	coordsResp := NewConvertToCoordsResponse()
	coordsErr := NewResponseError()

	if appError(resp) {
		err = json.Unmarshal(respBody, coordsErr)
		if err != nil {
			return nil, errors.Annotate(err, fmt.Sprintf("Status '%s' response for ConvertToCoords()", resp.Status))
		}
		return nil, coordsErr.AsError()
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
		q.Add("focus", fmt.Sprintf("%.13f,%.13f", areq.Focus.Latitude, areq.Focus.Longitude))
	}
	if areq.NFocusResults > 0 {
		q.Add("n-focus-results", fmt.Sprintf("%d", areq.NFocusResults))
	}
	if len(areq.ClipToCountry) > 0 {
		q.Add("clip-to-country", strings.Join(areq.ClipToCountry[:], ","))
	}
	if areq.ClipToBoundingBox != nil {
		q.Add("clip-to-bounding-box", areq.ClipToBoundingBox.String())
	}
	if areq.ClipToCircle != nil {
		q.Add("clip-to-circle", areq.ClipToCircle.String())
	}
	if areq.ClipToPolyGon != nil {
		q.Add("clip-to-polygon", areq.ClipToPolyGon.String())
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

	if httpError(resp) {
		return nil, errors.New(fmt.Sprintf("Status Code %d on GET request for AutoSuggest()", resp.StatusCode))
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to read Body in response for AutoSuggest()")
	}
	defer resp.Body.Close()

	autoResp := NewAutoSuggestResponse()
	autoErr := NewResponseError()

	if appError(resp) {
		err = json.Unmarshal(respBody, autoErr)
		if err != nil {
			return nil, errors.Annotate(err, fmt.Sprintf("Status '%s' response for AutoSuggest()", resp.Status))
		}
		return nil, autoErr.AsError()
	}

	err = json.Unmarshal(respBody, autoResp)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to unmarshal response for AutoSuggest()")
	}
	return autoResp, nil
}

// GridSectionImpl perform REST API request over HTTP.
func GridSectionImpl(geo *Geocoder, box *Box) (*GridSectionResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", geo.BaseURL().String(), nil)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to build GET request for GridSelection()")
	}
	req.Header.Add("Accept", "application/json")
	req.URL.Path = fmt.Sprintf("/%s/%s", geo.Version(), gridSectionPath)

	req.Header.Add("Accept", "application/json")
	req.URL.Path = fmt.Sprintf("/%s/%s", geo.Version(), gridSectionPath)
	q := req.URL.Query()
	q.Add("key", geo.APIKey())
	q.Add("format", geo.Format())
	q.Add("language", geo.Language())
	q.Add("bounding-box", box.String())
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to complete GET request for GridSection()")
	}

	if httpError(resp) {
		return nil, errors.New(fmt.Sprintf("Status Code %d on GET request for GridSection()", resp.StatusCode))
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to read Body in response for GridSection()")
	}
	defer resp.Body.Close()

	gridResp := NewGridSectionResponse()
	gridErr := NewResponseError()

	if appError(resp) {
		err = json.Unmarshal(respBody, gridErr)
		if err != nil {
			return nil, errors.Annotate(err, fmt.Sprintf("Status '%s' response for GridSection()", resp.Status))
		}
		return nil, gridErr.AsError()
	}

	err = json.Unmarshal(respBody, gridResp)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to unmarshal response for GridSection()")
	}
	return gridResp, nil
}

// AvailableLanguagesImpl perform REST API request over HTTP.
func AvailableLanguagesImpl(geo *Geocoder) (*LanguagesResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", geo.BaseURL().String(), nil)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to build GET request for AvailableLanguages()")
	}
	req.Header.Add("Accept", "application/json")
	req.URL.Path = fmt.Sprintf("/%s/%s", geo.Version(), availableLanguagesPath)
	q := req.URL.Query()
	q.Add("key", geo.APIKey())
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to complete GET request for AvailableLanguages()")
	}

	if httpError(resp) {
		return nil, errors.New(fmt.Sprintf("Status '%s' on GET request for AvailableLanguages()", resp.Status))
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to read Body in response for AvailableLanguages()")
	}
	defer resp.Body.Close()

	langResp := NewLanguagesResponse()
	langErr := NewResponseError()

	if appError(resp) {
		err = json.Unmarshal(respBody, langErr)
		if err != nil {
			return nil, errors.Annotate(err, fmt.Sprintf("Status '%s' response for AvailableLanguages()", resp.Status))
		}
		return nil, langErr.AsError()
	}

	err = json.Unmarshal(respBody, langResp)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to unmarshal response for AvailableLanguages()")
	}
	return langResp, nil
}

func httpError(resp *http.Response) bool {
	return (resp.StatusCode != 200 && (resp.StatusCode < 400 || 499 < resp.StatusCode))
}

func appError(resp *http.Response) bool {
	return (400 <= resp.StatusCode && resp.StatusCode <= 499)
}
