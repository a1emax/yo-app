package text

import (
	"github.com/a1emax/youngine/scene"
	"github.com/a1emax/youngine/scene/element/x/textscroller"

	"yo-app/pkg/global/assets"
	"yo-app/pkg/global/tools"
	"yo-app/pkg/window/common"
)

type Text[T any] interface {
	common.Element[T]
}

func New[T any](traitFunc scene.TraitFunc[T]) Text[T] {
	return textscroller.New(traitFunc, textscroller.Config{
		Clock: tools.Clock,
		Input: tools.Input,
	}, textscroller.Props{
		FontFace:  assets.FontFaces.MediumText,
		Text:      assets.Texts.LoremIpsum,
		TextColor: assets.Colors.Info,
	}.Func)
}
