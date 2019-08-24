
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
        api := what3words.NewGeocoder("[MK_API_KEY]")
        coord1, _ := w3w.NewCoordinates(52.207988, 0.116126)
	coord2, _ := w3w.NewCoordinates(52.208867, 0.117540)
	box, _ := w3w.NewBox(coord1, coord2)
	resp, err := api.GridSection(box)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
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

func ExampleAvailableLanguages() {
        api := what3words.NewGeocoder("[MK_API_KEY]")
	resp, err := api.AvailableLanguages()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Printf("======================\n")
	for _, language := range resp.Languages {
		fmt.Printf("Name       : %s\n", language.Name)
		fmt.Printf("Code       : %s\n", language.Code)
		fmt.Printf("Native Name: %s\n", language.NativeName)
		fmt.Printf("======================\n")
	}
        // Output:
        // ======================
        // Name       : German
        // Code       : de
        // Native Name: Deutsch
        // ======================
        // Name       : Hindi
        // Code       : hi
        // Native Name: हिन्दी
        // ======================        
}

func ExampleConvertTo3waGeoJSON() {
        api := what3words.NewGeocoder("[MK_API_KEY]")
        coords, err := what3words.NewCoordinates(51.520847, -0.195521)
        if err != nil {
                panic(err)
        }
	resp, err := api.ConvertTo3waGeoJSON(coords)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Printf("What3Names: %s\n", resp.Features[0].Properties["words"])
        // Output: What3Names: filled.count.soap
}

func ExampleConvertToCoordsGeoJSON() {
        api := what3words.NewGeocoder("[MK_API_KEY]")
	resp, err := api.ConvertToCoordsGeoJSON("filled.count.soap")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	coordsGeo := resp.Features[0].Geometry.Point
	fmt.Printf("Coords - Lat: %.6f, Lon: %0.6f\n", coordsGeo[1], coordsGeo[0])
	// Output: Coords - Lat: 51.520847, Lon: -0.195521
}

