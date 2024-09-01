package main

type Vertices []Vector2

type Polygon struct {
	InitialPoints Vertices
	Position      Vector2
	Rotation      float64 // radians

	Center  Vector2
	Points  Vertices
	Edges   []Vector2
	Normals []Vector2
}

// rotation unit is radians
func NewPolygon(
	position Vector2,
	rotation float64,
	initialPoints Vertices,
) Polygon {
	pLen := len(initialPoints)
	points := make(Vertices, pLen)
	edges := make([]Vector2, pLen)
	normals := make([]Vector2, pLen)

	polygon := Polygon{
		Position:      position,
		Rotation:      rotation,
		InitialPoints: initialPoints,
		Points:        points,
		Edges:         edges,
		Normals:       normals,
	}
	polygon.Recalculate()

	return polygon
}

func (v Vertices) NewTranslated(translation Vector2) Vertices {
	translated := make([]Vector2, len(v))

	for i, vertex := range v {
		translated[i] = Vector2{
			X: vertex.X + translation.X,
			Y: vertex.Y + translation.Y,
		}
	}

	return translated
}

func (v Vertices) NewScaled(scale Vector2) Vertices {
	scaled := make([]Vector2, len(v))

	for i, vertex := range v {
		scaled[i] = Vector2{
			X: vertex.X * scale.X,
			Y: vertex.Y * scale.Y,
		}
	}

	return scaled
}

func (p Polygon) Copy() Polygon {
	return p
}

// first one is next last one
func (v Vertices) GetNextVertex(i int) Vector2 {
	if i+1 == len(v) {
		return v[0]
	}
	return v[i+1]
}

func (p *Polygon) Recalculate() {
	numSides := len(p.InitialPoints)

	for i := 0; i < numSides; i++ {
		p.Points[i] = NewVector2(
			p.InitialPoints[i].X+p.Position.X,
			p.InitialPoints[i].Y+p.Position.Y,
		)
	}

	var sumX, sumY float64 = 0, 0
	for _, v := range p.Points {
		sumX += v.X
		sumY += v.Y
	}
	p.Center = NewVector2(sumX/float64(numSides), sumY/float64(numSides))

	for i := 0; i < numSides; i++ {
		if p.Rotation != 0 {
			p.Points[i].RotateAround(p.Center, p.Rotation)
		}
	}

	for i := 0; i < numSides; i++ {
		p.Edges[i] = p.Points.GetNextVertex(i).NewSubtracted(p.Points[i])
		p.Normals[i] = p.Edges[i].NewPerpendicular().NewNormalized()
	}
}
