package loremipsum

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/scene"
	"github.com/a1emax/youngine/scene/element/flexbox"
	"github.com/a1emax/youngine/scene/element/padding"

	"yo-app/pkg/global/consts"
	"yo-app/pkg/window/common"
	"yo-app/pkg/window/loremipsum/back"
	"yo-app/pkg/window/loremipsum/text"
)

type LoremIpsum[T any] interface {
	common.Element[T]
}

func New[T any](traitFunc scene.TraitFunc[T]) LoremIpsum[T] {
	return flexbox.New(traitFunc, flexbox.Props{
		AlignItems:     flexbox.AlignCenter,
		Direction:      flexbox.DirectionColumn,
		JustifyContent: flexbox.JustifyCenter,
	}.Func,
		padding.New(flexbox.Trait{}.Func, text.New(padding.Trait{
			Bottom: basic.SetOpt(consts.AboutPaddingVertical),
			Left:   basic.SetOpt(consts.AboutPaddingHorizontal),
			Right:  basic.SetOpt(consts.AboutPaddingHorizontal),
			Top:    basic.SetOpt(consts.AboutPaddingVertical),
		}.Func)),

		padding.New(flexbox.Trait{}.Func, back.New(padding.Trait{
			Bottom: basic.SetOpt(consts.BackPaddingBottom),
		}.Func)),
	)
}
