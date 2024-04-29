package Player

import "github.com/hajimehoshi/ebiten"

type Player struct {
	posx, posy    float64
	velX, velY    float64
	height, width float64
	speed         float64
	jump          float64
	isJumping     bool
	op            ebiten.DrawImageOptions
	img           *ebiten.Image
	gravity       float64
}

func (player *Player) PlayerMove(screenHeight float64, screenWidth float64) {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		if !player.isJumping {
			player.velY = -player.jump
			player.isJumping = true
		}
	}

	player.velY += player.gravity
	player.posy += player.velY

	if player.posy >= screenHeight/3*2 {
		player.posy = screenHeight / 3 * 2
		player.isJumping = false
		player.velY = 0
		player.jump = 15
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		player.velX -= player.speed / 10
		if player.velX < -player.speed {
			player.velX = -player.speed
		}

	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		player.velX += player.speed / 10
		if player.velX > player.speed {
			player.velX = player.speed
		}
	} else {
		if player.isJumping {
			if player.velX > 0 {
				player.velX *= 0.99

			} else if player.velX < 0 {
				player.velX *= 0.99

			}
		} else {
			if player.velX > 0 {
				player.velX *= 0.9

			} else if player.velX < 0 {
				player.velX *= 0.9

			}
		}

	}

	player.posx += player.velX

	if player.posx < 0 {
		player.posx = screenWidth
	} else if player.posx > screenWidth {
		player.posx = 0
	}

}
