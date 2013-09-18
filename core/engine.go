package core

// A framework to manage and update systems
type Engine struct {
	scene   *Scene
	systems []System
}

// NewEngine creates and initializes a new Engine instance
func NewEngine() *Engine {
	engine := &Engine{}
	engine.scene = &Scene{}
	return engine
}

// Update updates all of the attached systems
func (e *Engine) Update(dt float32) {
	for _, sys := range e.systems {
		sys.Begin(dt)
		// TODO: Fix this :(
		for ent := e.scene.BeginEnt(); ent != nil; ent = e.scene.NextEntity() {
			sys.ProcessEntity(ent, dt)
		}
		sys.End(dt)
	}
}

// AddSystem adds a new system
func (e *Engine) AddSystem(sys System) {
	e.systems = append(e.systems, sys)
}

// Scene returns the engine's scene
func (e *Engine) Scene() *Scene {
	return e.scene
}
