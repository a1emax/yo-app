package window

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/scene/element/flexbox"
	"github.com/a1emax/youngine/scene/element/overlay"
	"github.com/a1emax/youngine/scene/element/pageset"
	"github.com/a1emax/youngine/scene/element/turnoff"
	"github.com/a1emax/youngine/scene/element/x/colorarea"

	"yo-app/pkg/global/assets"
	"yo-app/pkg/global/tools"
	"yo-app/pkg/global/vars"
	"yo-app/pkg/window/common"
	"yo-app/pkg/window/debuginfo"
	"yo-app/pkg/window/loremipsum"
	"yo-app/pkg/window/mainmenu"
)

type Window interface {
	common.Element[basic.None]
}

func New() Window {
	return flexbox.New(func(basic.None) basic.None {
		return basic.None{}
	}, flexbox.Props{
		AlignItems:     flexbox.AlignCenter,
		Direction:      flexbox.DirectionColumn,
		JustifyContent: flexbox.JustifyCenter,
	}.Func,
		turnoff.New(flexbox.Trait{}.Func, func(turnoff.Props) turnoff.Props {
			return turnoff.Props{
				IsOff: !tools.Account.Debug(),
			}
		}, debuginfo.New(turnoff.Trait{}.Func)),

		overlay.New(flexbox.Trait{}.Func, overlay.Props{}.Func,
			colorarea.New(overlay.Trait{}.Func, colorarea.Props{
				Color: assets.Colors.Background,
			}.Func),

			pageset.New(overlay.Trait{}.Func, func(pageset.Props) pageset.Props {
				return pageset.Props{
					Page: vars.Window.Page,
				}
			},
				mainmenu.New(pageset.Trait{}.Func),
				loremipsum.New(pageset.Trait{}.Func),
			),
		),
	)
}