func ExampleGridSelectGeoJSON() {
        api := what3words.NewGeocoder("[MK_API_KEY]")
        coord1, _ := w3w.NewCoordinates(52.207988, 0.116126)
	coord2, _ := w3w.NewCoordinates(52.208867, 0.117540)
	box, _ := w3w.NewBox(coord1, coord2)
	respGeo, err := api.GridSectionGeoJSON(box)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Printf("======================\n")
	for _, line := range respGeo.Features[0].Geometry.MultiLineString {
		fmt.Printf("Line Start: %.13f, %.13f\n", line[0][1], line[0][0])
		fmt.Printf("Line End  : %.13f, %.13f\n", line[1][1], line[1][0])
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

func ExampleAutoSuggestVoice() {
        api := what3words.NewGeocoder("[MK_API_KEY]")

        // inputVoice is JSON output from Vocon Hybrid vocoder

	inputVoice := `{"_isInGrammar":"yes","_isSpeech":"yes","_hypotheses":[{"_score":342516,"_startRule":"whatthreewordsgrammar#_main_","_conf":6546,"_endTimeMs":6360,"_beginTimeMs":1570,"_lmScore":300,"_items":[{"_type":"terminal","_score":34225,"_orthography":"tend","_conf":6964,"_endTimeMs":2250,"_beginTimeMs":1580},{"_type":"terminal","_score":47670,"_orthography":"artichokes","_conf":7176,"_endTimeMs":3180,"_beginTimeMs":2260},{"_type":"terminal","_score":43800,"_orthography":"poached","_conf":6181,"_endTimeMs":4060,"_beginTimeMs":3220}]},{"_score":342631,"_startRule":"whatthreewordsgrammar#_main_","_conf":6498,"_endTimeMs":6360,"_beginTimeMs":1570,"_lmScore":300,"_items":[{"_type":"terminal","_score":34340,"_orthography":"tent","_conf":6772,"_endTimeMs":2250,"_beginTimeMs":1580},{"_type":"terminal","_score":47670,"_orthography":"artichokes","_conf":7176,"_endTimeMs":3180,"_beginTimeMs":2260},{"_type":"terminal","_score":43800,"_orthography":"poached","_conf":6181,"_endTimeMs":4060,"_beginTimeMs":3220}]},{"_score":342668,"_startRule":"whatthreewordsgrammar#_main_","_conf":6474,"_endTimeMs":6360,"_beginTimeMs":1570,"_lmScore":300,"_items":[{"_type":"terminal","_score":34225,"_orthography":"tend","_conf":6964,"_endTimeMs":2250,"_beginTimeMs":1580},{"_type":"terminal","_score":47670,"_orthography":"artichokes","_conf":7176,"_endTimeMs":3180,"_beginTimeMs":2260},{"_type":"terminal","_score":41696,"_orthography":"perch","_conf":5950,"_endTimeMs":4020,"_beginTimeMs":3220}]},{"_score":342670,"_startRule":"whatthreewordsgrammar#_main_","_conf":6474,"_endTimeMs":6360,"_beginTimeMs":1570,"_lmScore":300,"_items":[{"_type":"terminal","_score":34379,"_orthography":"tinge","_conf":6705,"_endTimeMs":2250,"_beginTimeMs":1580},{"_type":"terminal","_score":47670,"_orthography":"artichokes","_conf":7176,"_endTimeMs":3180,"_beginTimeMs":2260},{"_type":"terminal","_score":43800,"_orthography":"poached","_conf":6181,"_endTimeMs":4060,"_beginTimeMs":3220}]},{"_score":342783,"_startRule":"whatthreewordsgrammar#_main_","_conf":6426,"_endTimeMs":6360,"_beginTimeMs":1570,"_lmScore":300,"_items":[{"_type":"terminal","_score":34340,"_orthography":"tent","_conf":6772,"_endTimeMs":2250,"_beginTimeMs":1580},{"_type":"terminal","_score":47670,"_orthography":"artichokes","_conf":7176,"_endTimeMs":3180,"_beginTimeMs":2260},{"_type":"terminal","_score":41696,"_orthography":"perch","_conf":5950,"_endTimeMs":4020,"_beginTimeMs":3220}]},{"_score":342822,"_startRule":"whatthreewordsgrammar#_main_","_conf":6402,"_endTimeMs":6360,"_beginTimeMs":1570,"_lmScore":300,"_items":[{"_type":"terminal","_score":34379,"_orthography":"tinge","_conf":6705,"_endTimeMs":2250,"_beginTimeMs":1580},{"_type":"terminal","_score":47670,"_orthography":"artichokes","_conf":7176,"_endTimeMs":3180,"_beginTimeMs":2260},{"_type":"terminal","_score":41696,"_orthography":"perch","_conf":5950,"_endTimeMs":4020,"_beginTimeMs":3220}]}],"_resultType":"NBest"}`
	autoreq := w3w.NewAutoSuggestRequest(inputVoice)
	autoreq.SetFocus(coords)
	autoreq.SetInputType("vocon-hybrid")

	resp, err := api.AutoSuggest(autoreq1)
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
        // Country: ZZ
        // Nearest Place: Angaur State, Angaur
        // Words: tend.artichokes.poached
        // Distance to Focus km: 12220.000
        // Rank: 1
        // Language: en
        // ======================
        // Country: ZZ
        // Nearest Place: Berbera, Woqooyi Galbeed
        // Words: tent.artichokes.poached
        // Distance to Focus km: 6112.000
        // Rank: 2
        // Language: en
        // ======================
        // Country: CA
        // Nearest Place: Rouyn-Noranda, Quebec
        // Words: tend.artichokes.perch
        // Distance to Focus km: 5382.000
        // Rank: 3
        // Language: en
        // ======================

}
```
