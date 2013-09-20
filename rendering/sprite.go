package rend

type Sprite struct {
	texture *Texture
	rect    Rect
	verts   [4]Vertex
}

func NewSprite(t *Texture) *Sprite {
	spr := &Sprite{}
	spr.SetTexture(t)

	return spr
}

func (s *Sprite) Render(t *RenderTarget, states RenderStates) {
	states.texture = s.texture
	t.Render(s.verts[:], Quads, states)
}

func (s *Sprite) SetTexture(t *Texture) {
	s.SetRect(Rect{0, 0, t.Size().X, t.Size().Y})
	s.texture = t
}

func (s *Sprite) SetRect(rect Rect) {
	if rect != s.rect {
		s.rect = rect
		s.updatePositions()
		s.updateTexCoords()
	}
}

func (s *Sprite) updatePositions() {
	s.verts[0].Pos = Vector2{}
	s.verts[1].Pos = Vector2{0, s.rect.H}
	s.verts[2].Pos = Vector2{s.rect.W, s.rect.H}
	s.verts[3].Pos = Vector2{s.rect.W, 0}
}

func (s *Sprite) updateTexCoords() {
	left := s.rect.Left
	right := left + s.rect.W
	top := s.rect.Top
	bottom := top + s.rect.H

	s.verts[0].TexCoords = Vector2{left, top}
	s.verts[1].TexCoords = Vector2{left, bottom}
	s.verts[2].TexCoords = Vector2{right, bottom}
	s.verts[3].TexCoords = Vector2{right, top}
}
