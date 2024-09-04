package main

import (
	"math"
)

type Collision struct {
	Depth         float64
	Normal        Vector2
	ContactPoints []Vector2
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
		} else {
			if depth < response.Depth {
				response.Depth = depth
				response.Normal = polyA.Normals[i]
			}
		}
	}

	for i := 0; i < len(polyB.Points); i++ {
		if ok, depth := isSeparatingAxis(polyA, polyB, polyB.Normals[i]); ok {
			return false, response
		} else {
			if depth < response.Depth {
				response.Depth = depth
				response.Normal = polyB.Normals[i]
			}
		}
	}

	direction := polyB.Center.Copy().NewSubtracted(polyB.Center)
	if direction.Dot(response.Normal) < 0 {
		response.Normal.Reverse()
	}

	return true, response
}

func isSeparatingAxis(polyA, polyB Polygon, axis Vector2) (bool, float64) {
	minA, maxA := projectOnAxis(polyA.Points, axis)
	minB, maxB := projectOnAxis(polyB.Points, axis)
	minDepth := math.Min(maxB-minA, maxA-minB)
	return maxA < minB || maxB < minA, minDepth
}

func projectOnAxis(points Vertices, axis Vector2) (min, max float64) {
	min = math.MaxFloat64
	max = -math.MaxFloat64
	for i := 0; i < len(points); i++ {
		dot := axis.Dot(points[i])
		if dot < min {
			min = dot
		}
		if dot > max {
			max = dot
		}
	}
	return min, max
}
