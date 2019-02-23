package terminal

type page struct {
	lines map[int]*line
	line  int
	enil  int
}

func (p *page) lineOffset() int {
	return p.line - p.enil
}

func (p *page) pushScreenUp() {
	for i, line := range p.lines {
		for _, char := range line.chars {
			char.SetY(char.Y - 16)

			if i <= p.enil {
				char.RenderComponent.Hidden = true
			}
		}
	}

	p.enil++
}
