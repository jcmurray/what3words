// -*- mode: go; coding: utf-8 -*-
// vi: set syntax=go :
// cSpell.language:en-GB
// cSpell:disable

package what3words

// ResponseError response from REST API.
// Tags are used to map from the JSON response.
type ResponseError struct {
	Error ErrorDetails `json:"error"`
}

// ResponseErrorInterface interface definition
type ResponseErrorInterface interface {
	Code() string
	Message() string
	String() string
	AsError() error
}

// ErrorDetails response from REST API.
// Tags are used to map from the JSON response.
type ErrorDetails struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Verify this object implement the necessary interfaces
var _ ResponseErrorInterface = &ResponseError{}
