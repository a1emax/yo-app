package toggledebug

import (
	"github.com/a1emax/youngine/scene"
	"github.com/a1emax/youngine/scene/element/x/button"

	"yo-app/pkg/global/assets"
	"yo-app/pkg/global/consts"
	"yo-app/pkg/global/tools"
	"yo-app/pkg/window/common"
)

type ToggleDebug[T any] interface {
	common.Element[T]
}

func New[T any](traitFunc scene.TraitFunc[T]) ToggleDebug[T] {
	return button.New(traitFunc, button.Config{
		Clock: tools.Clock,
		Input: tools.Input,
	}, func(button.Props) button.Props {
		var text string
		if tools.Account.Debug() {
			text = assets.Texts.StopDebug
		} else {
			text = assets.Texts.StartDebug
		}

		return button.Props{
			Attrs: scene.Attrs{}.
				FixHeight(consts.TextButtonHeight).
				FixWidth(consts.TextButtonWidth),

			FontFace:         assets.FontFaces.LargeText,
			PressedColor:     assets.Colors.ButtonPressed,
			PressedTextColor: assets.Colors.Info,
			PrimaryColor:     assets.Colors.ButtonPrimary,
			PrimaryTextColor: assets.Colors.Info,
			Text:             text,

			OnClick: func(event button.ClickEvent) {
				tools.Account.ToggleDebug()
			},
		}
	})
}
