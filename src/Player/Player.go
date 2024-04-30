package Player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/takaryo1010/GolangGame/src/Platform"
)

type Player struct {
	Posx, Posy    float64
	VelX, VelY    float64
	Height, Width float64
	Speed         float64
	Jump          float64
	IsJumping     bool
	Op            ebiten.DrawImageOptions
	Img           *ebiten.Image
	Gravity       float64
}

func (player *Player) PlayerMove(screenHeight, screenWidth float64, platforms []Platform.Platform) {
	// ジャンプ処理
	isOnGround := false

	// 重力処理
	player.VelY += player.Gravity
	player.Posy += player.VelY

	// プラットフォームとの当たり判定
	for _, platform := range platforms {
		if player.Posy+player.Height >= platform.Posy && // プレイヤーの底辺がプラットフォームの上辺よりも上にあり
			player.Posy+player.Height <= platform.Posy+10 && // プレイヤーの底辺がプラットフォームの上辺から一定の距離以内にあり
			player.Posx+player.Width >= platform.Posx && // プレイヤーがプラットフォームの左端よりも右にいて
			player.Posx <= platform.Posx+platform.Width { // プレイヤーがプラットフォームの右端よりも左にいる場合
			player.Posy = platform.Posy - player.Height // プレイヤーをプラットフォームの上に乗せる
			player.IsJumping = false                    // ジャンプ状態を解除
			player.VelY = 0                             // 縦方向の速度をリセット
			isOnGround = true
			break
		}
	}
	// 台に下から当たった場合の処理
	if !isOnGround {
		for _, platform := range platforms {
			if player.Posy <= platform.Posy+platform.Height && // プレイヤーの頭が台の下辺よりも下にあり
				player.Posy >= platform.Posy && // プレイヤーの頭が台の上辺よりも上にあり
				player.Posx+player.Width >= platform.Posx && // プレイヤーが台の左端よりも右にいて
				player.Posx <= platform.Posx+platform.Width { // プレイヤーが台の右端よりも左にいる場合
				player.Posy = platform.Posy + platform.Height // プレイヤーを台の下に戻す
				player.VelY = 0                               // 縦方向の速度をリセット
				break
			}
		}
	}
	// 画面下部の床との当たり判定
	if player.Posy >= screenHeight-player.Height {
		player.Posy = screenHeight - player.Height
		player.IsJumping = false
		player.VelY = 0
		player.Jump = 10
		isOnGround = true
	}
	if isOnGround {
		if ebiten.IsKeyPressed(ebiten.KeyW) {
			if !player.IsJumping {
				player.VelY = -player.Jump
				player.IsJumping = true
			}
		}
	}
	// 横方向の当たり判定とすり抜け防止処理
	for _, platform := range platforms {
		if player.Posy+player.Height > platform.Posy && // プレイヤーがプラットフォームの上にいて
			player.Posy < platform.Posy+platform.Height && // プレイヤーがプラットフォームの下にいて
			player.Posx+player.Width > platform.Posx && // プレイヤーがプラットフォームの左端よりも右にいて
			player.Posx < platform.Posx+platform.Width { // プレイヤーがプラットフォームの右端よりも左にいる場合
			// プレイヤーが台の左側から当たった場合
			if player.Posx+player.Width > platform.Posx && player.Posx+player.Width < platform.Posx+player.Speed {
				player.Posx = platform.Posx - player.Width
				player.VelX = 0
			}
			// プレイヤーが台の右側から当たった場合
			if player.Posx < platform.Posx+platform.Width && player.Posx > platform.Posx+platform.Width-player.Speed {
				player.Posx = platform.Posx + platform.Width
				player.VelX = 0
			}
		}
	}
	// 横方向の移動処理
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		player.VelX -= player.Speed / 9
		if player.VelX < -player.Speed {
			player.VelX = -player.Speed
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		player.VelX += player.Speed / 9
		if player.VelX > player.Speed {
			player.VelX = player.Speed
		}
	} else {
		if player.IsJumping {
			if player.VelX > 0 {
				player.VelX *= 0.99
			} else if player.VelX < 0 {
				player.VelX *= 0.99
			}
		} else {
			if player.VelX > 0 {
				player.VelX *= 0.9
			} else if player.VelX < 0 {
				player.VelX *= 0.9
			}
		}
	}

	// 横方向の座標更新
	player.Posx += player.VelX
	if player.VelY > 10 {
		player.VelY = 10
	}
	// 画面端から出ないようにする処理
	if player.Posx < 0 {
		player.Posx = screenWidth
	} else if player.Posx > screenWidth {
		player.Posx = 0
	}
}

func (player *Player) MoveDebug(cameraX float64, cameraY float64) {
	if ebiten.IsKeyPressed(ebiten.KeyW) {

		player.Posy -= 20

	}else if ebiten.IsKeyPressed(ebiten.KeyS) {
		
		player.Posy += 20	
	}
}
