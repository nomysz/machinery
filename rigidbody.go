package main

type RigidBody struct {
	Velocity        Vector2
	AngularVelocity float64

	IsStatic bool

	Mass        float64
	InverseMass float64

	COR float64

	Collider Polygon
}

func NewRigidBody(
	position Vector2,
	velocity Vector2,
	rotation float64,
	angularVelocity float64,
	isStatic bool,
	mass float64,
	cor float64,
	collider Vertices,
) *RigidBody {
	if isStatic {
		velocity = Vector2Zero
		angularVelocity = 0
		mass = 1_000_000_000
	}
	return &RigidBody{
		Collider: NewPolygon(
			position,
			rotation,
			collider,
		),
		Velocity:        velocity,
		AngularVelocity: angularVelocity,
		IsStatic:        isStatic,
		Mass:            mass,
		InverseMass:     1 / mass,
		COR:             cor,
	}
}
