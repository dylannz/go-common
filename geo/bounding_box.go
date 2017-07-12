package geo

// BoundingBox contains two sets of lat longs which represent a cube area
type BoundingBox struct {
	NW   *Point
	SE   *Point
	SRID int
}

// NewBoundingBox returns a new BoundingBox from the given lat/long pairs
func NewBoundingBox(nwLat float64, nwLon float64, seLat float64, seLon float64, srid int) *BoundingBox {
	return &BoundingBox{
		NW:   NewPoint(nwLon, nwLat, srid),
		SE:   NewPoint(seLon, seLat, srid),
		SRID: srid,
	}
}

// Centroid calculates and returns the centroid of the bounding box
func (bb BoundingBox) Centroid() *Point {
	return CalculateCentroid([]*Point{bb.NW, bb.SE})
}
