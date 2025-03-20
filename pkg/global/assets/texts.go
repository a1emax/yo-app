package assets

import (
	"github.com/a1emax/youngine/asset/format/text"
	"github.com/a1emax/youngine/x/scope"
)

var Texts struct {
	Back       text.Asset
	LoremIpsum text.Asset
	StartDebug text.Asset
	StopDebug  text.Asset
	Yo         text.Asset
}

func initTexts(lc scope.Lifecycle) {
	Texts.Back = load[text.Asset](lc, "texts/back.txt")
	Texts.LoremIpsum = load[text.Asset](lc, "texts/lorem-ipsum.txt")
	Texts.StartDebug = load[text.Asset](lc, "texts/start-debug.txt")
	Texts.StopDebug = load[text.Asset](lc, "texts/stop-debug.txt")
	Texts.Yo = load[text.Asset](lc, "texts/yo.txt")
}
