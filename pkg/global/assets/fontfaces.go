package assets

import (
	"github.com/a1emax/youngine/asset/format/sfnt"
	"github.com/a1emax/youngine/x/scope"
)

var FontFaces struct {
	DebugInfoText sfnt.FaceAsset
	Icon          sfnt.FaceAsset
	LargeText     sfnt.FaceAsset
	MediumText    sfnt.FaceAsset
}

func initFontFaces(lc scope.Lifecycle) {
	FontFaces.DebugInfoText = load[sfnt.FaceAsset](lc, "font-faces/debug-info-text.sff")
	FontFaces.Icon = load[sfnt.FaceAsset](lc, "font-faces/icon.sff")
	FontFaces.LargeText = load[sfnt.FaceAsset](lc, "font-faces/large-text.sff")
	FontFaces.MediumText = load[sfnt.FaceAsset](lc, "font-faces/medium-text.sff")
}
