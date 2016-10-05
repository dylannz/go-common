package geo

// BoundingBox contains two sets of lat longs which represent a cube area
type BoundingBox struct {
	NW   Point
	SE   Point
	SRID int
}

// New returns a new BoundingBox from the given lat/long pairs
func NewBoundingBox(nwLat float64, nwLon float64, seLat float64, seLon float64, srid int) (*BoundingBox, error) {
	var bbBox BoundingBox
	var err error
	bbBox.SRID = srid

	bbBox.NW, err = NewPoint(nwLon, nwLat, srid)
	if err != nil {
		return nil, err
	}
	bbBox.SE, err = NewPoint(seLon, seLat, srid)
	if err != nil {
		return nil, err
	}

	return &bbBox, nil
}
