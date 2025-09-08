package main

import (
	"encoding/json"
	"os"
	_ "sort"
)

// A Ufosighting contains the details of a Ufosighting
type Ufosighting struct {
	Sightedatyear int `json:"sightedatyear"`
	Sightedatmonth int `json:"sightedatmonth"`
	Sightedatday int `json:"sightedatday"`
	Reportedat int `json:"reportedat"`
	Locationcity string `json:"locationcity"`
	Locationstate string `json:"locationstate"`
	Shape string `json:"shape"`
	Duration string `json:"duration"`
	Description string `json:"description"`
}

// loadUfosightings reads the file and returns the list of Ufosighting 
func loadUfosightings(filePath string) ([]Ufosighting, error) {
	// Open the file to get an io.Reader.
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Declare the variable in which the file will be decoded.
	var ufosightings []Ufosighting

	// Decode the file and store the content in the variable ufosightings.
	err = json.NewDecoder(f).Decode(&ufosightings)
	if err != nil {
		return nil, err
	}

	return ufosightings, nil
}

// locationstateCount registers all the ufosighting lcocationstate and their occurrences from the ufosightings
func locationstateCount(ufosightings []Ufosighting) map[string]float64 {
	locationstateCount := make(map[string]float64)

	for _, ufosighting := range ufosightings {
        	// If a ufosighting has two copies, that counts as two.
		locationstateCount[ufosighting.Locationstate]++
	}

	return locationstateCount
}
