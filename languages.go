// -*- mode: go; coding: utf-8 -*-
// vi: set syntax=go :
// cSpell.language:en-GB
// cSpell:disable

package what3words

// LanguagesResponse response from REST API.
// Tags are used to map from the JSON response.
type LanguagesResponse struct {
	Languages []Language `json:"languages"`
}

// Language response from REST API
// Tags are used to map from the JSON response.
type Language struct {
	NativeName string `json:"nativeName"`
	Code       string `json:"code"`
	Name       string `json:"name"`
}

// NewLanguagesResponse instantiate a LanguagesResponse object.
func NewLanguagesResponse() *LanguagesResponse {
	return &LanguagesResponse{}
}
