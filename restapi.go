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

	"github.com/juju/errors"
)

const (
	convertTo3waPath     = "convert-to-3wa"
	convertToCoordinates = "convert-to-coordinates"
)

// ConvertTo3waImpl perform REST API request over HTTP.
func ConvertTo3waImpl(geo *Geocoder, coords *Coordinates) (*ConvertTo3waResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", geo.BaseURL().String(), nil)
	if err != nil {
		return nil, errors.Annotate(err, "Unable to build GET request for ConvertTo3wa()")
	}
	req.Header.Add("Accept", "application/json")
	q := req.URL.Query()
	q.Add("key", geo.APIKey())
	q.Add("coordinates", fmt.Sprintf("%.6f,%.6f", coords.Latitude, coords.Longitude))
	q.Add("format", geo.Format())
	q.Add("language", geo.Language())
	req.URL.RawQuery = q.Encode()
	req.URL.Path = fmt.Sprintf("/%s/%s", geo.Version(), convertTo3waPath)

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
	q := req.URL.Query()
	q.Add("key", geo.APIKey())
	q.Add("words", words)
	q.Add("format", geo.Format())
	q.Add("language", geo.Language())
	req.URL.RawQuery = q.Encode()
	req.URL.Path = fmt.Sprintf("/%s/%s", geo.Version(), convertToCoordinates)

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
