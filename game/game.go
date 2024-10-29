package game

import (
	"fmt"
	"game/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
)

type Game struct {
	players           *Player
	laser             []*Laser
	meteors           []*Meteor
	meteorsSpawnTimer *Timer
	score             int
}

func NewGame() *Game {
	g := &Game{
		meteorsSpawnTimer: NewTime(24),
	}
	player := NewPlayer(g)
	g.players = player
	return g
}

// RODE em 60 FPS
// Responsavel para atualizar a logia do jogo
// 60x por segundo
func (g *Game) Update() error {
	g.players.Update()

	for _, l := range g.laser {
		l.Update()
	}

	g.meteorsSpawnTimer.Update()
	if g.meteorsSpawnTimer.IsReady() {
		g.meteorsSpawnTimer.Reset()
		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	for _, m := range g.meteors {
		if m.Collider().Intersects(g.players.Collider()) {
			fmt.Println("Voce perdeu")
			g.Reset()
		}
	}

	for i, m := range g.meteors {
		for j, l := range g.laser {
			if m.Collider().Intersects(l.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.laser = append(g.laser[:j], g.laser[j+1:]...)

				g.score += 1
			}
		}
	}

	return nil
}

// Responsavel por desenhar objetos na tela
// 60x por segundo
func (g *Game) Draw(screen *ebiten.Image) {
	g.players.Draw(screen)
	for _, l := range g.laser {
		l.Draw(screen)
	}

	for _, m := range g.meteors {
		m.Draw(screen)
	}

	text.Draw(screen, fmt.Sprintf("Pontos: %d", g.score), assets.FontUi, 20, 100, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) AddLasers(laser *Laser) {
	g.laser = append(g.laser, laser)
}

func (g *Game) Reset() {
	g.players = NewPlayer(g)
	g.meteors = nil
	g.laser = nil
	g.meteorsSpawnTimer.Reset()
	g.score = 0
}
