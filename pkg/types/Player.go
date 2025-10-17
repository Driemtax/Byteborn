package types

import "fmt"

type Player struct {
	Posx int
	Posy int
}

func NewPlayer() *Player {
	return &Player{
		Posx: 180,
		Posy: 180,
	}
}

func (p *Player) Move(d Direction) error {
	switch d {
	case UP:
		fmt.Println("Moving up..", p.Posy)
		if p.Posy >= 10 {
			p.Posy -= 10
			fmt.Println("Moved to:", p.Posy)
		}
	case DOWN:
		if p.Posy <= 389 {
			p.Posy += 10
		}
	case LEFT:
		if p.Posx >= 10 {
			p.Posx -= 10
		}
	case RIGHT:
		if p.Posx <= 390 {
			p.Posx += 10
		}
	}
	return nil
}
