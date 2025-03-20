package yo

import (
	"github.com/a1emax/youngine/scene"
	"github.com/a1emax/youngine/scene/element/x/button"

	"yo-app/pkg/global/assets"
	"yo-app/pkg/global/consts"
	"yo-app/pkg/global/tools"
	"yo-app/pkg/global/vars"
	"yo-app/pkg/window/common"
)

type Yo[T any] interface {
	common.Element[T]
}

func New[T any](traitFunc scene.TraitFunc[T]) Yo[T] {
	return button.New(traitFunc, button.Config{
		Clock: tools.Clock,
		Input: tools.Input,
	}, button.Props{
		Attrs: scene.Attrs{}.
			FixHeight(consts.TextButtonHeight).
			FixWidth(consts.TextButtonWidth),

		FontFace:         assets.FontFaces.LargeText,
		PressedColor:     assets.Colors.ButtonPressed,
		PressedTextColor: assets.Colors.Info,
		PrimaryColor:     assets.Colors.ButtonPrimary,
		PrimaryTextColor: assets.Colors.Info,
		Text:             assets.Texts.Yo,

		OnClick: func(event button.ClickEvent) {
			vars.Window.Page = vars.WindowPageLoremIpsum
		},
	}.Func)
}
