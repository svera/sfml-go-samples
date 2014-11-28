package "animation"

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
    m_frames = append(m_frames, rect)
}

func (this *Animation) SetSpriteSheet(texture sf.Texture) {
    m_texture = texture;
}

func (this *Animation) GetSpriteSheet() sf.Texture {
    return m_texture
}

func (this *Animation) GetSize() uint8 {
    return len(m_frames)
}

func (this *Animation) GetFrame(n uint8) sf.IntRect {
    return m_frames[n]
}
