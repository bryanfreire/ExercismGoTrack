package space

type Planet string

// Age calculated from planet
func Age(s float64, p Planet) (r float64) {
	var secsInEarth float64

	secsInEarth = 31557600
	switch p {
	case "Mercury":
		r = s / (secsInEarth * 0.2408467)
	case "Venus":
		r = s / (secsInEarth * 0.61519726)
	case "Earth":
		r = s / (secsInEarth * 1.0)
	case "Mars":
		r = s / (secsInEarth * 1.8808158)
	case "Jupiter":
		r = s / (secsInEarth * 11.862615)
	case "Saturn":
		r = s / (secsInEarth * 29.447498)
	case "Uranus":
		r = s / (secsInEarth * 84.016846)
	case "Neptune":
		r = s / (secsInEarth * 164.79132)
	}

	return
}
