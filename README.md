
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
        coords, err := what3words.NewCoordinates(51.520847, -0.195521)
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

func ExampleAutoSuggest() {
    api := what3words.NewGeocoder("[MK_API_KEY]")
        autoreq := what3words.NewAutoSuggestRequest("plan.clips.a")
        coords, err := what3words.NewCoordinates(51.520847, -0.195521)
	autoreq.SetFocus(coords)

	resp, err := api.AutoSuggest(autoreq)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Printf("======================\n")
	for _, suggestion := range resp.Suggestions {
		fmt.Printf("Country: %s\n", suggestion.Country)
		fmt.Printf("Nearest Place: %s\n", suggestion.NearestPlace)
		fmt.Printf("Words: %s\n", suggestion.Words)
		fmt.Printf("Distance to Focus km: %.3f\n", suggestion.DistanceToFocusKm)
		fmt.Printf("Rank: %d\n", suggestion.Rank)
		fmt.Printf("Language: %s\n", suggestion.Language)
		fmt.Printf("======================\n")
	}
        // Output:
        // ======================
        // Country: GB
        // Nearest Place: Brixton Hill, London
        // Words: plan.clips.area
        // Distance to Focus km: 11.000
        // Rank: 1
        // Language: en
        // ======================
        // Country: GB
        // Nearest Place: Borehamwood, Herts.
        // Words: plan.clips.arts
        // Distance to Focus km: 16.000
        // Rank: 2
        // Language: en
        // ======================
        // Country: GB
        // Nearest Place: Wood Green, London
        // Words: plan.slips.cage
        // Distance to Focus km: 13.000
        // Rank: 3
        // Language: en
        // ======================
}

func ExampleGridSelect() {
        coord1, _ := w3w.NewCoordinates(52.207988, 0.116126)
	coord2, _ := w3w.NewCoordinates(52.208867, 0.117540)
	box, _ := w3w.NewBox(coord1, coord2)
	resp, err := api.GridSection(box)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	} //2084411384217
	fmt.Printf("======================\n")
	for _, line := range resp.Lines {
		fmt.Printf("Line Start: %.13f, %.13f\n", line.Start.Latitude, line.Start.Longitude)
		fmt.Printf("Line End  : %.13f, %.13f\n", line.End.Latitude, line.End.Longitude)
		fmt.Printf("======================\n")
        }
        // Output:
        // ======================
        // Line Start: 52.2080099180681, 0.1161260000000
        // Line End  : 52.2080099180681, 0.1175400000000
        // ======================
        // Line Start: 52.2080368693402, 0.1161260000000
        // Line End  : 52.2080368693402, 0.1175400000000
        // ======================
        // Line Start: 52.2080638206123, 0.1161260000000
        // Line End  : 52.2080638206123, 0.1175400000000
        // ...
}
```
