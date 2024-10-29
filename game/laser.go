package game

import (
	"game/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Laser struct {
	image    *ebiten.Image
	position Vector
}

func NewLaser(position Vector) *Laser {
	image := assets.LaserSprite

	bounds := image.Bounds()
	halfWidth := float64(bounds.Dx()) / 2  //Metade da largura da imagem do laser
	halfHeight := float64(bounds.Dy()) / 2 // Metade da altura da imagem do laser

	position.X -= halfWidth
	position.Y -= halfHeight

	return &Laser{
		image:    image,
		position: position,
	}
}

func (l *Laser) Update() {
	spedd := 7.0
	l.position.Y += -spedd
}

func (l *Laser) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(l.position.X, l.position.Y)
	screen.DrawImage(l.image, op)
}

func (l *Laser) Collider() Rect {
	bounds := l.image.Bounds()
	return NewRect(l.position.X, l.position.Y, float64(bounds.Dx()), float64(bounds.Dy()))
}
