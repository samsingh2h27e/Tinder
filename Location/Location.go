package location

import "math"

type Location struct {
	Lat  float64
	Long float64
}

func (l *Location) DistanceInKM(l1 Location) float64 {
	const earthRadiusKm = 6371.0

	// Convert degrees to radians
	lat1 := l1.Lat * math.Pi / 180
	lat2 := l.Lat * math.Pi / 180
	lng1 := l1.Long * math.Pi / 180
	lng2 := l.Long * math.Pi / 180

	// Differences
	dLat := lat2 - lat1
	dLng := lng2 - lng1

	// Haversine formula
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dLng/2)*math.Sin(dLng/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadiusKm * c
}