package types

type Player struct {
	SizeX int
	SizeY int
	Posx  int
	Posy  int
}

func NewPlayer() *Player {
	return &Player{
		SizeX: 40,
		SizeY: 40,
		Posx:  380,
		Posy:  380,
	}
}

func (p *Player) Move(d Direction) error {
	switch d {
	case UNDEFINED:
	case UP:
		// if u ask yourself why i subtracted 10 here, you asked a very good question, but it works so get on with it :)
		if p.Posy >= (p.SizeY/2)-10 {
			p.Posy -= 10
		}
	case DOWN:
		if p.Posy <= 790-(p.SizeY) {
			p.Posy += 10
		}
	case LEFT:
		if p.Posx >= (p.SizeX/2)-10 {
			p.Posx -= 10
		}
	case RIGHT:
		if p.Posx <= 790-(p.SizeX) {
			p.Posx += 10
		}
	}
	return nil
}
