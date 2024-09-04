package main

const (
	GGravity = 9.81
	GDrag    = .98
)

type World struct {
	rigidBodies   []*RigidBody
	ContactPoints []Vector2
}

func NewWorld() *World {
	return &World{}
}

func (w *World) AddRigidBody(rb *RigidBody) {
	w.rigidBodies = appendIfMissing(w.rigidBodies, rb)
}

func (w *World) DeleteRigidBody(rb *RigidBody) {
	for i, rigidBody := range w.rigidBodies {
		if rb == rigidBody {
			copy(w.rigidBodies[i:], w.rigidBodies[i+1:])
			w.rigidBodies[len(w.rigidBodies)-1] = nil
			w.rigidBodies = w.rigidBodies[:len(w.rigidBodies)-1]
			return
		}
	}
}

func (w *World) GetRigidBodies() []*RigidBody {
	return w.rigidBodies
}

func (w *World) DeleteRigidBodies() {
	w.rigidBodies = []*RigidBody{}
}

func (w *World) Update() {
	w.ContactPoints = []Vector2{}

	for _, rb := range w.rigidBodies {
		rb.Update(w)
	}
}
