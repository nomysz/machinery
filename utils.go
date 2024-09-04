package main

// given point and segment defined by two other points
// provide point on segment that is closest to given point
func GetPointOnSegmentClosestToPoint(p, segmentPA, segmentPB Vector2) Vector2 {
	pAToP := p.NewSubtracted(segmentPA)
	pAToPB := segmentPB.NewSubtracted(segmentPA)

	projection := pAToP.Dot(pAToPB) / pAToPB.Dot(pAToPB)

	var contactPoint Vector2

	if projection <= 0 {
		contactPoint = segmentPA
	} else if projection >= 1 {
		contactPoint = segmentPB
	} else {
		contactPoint = segmentPA.NewAdded(pAToPB.NewScaled(projection))
	}

	return contactPoint
}

func GetCircleMomentOfInertia(mass, radius float64) float64 {
	return .5 * mass * radius * radius
}

func appendIfMissing[T comparable](slice []T, i T) []T {
	for _, item := range slice {
		if item == i {
			return slice
		}
	}
	return append(slice, i)
}
