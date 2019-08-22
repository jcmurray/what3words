// -*- mode: go; coding: utf-8 -*-
// vi: set syntax=go :
// cSpell.language:en-GB
// cSpell:disable

package what3words

import (
	"reflect"
	"testing"
)

func TestNewCoordinates(t *testing.T) {
	type args struct {
		lat float64
		lon float64
	}
	tests := []struct {
		name    string
		args    args
		want    *Coordinates
		wantErr bool
	}{
		{"1 Good Lat Lon", args{lat: 21.9292, lon: -10.282}, &Coordinates{Latitude: 21.9292, Longitude: -10.282}, false},
		{"2 Bad Lat", args{lat: -121.9292, lon: -10.282}, nil, true},
		{"3 Bad Lat", args{lat: 130.9292, lon: -10.282}, nil, true},
		{"4 Bad Lon", args{lat: 21.9292, lon: -102.282}, nil, true},
		{"5 Bad Lat", args{lat: -130.9292, lon: 10.282}, nil, true},
		{"6 Bad Lon", args{lat: -21.9292, lon: 102.282}, nil, true},
		{"7 Bad Lat lon", args{lat: 221.9292, lon: 102.282}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCoordinates(tt.args.lat, tt.args.lon)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCoordinates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCoordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoordinates_SetLat(t *testing.T) {
	type fields struct {
		Latitude  float64
		Longitude float64
	}
	type args struct {
		lat float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"1 Good Lat", fields{Latitude: 21.9292, Longitude: 0}, args{lat: 21.9292}, false},
		{"2 Bad Lat", fields{Latitude: 0, Longitude: 0}, args{lat: 121.9292}, true},
		{"3 Bad Lat", fields{Latitude: 0, Longitude: 0}, args{lat: -121.9292}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coord := &Coordinates{
				Latitude:  tt.fields.Latitude,
				Longitude: tt.fields.Longitude,
			}
			if err := coord.SetLat(tt.args.lat); (err != nil) != tt.wantErr {
				t.Errorf("Coordinates.SetLat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCoordinates_SetLon(t *testing.T) {
	type fields struct {
		Latitude  float64
		Longitude float64
	}
	type args struct {
		lon float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"1 Good Lon", fields{Latitude: 0, Longitude: 21.9292}, args{lon: 21.9292}, false},
		{"2 Bad Lon", fields{Latitude: 0, Longitude: 0}, args{lon: 121.9292}, true},
		{"3 Bad Lon", fields{Latitude: 0, Longitude: 0}, args{lon: -121.9292}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coord := &Coordinates{
				Latitude:  tt.fields.Latitude,
				Longitude: tt.fields.Longitude,
			}
			if err := coord.SetLon(tt.args.lon); (err != nil) != tt.wantErr {
				t.Errorf("Coordinates.SetLon() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
