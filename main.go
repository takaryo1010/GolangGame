package main

import (
    "fmt"
    "image/color"
    "log"
 
    "strings"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
    "github.com/hajimehoshi/ebiten/v2/text"
    "github.com/takaryo1010/GolangGame/src/Platform"
    "github.com/takaryo1010/GolangGame/src/Player"
    "golang.org/x/image/font"
    "golang.org/x/image/font/opentype"
)

var player Player.Player
var platforms Platform.Platforms
var mplusNormalFont font.Face
var lastY float64
var IsKeyPressed0 bool

func init() {
    player = Player.Player{
        Posx:    screenWidth / 2,
        Posy:    screenHeight,
        Height:  8,
        Width:   8.0,
        Gravity: 0.6,
        VelX:    0, // 横方向の速度を初期化
        VelY:    0,
        Jump:    10,
        Speed:   5,
    }
    lastY = screenHeight
    player.Img = ebiten.NewImage(int(player.Width), int(player.Height))
    player.Img.Fill(color.White)
    platforms, lastY = Platform.InitPlatform(100, screenWidth, screenHeight, screenHeight)
    tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
    if err != nil {
        log.Fatal(err)
    }

    const dpi = 72
    mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
        Size:    32,
        DPI:     dpi,
        Hinting: font.HintingVertical,
    })
    if err != nil {
        log.Fatal(err)
    }
}

const (
    screenWidth  = 640
    screenHeight = 480
)

type Game struct {
    playerName      string  // プレイヤーの名前
    currentInput    string  // 現在の入力中の文字列
    inputInProgress bool    // 入力が進行中かどうかのフラグ
    greeting        string
    cameraX         float64 // カメラのX座標
    cameraY         float64 // カメラのY座標
    hiscore         int     // ハイスコア
    started         bool    // ゲームが開始されたかどうかのフラグ
}

func (g *Game) Update() error {
    if !g.started {
        // ゲームが開始されていない場合は、名前の入力を受け付ける
        inputChars := ebiten.InputChars()
        for _, char := range inputChars {
            if ebiten.IsKeyPressed(ebiten.KeyEnter) {
                // 改行が入力されたら、入力を終了し名前を確定する
                g.playerName = g.currentInput
                g.currentInput = ""
                g.inputInProgress = false
                g.started = true // ゲームを開始する
			} else if(ebiten.IsKeyPressed(ebiten.KeyBackspace)){
                // バックスペースが入力されたら、入力文字列から最後の文字を削除する
                if len(g.currentInput) > 0 {
                    g.currentInput = g.currentInput[:len(g.currentInput)-1]
                }
			}else{
				// それ以外の場合は入力文字列に追加する
				g.currentInput += string(char)
				fmt.Println(char)
			}
			
        }

        if !g.inputInProgress {
            // 入力が進行中でない場合は、名前の入力を促す
            g.greeting = "Enter your name: " + g.currentInput
        }
        return nil
    }

    // ゲームが開始されている場合はゲームの更新を行う
    greeting := fmt.Sprintf("NowHeight:%dbit", -(int(player.Posy)-472))
    g.greeting = greeting
    if g.hiscore < -(int(player.Posy)-464) {
        g.hiscore = -(int(player.Posy) - 464)
    }
    // プレイヤーの位置をカメラの中心にする
    g.cameraX = player.Posx - float64(screenWidth)/2
    g.cameraY = player.Posy - float64(screenHeight)/3*2

    if -(player.Posy-platforms[0].Posy) > 5000 {
        // 50ピクセル以上下に移動した場合、プラットフォームを削除し、新しいプラットフォームを生成
        newplatforms, newY := Platform.InitPlatform(1, screenWidth, screenHeight, lastY)
        lastY = newY
        platforms = append(platforms[1:], newplatforms...)
    }
    if len(platforms) > 400 {
        platforms = platforms[1:]
    }
    player.PlayerMove(screenHeight, screenWidth, platforms)
    // player.MoveDebug(g.cameraX, g.cameraY)
    if !ebiten.IsKeyPressed(ebiten.Key0) {
        IsKeyPressed0 = false
    }
    if ebiten.IsKeyPressed(ebiten.Key0) && !IsKeyPressed0 {
        // ゲームを再開するためにプレイヤーの位置、プラットフォームなどをリセット
        player.Posx = screenWidth / 2
        player.Posy = screenHeight
        lastY = screenHeight
        platforms = nil
        platforms, lastY = Platform.InitPlatform(100, screenWidth, screenHeight, screenHeight)
        IsKeyPressed0 = true
    }
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    if !g.started {
        // ゲームが開始されていない場合は、名前の入力を描画する
        text.Draw(screen, g.greeting, mplusNormalFont, 200, 200, color.White)
        return
    }

    // ゲームが開始されている場合は、ゲームの要素を描画する
    text.Draw(screen, fmt.Sprintf("Player: %s", g.playerName), mplusNormalFont, 200, 200, color.White)
    screen.Fill(color.RGBA{0, 0, 0, 0})

    // カメラの位置を考慮してプレイヤーを描画
    player.Op = ebiten.DrawImageOptions{}
    player.Op.GeoM.Translate(player.Posx, player.Posy-g.cameraY)
    screen.DrawImage(player.Img, &player.Op)

    // カメラの位置を考慮してプラットフォームを描画
    platforms.Draw(screen, g.cameraX, g.cameraY)

    text.Draw(screen, g.greeting, mplusNormalFont, 0, 32, color.White)
    text.Draw(screen, fmt.Sprintf("HiScore      :%dbit", g.hiscore), mplusNormalFont, 0, 64, color.White)

    ground := ebiten.NewImage(screenWidth, 300)
    ground.Fill(color.RGBA{0, 255, 0, 0})
    opGround := ebiten.DrawImageOptions{}
    opGround.GeoM.Translate(0, screenHeight-g.cameraY)
    screen.DrawImage(ground, &opGround)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return 640, 480
}

func (g *Game) SetName(name string) {
    // 入力された名前を検証して設定
    g.playerName = strings.TrimSpace(name)
    // ゲームを開始するフラグを設定
    g.started = true
}

func main() {
    ebiten.SetWindowSize(screenWidth, screenHeight)
    ebiten.SetWindowTitle("Hello, World!")

    // ゲームのインスタンスを作成
    game := &Game{
        hiscore: 0,
        started:true,
    }

    // ゲームを実行
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}
