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

// ErrorDetails response from REST API.
// Tags are used to map from the JSON response.
type ErrorDetails struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// NewResponseError instantiate a ResponseError object.
func NewResponseError() *ResponseError {
	return &ResponseError{}
}

// Code of error message.
func (respErr *ResponseError) Code() string {
	return respErr.Error.Code
}

// Message of error message.
func (respErr *ResponseError) Message() string {
	return respErr.Error.Message
}
