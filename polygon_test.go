package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPolygonArea(t *testing.T) {
	// simplest 1/1 polygon
	assert.Equal(
		t,
		GetPolygonArea(Vertices{
			NewVector2(0, 0),
			NewVector2(0, 1),
			NewVector2(1, 1),
			NewVector2(1, 0),
		}),
		1.0,
	)
	// quater of simplest 1/1 polygon
	assert.Equal(
		t,
		GetPolygonArea(Vertices{
			NewVector2(0, 0),
			NewVector2(0, 0.5),
			NewVector2(0.5, 0.5),
			NewVector2(0.5, 0),
		}),
		0.25,
	)
	// simple 10/10 polygon
	assert.Equal(
		t,
		GetPolygonArea(Vertices{
			NewVector2(0, 0),
			NewVector2(0, 10),
			NewVector2(10, 10),
			NewVector2(10, 0),
		}),
		100.0,
	)
	// triangle as a half of simple 10/10 polygon
	assert.Equal(
		t,
		GetPolygonArea(Vertices{
			NewVector2(0, 0),
			NewVector2(0, 10),
			NewVector2(10, 0),
		}),
		50.0,
	)
	// more irregular yet still convex polygon
	assert.Equal(
		t,
		GetPolygonArea(Vertices{
			NewVector2(0, 0),
			NewVector2(0, 110),
			NewVector2(10, 10),
			NewVector2(10, 0),
		}),
		600.0,
	)
	// non-convex polygon
	assert.Equal(
		t,
		GetPolygonArea(Vertices{
			NewVector2(0, 0),
			NewVector2(0, 110),
			NewVector2(10, 10),
			NewVector2(20, 110),
			NewVector2(20, 0),
		}),
		1200.0,
	)
}
