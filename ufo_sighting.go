package goufos

import (
	"fmt"
	"encoding/json"
	"os"
	_ "sort"
)
// Here we see the header-line/column-names of one ufosighting as csv and one actual ufosighting
//
//sightedatyear,sightedatmonth,sightedatday,reportedat,locationcity,locationstate,shape,duration,description
//
//1996,12,25,19961225,Fresno/Stockton (from,to Stockton), CA, light,several minutes,SUMMARY:  Radio show call in single object obsevred as a bright light traveling at a high rate of speed in a SE direction from STCN to FAT light was obsevered to be constant luminance until out of sight.  Multiple witnesses radio station swamped w/ calls.Description same as above observerser backgrounds are wide and varied according to information availiable.This information is is provided by an individual who was listening to the radio at the time of occurance.

// A UfoSighting contains the details of a UfoSighting
type UfoSighting struct {
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

// UfoSighting holds all the information about a specific UFO sighting
//type UfoSighting struct {
//	sightedatyear int
//	sightedatmonth int
//	sightedatday int
//	reportedat int
//	locationcity string
//	locationstate string
//	shape string
//	duration string
//	description string
//}

// UfoSighting implements fmt.Stringer interface for easy printing
func (u *UfoSighting) String() string {
	//return string(u.reportedat)+u.locationcity+u.locationstate+u.shape+u.duration+u.description
	return fmt.Sprintf("%s", u.description)
}

// New returns a UfoSighting variable:
func New(sightedatyear int, sightedatmonth int, sightedatday int, reportedat int, locationcity string, locationstate string, shape string, duration string, description string) (UfoSighting, error) {
	u := UfoSighting{
		Sightedatyear: sightedatyear,
		Sightedatmonth: sightedatmonth,
		Sightedatday: sightedatday,
		Reportedat: reportedat,
		Locationcity: locationcity,
		Locationstate: locationstate,
		Shape: shape,
		Duration: duration,
		Description: description,
	}

	return u, nil
}

// loadUfoSightings reads the file and returns the list of UfoSighting 
func loadUfoSightings(filePath string) ([]UfoSighting, error) {
	// Open the file to get an io.Reader.
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Declare the variable in which the file will be decoded.
	var ufosightings []UfoSighting

	// Decode the file and store the content in the variable ufosightings.
	err = json.NewDecoder(f).Decode(&ufosightings)
	if err != nil {
		return nil, err
	}

	return ufosightings, nil
}

// locationstateCount registers all the ufosighting lcocationstate and their occurrences from the ufosightings
func locationstateCount(ufosightings []UfoSighting) map[string]float64 {
	locationstateCount := make(map[string]float64)

	for _, ufosighting := range ufosightings {
        	// If a ufosighting has two copies, that counts as two.
		locationstateCount[ufosighting.Locationstate]++
	}

	return locationstateCount
}
