package ufos
// Here we see the header-line/column-names of one ufosighting as csv and one actual ufosighting
//
//sightedatyear,sightedatmonth,sightedatday,reportedat,locationcity,locationstate,shape,duration,description
//
//1996,12,25,19961225,Fresno/Stockton (from,to Stockton), CA, light,several minutes,SUMMARY:  Radio show call in single object obsevred as a bright light traveling at a high rate of speed in a SE direction from STCN to FAT light was obsevered to be constant luminance until out of sight.  Multiple witnesses radio station swamped w/ calls.Description same as above observerser backgrounds are wide and varied according to information availiable.This information is is provided by an individual who was listening to the radio at the time of occurance.

// UfoSighting holds all the information about a specific UFO sighting
type UfoSighting struct {
	sightedatyear int
	sightedatmonth int
	sightedatday int
	reportedat int
	locationcity string
	locationstate string
	shape string
	duration string
	description string
}

// UfoSighting implements fmt.Stringer interface for easy printing
func (u *UfoSighting) String() {
	return string(u.reportedat)+u.locationcity+u.locationstate+u.shape+u.duration+u.description
}
