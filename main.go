package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/takaryo1010/GolangGame/src/Player"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var player Player.Player

func init() {
	player = Player.Player{
		posx:    screenWidth / 2,
		posy:    screenHeight / 2,
		height:  8,
		width:   8.0,
		gravity: 0.6,
		velX:    0, // 横方向の速度を初期化
		velY:    0,
		jump:    15,
		speed:   5,
	}
	player.img = ebiten.NewImage(int(player.width), int(player.height))
	player.img.Fill(color.White)
}

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	greeting string
}

func (g *Game) Update() error {
	greeting := fmt.Sprintf("%d:%d", int(player.posx), int(player.posy))
	g.greeting = greeting

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 0})
	player.op = ebiten.DrawImageOptions{}
	player.op.GeoM.Translate(player.posx, player.posy)
	screen.DrawImage(player.img, &player.op)
	ebitenutil.DebugPrint(screen, g.greeting)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
