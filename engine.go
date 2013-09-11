package fission

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
func (this *Engine) Update(dt float32) {
	for _, sys := range this.systems {
		sys.Begin(dt)
		// TODO: Fix this :(
		for ent := this.scene.BeginEnt(); ent != nil; ent = this.scene.NextEntity() {
			sys.ProcessEntity(ent, dt)
		}
		sys.End(dt)
	}
}

// AddSystem adds a new system
func (this *Engine) AddSystem(s System) {
	this.systems = append(this.systems, s)
}

// Scene returns the engine's scene
func (this *Engine) Scene() *Scene {
	return this.scene
}
