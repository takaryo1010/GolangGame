package Platform

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Platform struct {
	Posx, Posy    float64
	Height, Width float64
}

type Platforms []Platform

func InitPlatform(n int, screenHeight float64, screenWidth float64, lastY float64) (Platforms, float64) {
	rand.Seed(time.Now().UnixNano())
	platforms := Platforms{}

	// 台の生成範囲を調整するためのパラメータ
	minPlatformHeight := 20
	maxPlatformHeight := 21
	minPlatformWidth := 100
	maxPlatformWidth := 200
	minVerticalGap := 50
	maxVerticalGap := 70
	minHorizontalGap := 50
	maxHorizontalGap := 200 // 台同士の間隔の範囲を設定

	for i := 0; i < n; i++ {
		// 台の幅と高さをランダムに決定
		width := float64(rand.Intn(maxPlatformWidth-minPlatformWidth) + minPlatformWidth)
		height := float64(rand.Intn(maxPlatformHeight-minPlatformHeight) + minPlatformHeight)

		// 台のX座標をランダムに決定
		x := rand.Float64() * (screenWidth - width)

		// 台のY座標を直前に生成した台のY座標から一定の間隔で下へ配置
		y := lastY - float64(rand.Intn(maxVerticalGap-minVerticalGap)+minVerticalGap)

		// 台が画面の端からはみ出た場合は修正する
		if x+width > screenWidth {
			x = screenWidth - width // 台を画面内に収める
		}
		if x < 0 {
			x = 0 // 台を画面内に収める
		}

		platform := Platform{Posx: x, Posy: y, Height: height, Width: width}
		platforms = append(platforms, platform)

		// 同じ高さに2から3つの台を生成する
		for j := 0; j < rand.Intn(2)+2; j++ {
			platform := Platform{Posx: x, Posy: y, Height: height, Width: width}
			platforms = append(platforms, platform)

			// 台のX座標を少しずらして次の台を生成する
			x += width + float64(rand.Intn(maxHorizontalGap-minHorizontalGap)+minHorizontalGap)
		}

		lastY = y
	}

	return platforms, lastY
}


// Draw は台を描画する
func (platforms *Platforms) Draw(screen *ebiten.Image, cameraX float64, cameraY float64) {
	for _, platform := range *platforms {
		platformImg := ebiten.NewImage(int(platform.Width), int(platform.Height))
		platformImg.Fill(color.RGBA{0, 255, 0, 255})
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(platform.Posx, platform.Posy-cameraY) // カメラのY座標を適用
		screen.DrawImage(platformImg, &op)
	}
}
