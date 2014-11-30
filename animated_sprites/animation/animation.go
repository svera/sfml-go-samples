package animation

import (
    sf "bitbucket.org/krepa098/gosfml2"
)

type Animation struct {
    m_frames []sf.IntRect
    m_texture *sf.Texture
}

func NewAnimation() *Animation {
    a := new(Animation)
    a.m_frames = make([]sf.IntRect, 2)
    return a
}

func (this *Animation) AddFrame(rect sf.IntRect) {
    this.m_frames = append(this.m_frames, rect)
}

func (this *Animation) SetSpriteSheet(texture *sf.Texture) {
    this.m_texture = texture;
}

func (this *Animation) GetSpriteSheet() *sf.Texture {
    return this.m_texture
}

func (this *Animation) GetSize() uint8 {
    return uint8(len(this.m_frames))
}

func (this *Animation) GetFrame(n uint8) sf.IntRect {
    return this.m_frames[n]
}
