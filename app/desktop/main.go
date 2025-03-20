package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"

	"yo-app/pkg/global/vars"
	"yo-app/pkg/kernel"
)

func main() {
	err := run()
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}

func run() (_err error) {
	defer func() {
		_err = errors.Join(_err, kernel.Close())
	}()

	vars.Extern.FilesDir = "."

	kernel.Activate()

	ebiten.SetWindowTitle("Yo App")
	ebiten.SetWindowSize(392, 822) // 384x799
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	return ebiten.RunGame(kernel.EbitenGame())
}
