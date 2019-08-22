// -*- mode: go; coding: utf-8 -*-
// vi: set syntax=go :
// cSpell.language:en-GB
// cSpell:disable

package what3words

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGeocoder_SetAPIKey(t *testing.T) {
	type fields struct {
		APIKey string
	}
	type args struct {
		apiKey string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"Set XXX", fields{APIKey: "XXX"}, args{apiKey: "XXX"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			geocoder := &Geocoder{
				apiKey: tt.fields.APIKey,
			}
			geocoder.SetAPIKey(tt.args.apiKey)
		})
	}
}

func TestGeocoder_APIKey(t *testing.T) {
	type fields struct {
		APIKey string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Get XXX", fields{APIKey: "XXX"}, "XXX"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			geocoder := &Geocoder{
				apiKey: tt.fields.APIKey,
			}
			if got := geocoder.APIKey(); got != tt.want {
				t.Errorf("Geocoder.GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGeocoder(t *testing.T) {
	url, _ := url.Parse(baseURL)
	type args struct {
		apiKey string
	}
	tests := []struct {
		name string
		args args
		want *Geocoder
	}{
		{"New APIKey: XXX", args{apiKey: "XXX"},
			&Geocoder{apiKey: "XXX", baseURL: url, format: "json", language: "en", version: "v3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGeocoder(tt.args.apiKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGeocoder() = %v, want %v", got, tt.want)
			}
		})
	}
}
