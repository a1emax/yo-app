package debuginfo

import (
	"fmt"
	"runtime"

	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/scene"
	"github.com/a1emax/youngine/scene/element/flexbox"
	"github.com/a1emax/youngine/scene/element/overlay"
	"github.com/a1emax/youngine/scene/element/x/colorarea"
	"github.com/a1emax/youngine/scene/element/x/textviewer"
	"github.com/hajimehoshi/ebiten/v2"

	"yo-app/pkg/global/assets"
	"yo-app/pkg/global/consts"
	"yo-app/pkg/global/vars"
	"yo-app/pkg/window/common"
)

type DebugInfo[T any] interface {
	common.Element[T]
}

func New[T any](traitFunc scene.TraitFunc[T]) DebugInfo[T] {
	var maxMemSys uint64
	var maxMemPauseNs uint64

	return overlay.New(traitFunc, overlay.Props{
		Attrs: scene.Attrs{}.
			FixHeight(consts.DebugInfoHeight),
	}.Func,
		colorarea.New(overlay.Trait{}.Func, colorarea.Props{
			Color: assets.Colors.DebugInfoBackground,
		}.Func),

		flexbox.New(overlay.Trait{}.Func, flexbox.Props{
			AlignItems:     flexbox.AlignCenter,
			Direction:      flexbox.DirectionRow,
			JustifyContent: flexbox.JustifySpaceBetween,
		}.Func,
			textviewer.New(flexbox.Trait{}.Func, func(textviewer.Props) textviewer.Props {
				text := fmt.Sprintf("%.2f / %.2f", ebiten.ActualFPS(), ebiten.ActualTPS())

				return textviewer.Props{
					Attrs: scene.Attrs{}.
						FixWidth(consts.DebugInfoItemWidth),

					FontFace:  assets.FontFaces.DebugInfoText,
					Text:      text,
					TextColor: assets.Colors.DebugInfoText,
				}
			}),

			textviewer.New(flexbox.Trait{}.Func, func(textviewer.Props) textviewer.Props {
				text := fmt.Sprintf("%dx%d / %dx%d",
					vars.Ebiten.ScreenWidth, vars.Ebiten.ScreenHeight,
					vars.Ebiten.OutsideWidth, vars.Ebiten.OutsideHeight,
				)

				return textviewer.Props{
					Attrs: scene.Attrs{}.
						FixWidth(consts.DebugInfoItemWidth),

					FontFace:  assets.FontFaces.DebugInfoText,
					Text:      text,
					TextColor: assets.Colors.DebugInfoText,
				}
			}),

			textviewer.New(flexbox.Trait{}.Func, func(textviewer.Props) textviewer.Props {
				var mem runtime.MemStats
				runtime.ReadMemStats(&mem)

				if mem.Sys > maxMemSys {
					maxMemSys = mem.Sys
				}

				memPauseNs := mem.PauseNs[(mem.NumGC+255)%256]
				if memPauseNs > maxMemPauseNs {
					maxMemPauseNs = memPauseNs
				}

				text := fmt.Sprintf("%.2f MiB (%.2f ms)",
					basic.Float(maxMemSys)/(1024*1024),
					basic.Float(maxMemPauseNs)/1_000_000,
				)

				return textviewer.Props{
					Attrs: scene.Attrs{}.
						FixWidth(consts.DebugInfoItemWidth),

					FontFace:  assets.FontFaces.DebugInfoText,
					Text:      text,
					TextColor: assets.Colors.DebugInfoText,
				}
			}),
		),
	)
}
