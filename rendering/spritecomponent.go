package rnd

import (
	"github.com/tedsta/fission/core"
	"github.com/tedsta/gosfml"
)

var SpriteComponentType core.TypeBits

type SpriteComponent struct {
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

func SpriteComponentFactory() core.Component {
	return &SpriteComponent{}
}

func NewSpriteComponent(fileName string, frames, framesPerRow int) *SpriteComponent {
	s := &SpriteComponent{texturePath: fileName}
	s.Sprite = sf.NewSprite(sf.NewTextureFromFile(fileName))

	sprSize := s.Sprite.Texture().Size()

	// Calculate frame stuff
	s.EndFrame = frames - 1
	s.FrameStep = 1
	s.FrameDelay = 0.125
	s.LoopAnim = true

	s.frames = frames
	s.framesPerRow = framesPerRow
	s.animClock.Restart()

	s.frameDim.X = sprSize.X / float32(s.framesPerRow)
	s.frameDim.Y = sprSize.Y / float32(s.frames/s.framesPerRow)

	return s
}

func (s *SpriteComponent) Serialize(p *core.OutPacket) {
	p.Write(s.RelPos.X)
	p.Write(s.RelPos.Y)
	p.Write(s.RelRot)

	p.Write(s.CurrentFrame)
	p.Write(s.StartFrame)
	p.Write(s.EndFrame)
	p.Write(s.FrameStep)
	p.Write(s.FrameDelay)
	p.Write(s.LoopAnim)

	p.Write(s.texturePath)
	p.Write(s.frames)
	p.Write(s.framesPerRow)
	p.Write(s.frameDim.X)
	p.Write(s.frameDim.Y)
}

func (s *SpriteComponent) Deserialize(p *core.InPacket) {
	p.Read(&s.RelPos.X)
	p.Read(&s.RelPos.Y)
	p.Read(&s.RelRot)

	p.Read(&s.CurrentFrame)
	p.Read(&s.StartFrame)
	p.Read(&s.EndFrame)
	p.Read(&s.FrameStep)
	p.Read(&s.FrameDelay)
	p.Read(&s.LoopAnim)

	p.Read(&s.texturePath)
	p.Read(&s.frames)
	p.Read(&s.framesPerRow)
	p.Read(&s.frameDim.X)
	p.Read(&s.frameDim.Y)

	s.Sprite = sf.NewSprite(sf.NewTextureFromFile(s.texturePath))
	s.animClock.Restart()
}

func (s *SpriteComponent) TypeBits() core.TypeBits {
	return SpriteComponentType
}

func (s *SpriteComponent) Render(t *sf.RenderTarget, states sf.RenderStates) {
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

	states.Transform.Translate(s.RelPos)

	//states.transform.Combine(transform)
	s.Sprite.Render(t, states)
}
