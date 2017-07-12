package geo

import "math"

// CalculateCentroid calculates a centroid based on an array of points
func CalculateCentroid(coords []*Point) *Point {
	if len(coords) == 0 {
		return nil
	}
	x := float64(0)
	y := float64(0)
	z := float64(0)
	for _, coord := range coords {
		// degr to rad
		lat := coord.Lat * math.Pi / 180
		long := coord.Long * math.Pi / 180

		// increment sums
		x += math.Cos(lat) * math.Cos(long)
		y += math.Cos(lat) * math.Sin(long)
		z += math.Sin(lat)
	}

	// calculate averages
	total := float64(len(coords))
	x /= total
	y /= total
	z /= total

	// average xyz to lat/long
	cLong := math.Atan2(y, x)
	cSqrt := math.Sqrt(x*x + y*y)
	cLat := math.Atan2(z, cSqrt)

	// rad back to degr
	lat := cLat * 180 / math.Pi
	long := cLong * 180 / math.Pi

	return NewPoint(long, lat, coords[0].SRID)
}
