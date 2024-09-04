package main

import "math"

const broadCollisionRadiusTolerance = 5

type AfterMoveHandler func(self *RigidBody)
type CollisionHandler func(collidedWith *RigidBody)

type RigidBody struct {
	Velocity        Vector2
	AngularVelocity float64

	force Vector2

	IsStatic               bool
	Mass                   float64
	inverseMass            float64
	momentOfInertia        float64
	inverseMomentOfInertia float64
	COR                    float64

	Collider             Polygon
	BroadCollisionRadius float64

	RelatedObject      interface{}
	AfterMoveHandler   AfterMoveHandler
	OnCollisionHandler CollisionHandler
}

func NewRigidBody(
	position Vector2,
	velocity Vector2,
	rotation float64,
	angularVelocity float64,
	isStatic bool,
	dentisity float64,
	cor float64,
	collider Vertices,
	relatedObject interface{},
	afterMoveHandler AfterMoveHandler,
	onCollisionHandler CollisionHandler,
) *RigidBody {
	var (
		numSides              int     = len(collider)
		mass, momentOfInertia float64 = 0, 0
		area                  float64 = GetPolygonArea(collider)
		sideLen               float64 = collider[0].Distance(collider[1])
		broadCollisionRadius  float64 = GetRadiusOfCircumscribedCircleInRegularPolygon(numSides, sideLen) +
			broadCollisionRadiusTolerance
	)

	if isStatic {
		velocity = Vector2Zero
		angularVelocity = 0
	} else {
		mass = area * dentisity
		momentOfInertia = GetCircleMomentOfInertia(mass, broadCollisionRadius)
	}

	return &RigidBody{
		Collider: NewPolygon(
			position,
			rotation,
			collider,
		),
		Velocity:               velocity,
		AngularVelocity:        angularVelocity,
		force:                  Vector2Zero,
		IsStatic:               isStatic,
		Mass:                   mass,
		inverseMass:            1 / mass,
		momentOfInertia:        momentOfInertia,
		inverseMomentOfInertia: 1 / momentOfInertia,
		COR:                    cor,
		BroadCollisionRadius:   broadCollisionRadius,
		RelatedObject:          relatedObject,
		AfterMoveHandler:       afterMoveHandler,
		OnCollisionHandler:     onCollisionHandler,
	}
}

func (r *RigidBody) AddForce(force Vector2) {
	r.force.Add(force)
}

func (rb *RigidBody) Update(w *World) {
	rb.Velocity.Add(rb.force.NewScaled(rb.inverseMass))
	rb.force = Vector2Zero

	rb.Collider.Position.Add(rb.Velocity.NewScaled(GDrag))
	rb.Collider.Rotation = math.Mod(rb.Collider.Rotation+rb.AngularVelocity, PI2)

	rb.AfterMoveHandler(rb)

	rb.Collider.Recalculate()

	for _, rbB := range w.GetRigidBodies() {
		if rb == rbB {
			continue
		}

		if rb.IsStatic && rbB.IsStatic {
			continue
		}

		distanceBetweenCenters := rb.Collider.Center.Distance(rbB.Collider.Center)
		if distanceBetweenCenters > rb.BroadCollisionRadius+rbB.BroadCollisionRadius {
			continue
		}

		if collided, collision := CheckPolyPolyCollision(rb.Collider, rbB.Collider); collided {
			rb.OnCollisionHandler(rbB)

			if rb.IsStatic {
				rbB.Collider.Position.Add(collision.Normal.NewScaled(collision.Depth))
				rbB.Collider.Recalculate()
			} else if rbB.IsStatic {
				rb.Collider.Position.Add(collision.Normal.NewScaled(collision.Depth).NewReversed())
				rb.Collider.Recalculate()
			} else {
				rbB.Collider.Position.Add(collision.Normal.NewScaled(collision.Depth * .5))
				rbB.Collider.Recalculate()
				rb.Collider.Position.Add(collision.Normal.NewScaled(collision.Depth * .5).NewReversed())
				rb.Collider.Recalculate()
			}

			collision.ContactPoints = GetContactPoints(*rb, *rbB, collision)
			w.ContactPoints = append(w.ContactPoints, collision.ContactPoints...)

			// missing collision resolution here
		}
	}
}
