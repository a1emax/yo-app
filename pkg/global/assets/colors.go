package assets

import (
	"github.com/a1emax/youngine/asset/format/rgba"
	"github.com/a1emax/youngine/x/scope"
)

var Colors struct {
	Background          rgba.Asset
	ButtonPressed       rgba.Asset
	ButtonPrimary       rgba.Asset
	DebugInfoBackground rgba.Asset
	DebugInfoText       rgba.Asset
	Info                rgba.Asset
}

func initColors(lc scope.Lifecycle) {
	Colors.Background = load[rgba.Asset](lc, "colors/background.rgba")
	Colors.ButtonPressed = load[rgba.Asset](lc, "colors/button-pressed.rgba")
	Colors.ButtonPrimary = load[rgba.Asset](lc, "colors/button-primary.rgba")
	Colors.DebugInfoBackground = load[rgba.Asset](lc, "colors/debug-info-background.rgba")
	Colors.DebugInfoText = load[rgba.Asset](lc, "colors/debug-info-text.rgba")
	Colors.Info = load[rgba.Asset](lc, "colors/info.rgba")
}
