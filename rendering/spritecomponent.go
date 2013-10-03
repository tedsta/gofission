package rnd

import (
	"github.com/tedsta/fission/core"
	"github.com/tedsta/gosfml"
)

type spriteComponent struct {
	Sprite *sf.Sprite
	RelPos sf.Vector2 // The offset of the sprite in relation to the entity
	RelRot float32    // The rotation of the sprite in relation to the entity

	CurrentFrame int     // The current frame number in the animation
	StartFrame   int     // The start frame of the frame loop
	EndFrame     int     // The end frame of the frame loop
	FrameStep    int     // The number of frames to step
	FrameDelay   float32 // The time (in seconds) between frames
	LoopAnim     bool    // Whether or not to loop the animation

	texturePath  string     // The file path to this sprite's texture
	frames       int        // The total number of frames in the animation
	framesPerRow int        // The number of frames per row in the sprite sheet
	animClock    sf.Clock   // The SFML clock
	frameDim     sf.Vector2 // The sprite animation's frame dimensions
}

func NewSpriteComponent(fileName string, frames, framesPerRow int) *RenderComponent {
	s := &spriteComponent{texturePath: fileName}
	s.Sprite = sf.NewSprite(sf.NewTextureFromFile(fileName))

	sprSize := s.Sprite.Texture().Size()

	// Calculate frame stuff
	s.EndFrame = frames - 1
	s.FrameStep = 1
	s.FrameDelay = 0.1
	s.LoopAnim = true

	s.frames = frames
	s.framesPerRow = framesPerRow
	s.animClock.Restart()

	s.frameDim.X = sprSize.X / float32(s.framesPerRow)
	s.frameDim.Y = sprSize.Y / float32(s.frames/s.framesPerRow)

	return NewRenderComponent(s.Serialize, s.Deserialize, s.Render)
}

func (s *spriteComponent) Serialize(p *core.Packet) {
}

func (s *spriteComponent) Deserialize(p *core.Packet) {
}

func (s *spriteComponent) Render(t *sf.RenderTarget, states sf.RenderStates) {
	if float32(s.animClock.ElapsedTime().Seconds()) >= s.FrameDelay &&
		(s.LoopAnim || s.CurrentFrame != s.EndFrame) {

		s.animClock.Restart()
		s.CurrentFrame += s.FrameStep

		if s.CurrentFrame > s.EndFrame && s.FrameStep >= 1 && s.LoopAnim {
			s.CurrentFrame = s.StartFrame
		} else if s.CurrentFrame < s.StartFrame && s.FrameStep <= -1 && s.LoopAnim {
			s.CurrentFrame = s.EndFrame
		}
	}

	// Calculate frame position
	frameX := float32(s.CurrentFrame%s.framesPerRow) * s.frameDim.X
	frameY := float32(s.CurrentFrame/s.framesPerRow) * s.frameDim.Y

	s.Sprite.SetTextureRect(sf.Rect{frameX, frameY, s.frameDim.X, s.frameDim.Y})

	//states.transform.Combine(transform)
	s.Sprite.Render(t, states)
}
