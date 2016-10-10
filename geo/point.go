package geo

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
)

// Point represents a lat/long
type Point struct {
	Lat  float64
	Long float64
	SRID int
}

// NewPoint creates and returns a new Point
func NewPoint(x float64, y float64, srid int) (*Point, error) {
	return &Point{
		Long: x,
		Lat:  y,
		SRID: srid,
	}, nil
}

// IsNull returns a boolean indicating whether the point is considered null.
// Currently returns true when the SRID has not been set.
func (p Point) IsNull() bool {
	return 0 == p.SRID
}

// MarshalJSON implements the json.Marshaler interface
func (p Point) MarshalJSON() ([]byte, error) {
	if p.IsNull() {
		return []byte("null"), nil
	}
	return json.Marshal(map[string]float64{
		"lat":  p.Lat,
		"long": p.Long,
	})
}

func (p Point) String() string {
	if p.IsNull() {
		return "NULL"
	}
	point := fmt.Sprintf(
		"ST_GeometryFromText('POINT(%s %s)', %d)",
		strconv.FormatFloat(p.Long, 'f', -1, 64),
		strconv.FormatFloat(p.Lat, 'f', -1, 64),
		p.SRID,
	)
	if p.SRID != WGS84 {
		point = fmt.Sprintf("ST_Transform(%s,%d)", point, WGS84)
	}
	return point
}

// Compare returns true if the SRIDs of the given points are the same, and they
// are located within "diff" distance of each other.
// This is not a safe way to compare a lat/long in general, but for the NZ
// coordinates we are dealing with it should work just fine.
func (p Point) Compare(c Point, diff float64) bool {
	if p.SRID != c.SRID {
		return false
	}
	return math.Abs(p.Lat-c.Lat) < diff && math.Abs(p.Long-c.Long) < diff
}
