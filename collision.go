package machinery

import (
	"math"
)

const (
	overlapCorrectionPercent = 0.8
	overlapCorrectionSlop    = 0.01
)

type Collision struct {
	Depth         float64
	Normal        Vector2
	ContactPoints []Vector2
}

func ResolveCollision(rbA, rbB *RigidBody, c Collision) {
	for _, cp := range c.ContactPoints {
		if cp.IsEqual(Vector2Zero) {
			continue
		}

		armA := cp.NewSubtracted(rbA.Collider.Center)
		rotationalVelocityA := NewVector2(-rbA.AngularVelocity*armA.Y, rbA.AngularVelocity*armA.X)
		closingVelocityA := rbA.Velocity.NewAdded(rotationalVelocityA)

		armB := cp.NewSubtracted(rbB.Collider.Center)
		rotationalVelocityB := NewVector2(-rbB.AngularVelocity*armB.Y, rbB.AngularVelocity*armB.X)
		closingVelocityB := rbB.Velocity.NewAdded(rotationalVelocityB)

		impulseAugmentationA := armA.Cross(c.Normal)
		impulseAugmentationA = impulseAugmentationA * rbA.inverseMomentOfInertia * impulseAugmentationA

		impulseAugmentationB := armB.Cross(c.Normal)
		impulseAugmentationB = impulseAugmentationB * rbB.inverseMomentOfInertia * impulseAugmentationB

		relativeVelocity := closingVelocityA.NewSubtracted(closingVelocityB)
		separatingVelocity := relativeVelocity.Dot(c.Normal)

		newSeparatingVelocity := math.Max(0, -math.Min(rbA.COR, rbB.COR)*separatingVelocity)
		separatingVelocityDifference := newSeparatingVelocity - separatingVelocity

		impulseMagnitude := separatingVelocityDifference / (rbA.inverseMass + rbB.inverseMass + impulseAugmentationA + impulseAugmentationB)

		impulse := c.Normal.NewScaled(impulseMagnitude)

		rbA.Velocity.Add(impulse.NewScaled(rbA.inverseMass))
		rbB.Velocity.Add(impulse.NewScaled(-rbB.inverseMass))

		rbA.AngularVelocity += armA.Cross(impulse) * rbA.inverseMomentOfInertia
		rbB.AngularVelocity -= armB.Cross(impulse) * rbB.inverseMomentOfInertia
	}

	if c.Depth-overlapCorrectionSlop > 0 {
		correction := c.Normal.NewScaled(
			math.Max(c.Depth-overlapCorrectionSlop, 10e-4) / (rbA.inverseMass + rbB.inverseMass) * overlapCorrectionPercent,
		)
		rbA.Collider.Position.Add(correction.NewScaled(rbA.inverseMass))
		rbB.Collider.Position.Add(correction.NewScaled(-rbB.inverseMass))
	}
}

func GetContactPoints(rbA, rbB RigidBody, c Collision) []Vector2 {
	var (
		cp1, cp2      Vector2 = Vector2Zero, Vector2Zero
		minDistanceSq         = math.MaxFloat64
	)

	for _, rbAPoint := range rbA.Collider.Points {
		for i, rbBPoint := range rbB.Collider.Points {
			cp := GetPointOnSegmentClosestToPoint(rbAPoint, rbBPoint, rbB.Collider.Points.GetNextVertex(i))

			distance := rbAPoint.Distance(cp)
			distanceSq := distance * distance

			if IsEqualish(distanceSq, minDistanceSq) {
				if !cp.IsEqualish(cp1) && !cp.IsEqualish(cp2) {
					cp2 = cp
				}
			} else if distanceSq < minDistanceSq {
				minDistanceSq = distanceSq
				cp1 = cp
			}
		}
	}

	for _, rbBPoint := range rbB.Collider.Points {
		for i, rbAPoint := range rbA.Collider.Points {
			cp := GetPointOnSegmentClosestToPoint(rbBPoint, rbAPoint, rbA.Collider.Points.GetNextVertex(i))

			distance := rbBPoint.Distance(cp)
			distanceSq := distance * distance

			if IsEqualish(distanceSq, minDistanceSq) {
				if !cp.IsEqualish(cp1) && !cp.IsEqualish(cp2) {
					cp2 = cp
				}
			} else if distanceSq < minDistanceSq {
				minDistanceSq = distanceSq
				cp1 = cp
			}
		}
	}

	if cp2 == Vector2Zero {
		return []Vector2{cp1}
	}
	return []Vector2{cp1, cp2}
}

func CheckPolyPolyCollision(polyA, polyB Polygon) (bool, Collision) {
	response := Collision{Depth: math.MaxFloat64, Normal: Vector2Zero}

	for i := 0; i < len(polyA.Points); i++ {
		if ok, depth := isSeparatingAxis(polyA, polyB, polyA.Normals[i]); ok {
			return false, response
		} else if depth < response.Depth {
			response.Depth = depth
			response.Normal = polyA.Normals[i]
		}
	}

	for i := 0; i < len(polyB.Points); i++ {
		if ok, depth := isSeparatingAxis(polyA, polyB, polyB.Normals[i]); ok {
			return false, response
		} else if depth < response.Depth {
			response.Depth = depth
			response.Normal = polyB.Normals[i]
		}
	}

	centerDelta := polyB.Center.NewSubtracted(polyA.Center)
	if centerDelta.Dot(response.Normal) < 0 {
		response.Normal = response.Normal.NewScaled(-1)
	}

	return true, response
}

func isSeparatingAxis(polyA, polyB Polygon, axis Vector2) (bool, float64) {
	var (
		minA, maxA = projectOnAxis(polyA.Points, axis)
		minB, maxB = projectOnAxis(polyB.Points, axis)
	)

	if maxA < minB+1e-9 || maxB < minA+1e-9 {
		return true, 0
	}

	var (
		overlapA = maxB - minA
		overlapB = maxA - minB
	)

	if overlapB < overlapA {
		return false, overlapB
	}

	return false, overlapA
}

func projectOnAxis(points Vertices, axis Vector2) (min, max float64) {
	min = math.MaxFloat64
	max = -math.MaxFloat64
	for _, p := range points {
		dot := axis.Dot(p)
		if dot < min {
			min = dot
		}
		if dot > max {
			max = dot
		}
	}
	return min, max
}
