package main

// given point and segment defined by two other points
// provide point on segment that is closest to given point
func GetPointOnSegmentClosestToPoint(p, segmentP1, segmentP2 Vector2) Vector2 {
	p1ToP := p.NewSubtracted(segmentP1)
	p1ToP2 := segmentP2.NewSubtracted(segmentP1)

	projection := p1ToP.Dot(p1ToP2) / p1ToP2.Dot(p1ToP2)

	var contactPoint Vector2

	if projection <= 0 {
		contactPoint = segmentP1
	} else if projection >= 1 {
		contactPoint = segmentP2
	} else {
		contactPoint = segmentP1.NewAdded(p1ToP2.NewScaled(projection))
	}

	return contactPoint
}
