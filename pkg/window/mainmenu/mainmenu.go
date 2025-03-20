package mainmenu

import (
	"github.com/a1emax/youngine/scene"
	"github.com/a1emax/youngine/scene/element/flexbox"
	"github.com/a1emax/youngine/scene/element/overlay"

	"yo-app/pkg/global/consts"
	"yo-app/pkg/window/common"
	"yo-app/pkg/window/mainmenu/toggledebug"
	"yo-app/pkg/window/mainmenu/yo"
)

type MainMenu[T any] interface {
	common.Element[T]
}

func New[T any](traitFunc scene.TraitFunc[T]) MainMenu[T] {
	return overlay.New(traitFunc, overlay.Props{}.Func,
		flexbox.New(overlay.Trait{}.Func, flexbox.Props{
			Attrs: scene.Attrs{}.
				FixHeight(consts.MainMenuButtonsHeight),

			AlignItems:     flexbox.AlignCenter,
			Direction:      flexbox.DirectionColumn,
			JustifyContent: flexbox.JustifySpaceBetween,
		}.Func,
			yo.New(flexbox.Trait{}.Func),
			toggledebug.New(flexbox.Trait{}.Func),
		),
	)
}
