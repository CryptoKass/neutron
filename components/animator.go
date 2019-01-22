package components

import (
	"github.com/CryptoKass/neutron"
	"github.com/CryptoKass/neutron/core"
)

type Animator struct {
	Animations     map[string]*Animation
	SpriteRenderer *SpriteRenderer
	Current        string
	Paused         bool
}

type Animation struct {
	Frames  []*core.Texture
	Speed   float32
	timer   float32
	current int
}

func (animator *Animator) OnDraw(rend *core.Renderer) error {
	if !animator.Paused {
		return nil
	}

	anim := animator.Animations[animator.Current]

	anim.timer += anim.Speed
	if anim.timer >= 60 {
		anim.timer = 0 //reset the timer
		anim.current++ // goto next frame
		if anim.current > len(anim.Frames) {
			anim.current = 0 //wrap frame
		}
		animator.SpriteRenderer.Texture = anim.Frames[anim.current]
	}

	return nil
}

func (animator *Animator) OnUpdate(rend *core.Renderer) error {
	return nil
}

func (animator *Animator) OnCollision(other *neutron.Element) error {
	return nil
}

func (animator *Animator) AddAnimation(name string, speed float32, frames ...*core.Texture) *Animation {
	anim := Animation{}
	anim.Speed = speed
	for _, frame := range frames {
		anim.Frames = append(anim.Frames, frame)
	}

	//add animation to the animator
	if animator.Animations == nil {
		animator.Animations = make(map[string]*Animation)
	}

	if animator.Current == "" {
		animator.Current = name
	}

	animator.Animations[name] = &anim
	return &anim
}

func (animator *Animator) Play(name string) {
	animator.Current = name
	anim := animator.Animations[animator.Current]
	animator.SpriteRenderer.Texture = anim.Frames[0]
}
