package terminal

import (
	"github.com/lukevers/arp147/ui"
)

type page struct {
	commands []string
	cmdindex int

	lines map[int]*line
	line  int
	enil  int

	cursor *ui.Text
	cpoint int

	escapable bool
	readonly  bool
	editable  bool
}

func (p *page) ToString() string {
	contents := ""
	for i := 0; i <= p.line; i++ {
		contents += p.lines[i].String() + "\n"
	}

	return contents
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

	p.cursor.RenderComponent.Hidden = false
}

func (p *page) hide() {
	p.cursor.RenderComponent.Hidden = true

	for _, line := range p.lines {
		for _, char := range line.chars {
			char.RenderComponent.Hidden = true
		}
	}
}

func (p *page) pushScreenUp() {
	lastVisibleIndex := p.getLastVisibleLineIndex()
	for i, line := range p.lines {
		for _, char := range line.chars {
			char.SetY(char.Y - FontSize)

			if i <= p.enil {
				char.RenderComponent.Hidden = true
			} else if i == lastVisibleIndex+1 {
				char.RenderComponent.Hidden = false
			}
		}
	}

	p.enil++
}

func (p *page) getLastVisibleLineIndex() int {
	for i := len(p.lines) - 1; i > 0; i-- {
		if p.lines[i] == nil {
			continue
		}

		if len(p.lines[i].chars) > 0 {
			if !p.lines[i].chars[0].RenderComponent.Hidden {
				return i
			}
		} else {
			// Handle empty lines that still exist
			if float32(i)*FontSize <= 704 {
				return i
			}
		}
	}

	return 0
}

func (p *page) pushScreenDown() {
	lastVisibleIndex := p.getLastVisibleLineIndex()
	for i, line := range p.lines {
		for _, char := range line.chars {
			char.SetY(char.Y + FontSize)

			if i == p.line && i <= p.enil {
				char.RenderComponent.Hidden = false
			} else if i == lastVisibleIndex {
				char.RenderComponent.Hidden = true
			}
		}
	}

	p.enil--
}
