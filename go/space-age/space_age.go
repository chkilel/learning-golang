package space

// Planet against which age is calculated
type Planet string

//Age returns the age given in seconds representing the orbital year of the given planet
func Age(ts float64, p Planet) float64 {

	var age float64

	switch p {
	case "Mercury":
		age = float64(ts) / 31557600.0 / 0.2408467
	case "Venus":
		age = float64(ts) / 31557600.0 / 0.61519726
	case "Earth":
		age = float64(ts) / 31557600.0
	case "Mars":
		age = float64(ts) / 31557600.0 / 1.8808158
	case "Jupiter":
		age = float64(ts) / 31557600.0 / 11.862615
	case "Saturn":
		age = float64(ts) / 31557600.0 / 29.447498
	case "Uranus":
		age = float64(ts) / 31557600.0 / 84.016846
	case "Neptune":
		age = float64(ts) / 31557600.0 / 164.79164
	}
	return age
}
