
# API for [What3words](https://docs.what3words.com/api/v3/) for Golang

This project is licensed under the terms of the MIT license.

Supports:

- What3words API V3

## Installation

Install:

```shell
go get -u github.com/jcmurray/what3words
```

Import:

```go
import "github.com/jcmurray/what3words"
```

## Quickstart

```go
func ExampleConvertTo3wa() {
        api := what3words.NewGeocoder("[MK_API_KEY]")
        coords, err := w3w.NewCoordinates(51.520847, -0.195521)
        if err != nil {
                panic(err)
        }
        resp, err := api.ConvertTo3wa(coords)
        if err != nil {
                panic(err)
        }
        fmt.Printf("What3Names: %s\n", resp.Words)
        // Output: What3Names: filled.count.soap
}

func ExampleConvertToCoords() {
    api := what3words.NewGeocoder("[MK_API_KEY]")
        resp, err := api.ConvertToCoords("filled.count.soap")
        if err != nil {
            panic(err)
        }
        fmt.Printf("Coords - Lat: %.6f, Lon: %0.6f\n", resp.Coordinates.Latitude, resp.Coordinates.Longitude)}
        // Output: Coords - Lat: 51.520847, Lon: -0.195521
}
```

