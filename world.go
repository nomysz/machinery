package main

const (
	GGravity = 9.81
	GDrag    = .98
)

type PhysicalWorld struct {
	rigidBodies   []*RigidBody
	ContactPoints []Vector2
}

func NewPhysicalWorld() *PhysicalWorld {
	return &PhysicalWorld{}
}

func (w *PhysicalWorld) AddRigidBody(rb *RigidBody) {
	w.rigidBodies = appendIfMissing(w.rigidBodies, rb)
}

func (w *PhysicalWorld) DeleteRigidBody(rb *RigidBody) {
	for i, rigidBody := range w.rigidBodies {
		if rb == rigidBody {
			copy(w.rigidBodies[i:], w.rigidBodies[i+1:])
			w.rigidBodies[len(w.rigidBodies)-1] = nil
			w.rigidBodies = w.rigidBodies[:len(w.rigidBodies)-1]
			return
		}
	}
}

func (w *PhysicalWorld) GetRigidBodies() []*RigidBody {
	return w.rigidBodies
}

func (w *PhysicalWorld) DeleteRigidBodies() {
	w.rigidBodies = []*RigidBody{}
}

func (w *PhysicalWorld) Update(debugMessages *[]string) {
	w.ContactPoints = []Vector2{}
	for _, rb := range w.rigidBodies {
		rb.Update(w)
	}
}
