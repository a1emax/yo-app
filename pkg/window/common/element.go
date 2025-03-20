package common

import (
	"github.com/a1emax/youngine/scene"
	"github.com/hajimehoshi/ebiten/v2"
)

type Element[T any] interface {
	scene.Element[*ebiten.Image, T]
}

type BaseElement[T, P any] struct {
	scene.BaseElement[*ebiten.Image, T, P]
}
