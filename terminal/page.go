package terminal

type page struct {
	lines map[int]*line
	line  int
	enil  int

	escapable bool
	readonly  bool
}

func (p *page) lineOffset() int {
	return p.line - p.enil
}

func (p *page) show() {
	for i, line := range p.lines {
		for _, char := range line.chars {
			if i >= p.enil {
				char.RenderComponent.Hidden = false
			}
		}
	}
}

func (p *page) hide() {
	for _, line := range p.lines {
		for _, char := range line.chars {
			char.RenderComponent.Hidden = true
		}
	}
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

// func (p *page) pushScreenDown() {
// 	for i, line := range p.lines {
// 		for _, char := range line.chars {
// 			char.SetY(char.Y + 16)

// 			if i >= p.enil {
// 				char.RenderComponent.Hidden = false
// 			}

// 			// if i <= p.enil {
// 			// 	char.RenderComponent.Hidden = false
// 			// } else if i >= p.line {
// 			// 	char.RenderComponent.Hidden = true
// 			// }
// 		}
// 	}

// 	p.enil--
// }
