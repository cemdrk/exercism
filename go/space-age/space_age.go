// Package space provides utility functions about space
package space

// Planet is type for planet
type Planet string

var periods = []struct {
	planet Planet
	period float64
}{
	{"Mercury", 0.2408467},
	{"Venus", 0.61519726},
	{"Earth", 1},
	{"Mars", 1.8808158},
	{"Jupiter", 11.862615},
	{"Saturn", 29.447498},
	{"Uranus", 84.016846},
	{"Neptune", 164.79132},
}

const earthYearInSec float64 = 31557600

// Age calculates age in Earth-years
func Age(ageInSec float64, planet Planet) float64 {
	var ageOnPlanet float64

	for _, p := range periods {
		if p.planet == planet {
			ageOnEarth := ageInSec / earthYearInSec
			ageOnPlanet = ageOnEarth / p.period
			break
		}
	}

	return ageOnPlanet
}
