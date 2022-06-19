package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/magnets/crane"
	"github.com/jakecoffman/magnets/crane/constant"
	"log"
)

func main() {
	ebiten.SetWindowSize(constant.ScreenWidth, constant.ScreenHeight)
	ebiten.SetWindowTitle("Magnets")
	if err := ebiten.RunGame(crane.NewGame()); err != nil {
		log.Fatal(err)
	}
}
