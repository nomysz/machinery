package machinery

import "math"

type Vector2 struct {
	X, Y float64
}

func NewVector2(X, Y float64) Vector2 {
	return Vector2{X, Y}
}

var (
	Vector2Zero  = NewVector2(0, 0)
	Vector2Right = NewVector2(1, 0)
	Vector2Left  = NewVector2(-1, 0)
	Vector2Up    = NewVector2(0, -1)
	Vector2Down  = NewVector2(0, 1)
)

// angle unit is radians
func NewVector2FromAngle(angle float64, len float64) Vector2 {
	return NewVector2(
		len*math.Cos(angle),
		len*math.Sin(angle),
	)
}

// angle unit is radians
func NewVector2AngledFromSection(v1, v2 Vector2, angle float64) Vector2 {
	dx := v2.X - v1.X
	dy := v2.Y - v1.Y
	distance := math.Sqrt(dx*dx + dy*dy)
	newAngle := math.Atan2(dy, dx) + angle

	return NewVector2(
		v2.X+distance*math.Cos(newAngle),
		v2.Y+distance*math.Sin(newAngle),
	)
}

func (v Vector2) IsEqual(v2 Vector2) bool {
	return v.X == v2.X && v.Y == v2.Y
}

func (v Vector2) IsEqualish(v2 Vector2) bool {
	return v.Distance(v2) < .00001
}

func (v Vector2) Copy() Vector2 {
	return v
}

func (v *Vector2) Normalize() {
	normalized := v.NewNormalized()
	v.X = normalized.X
	v.Y = normalized.Y
}

func (v Vector2) NewNormalized() Vector2 {
	d := v.Len()
	return NewVector2(v.X/d, v.Y/d)
}

func (v *Vector2) Add(addition Vector2) {
	added := v.NewAdded(addition)
	v.X = added.X
	v.Y = added.Y
}

func (v Vector2) NewAdded(another Vector2) Vector2 {
	return NewVector2(v.X+another.X, v.Y+another.Y)
}

func (v *Vector2) Subtract(another Vector2) {
	substracted := v.NewSubtracted(another)
	v.X = substracted.X
	v.Y = substracted.Y
}

func (v Vector2) NewSubtracted(another Vector2) Vector2 {
	return NewVector2(v.X-another.X, v.Y-another.Y)
}

func (v *Vector2) Scale(scalar float64) {
	scaled := v.NewScaled(scalar)
	v.X = scaled.X
	v.Y = scaled.Y
}

func (v Vector2) NewScaled(scalar float64) Vector2 {
	return NewVector2(
		v.X*scalar,
		v.Y*scalar,
	)
}

func (v *Vector2) Reverse() {
	reversed := v.NewReversed()
	v.X = reversed.X
	v.Y = reversed.Y
}

func (v Vector2) NewReversed() Vector2 {
	return NewVector2(-v.X, -v.Y)
}

// angle unit is radians
func (v *Vector2) Rotate(angle float64) {
	rotated := v.NewRotated(angle)
	v.X = rotated.X
	v.Y = rotated.Y
}

// angle unit is radians
func (v Vector2) NewRotated(angle float64) Vector2 {
	return NewVector2(
		v.X*math.Cos(angle)-v.Y*math.Sin(angle),
		v.X*math.Sin(angle)+v.Y*math.Cos(angle),
	)
}

func (v Vector2) NewPerpendicular() Vector2 {
	return NewVector2(v.Y, -v.X)
}

func (v Vector2) Dot(another Vector2) float64 {
	return v.X*another.X + v.Y*another.Y
}

func (v Vector2) Cross(another Vector2) float64 {
	return v.X*another.Y - v.Y*another.X
}

func (v Vector2) Distance(another Vector2) float64 {
	dX := v.X - another.X
	dY := v.Y - another.Y
	return math.Sqrt(dX*dX + dY*dY)
}

func (v Vector2) Len() float64 {
	return math.Sqrt(v.Dot(v))
}

// angle unit is radians
func (v *Vector2) RotateAround(center Vector2, angle float64) {
	rotated := v.NewRotatedAround(center, angle)
	v.X = rotated.X
	v.Y = rotated.Y
}

// angle unit is radians
func (v Vector2) NewRotatedAround(center Vector2, angle float64) Vector2 {
	dX := v.X - center.X
	dY := v.Y - center.Y
	cosTheta := math.Cos(angle)
	sinTheta := math.Sin(angle)
	rotatedX := dX*cosTheta - dY*sinTheta
	rotatedY := dX*sinTheta + dY*cosTheta
	return NewVector2(rotatedX+center.X, rotatedY+center.Y)
}
