package game

import (
	"game/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image             *ebiten.Image
	position          Vector
	game              *Game
	LaserLoadingTimer *Timer
}

func NewPlayer(game *Game) *Player {
	image := assets.PlayerSprite

	bounds := image.Bounds()
	halfWidth := float64(bounds.Dx()) / 2

	position := Vector{
		X: (screenWidth / 2) - halfWidth,
		Y: 500,
	}
	return &Player{
		image:             image,
		position:          position,
		game:              game,
		LaserLoadingTimer: NewTime(12),
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	//Posicao na tela que a imagem irar ficar na tela
	op.GeoM.Translate(p.position.X, p.position.Y)
	//Desenhar imagem
	screen.DrawImage(p.image, op)
}

func (p *Player) Update() {

	spped := 6.0
	// Movimentar nave para esquerda
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.X -= spped
	}
	//Movimentar nave para direita
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.X += spped
	}

	p.LaserLoadingTimer.Update()
	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.LaserLoadingTimer.IsReady() {
		p.LaserLoadingTimer.Reset()
		bounds := p.image.Bounds()
		halfWidth := float64(bounds.Dx()) / 2
		halfHeight := float64(bounds.Dy()) / 2

		spawnPos := Vector{
			p.position.X + halfWidth,
			p.position.Y - halfHeight/2,
		}

		laser := NewLaser(spawnPos)

		p.game.AddLasers(laser)

	}
}

//func (p *Player) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {}

func (p *Player) Collider() Rect {
	bounds := p.image.Bounds()
	return NewRect(p.position.X, p.position.Y, float64(bounds.Dx()), float64(bounds.Dy()))
}
