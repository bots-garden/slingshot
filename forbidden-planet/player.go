package main

type Player struct {
	avatar    string
	name      string
	x         int
	y         int
	direction Direction
}

func (p *Player) Name() string {
	return p.name
}
func (p *Player) Avatar() string {
	return p.avatar
}

func (p *Player) LookBelow(bords Bords) Boxe {
	objectBoxe := bords.objects.Content(p.x, p.y)
	groundBoxe := bords.ground.Content(p.x, p.y)
	if objectBoxe == Nothing { // nothing is "⬛️"
		return groundBoxe
	} else {
		return objectBoxe
	}
}

func (p *Player) LookAt(direction Direction, bords Bords) Boxe {
	var objectBoxe Boxe
	var groundBoxe Boxe

	switch direction {
	case North:
		// out of map
		if p.y-1 < 0 {
			objectBoxe = Nothing
			groundBoxe = Nothing
		} else {
			objectBoxe = bords.objects.Content(p.x, p.y-1)
			groundBoxe = bords.ground.Content(p.x, p.y-1)
		}

	case South:
		// out of map
		if p.y+1 > bords.ground.height-1 {
			objectBoxe = Nothing
			groundBoxe = Nothing
		} else {
			objectBoxe = bords.objects.Content(p.x, p.y+1)
			groundBoxe = bords.ground.Content(p.x, p.y+1)
		}

	case East:
		// out of map
		if p.x+1 > bords.ground.width-1 {
			objectBoxe = Nothing
			groundBoxe = Nothing
		} else {
			objectBoxe = bords.objects.Content(p.x+1, p.y)
			groundBoxe = bords.ground.Content(p.x+1, p.y)
		}

	case West:
		// out of map
		if p.x-1 < 0 {
			objectBoxe = Nothing
			groundBoxe = Nothing
		} else {
			objectBoxe = bords.objects.Content(p.x-1, p.y)
			groundBoxe = bords.ground.Content(p.x-1, p.y)
		}
	}
	if objectBoxe == Nothing { // nothing is "⬛️"
		return groundBoxe
	} else {
		return objectBoxe
	}
}

func (p *Player) Move(direction Direction, bords Bords) (Direction, int, int) {

	p.direction = direction

	switch direction {
	case North:
		p.y -= 1
		if p.y < 0 {
			p.y = 0
		}
	case South:
		p.y += 1
		if p.y > bords.ground.height-1 {
			p.y = bords.ground.height - 1
		}
	case East:
		p.x += 1
		if p.x > bords.ground.width-1 {
			p.x = bords.ground.width - 1
		}
	case West:
		p.x -= 1
		if p.x < 0 {
			p.x = 0
		}
	}
	return p.direction, p.x, p.y
}