package terminal

import (
	"log"
	"math"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/lukevers/arp147/input"
	"github.com/lukevers/arp147/ui"
)

// TerminalSystem is a scrollable, visual and text input-able system.
type TerminalSystem struct {
	pages map[int]*page
	page  int

	world *ecs.World
}

type TerminalViewer struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// Remove is called whenever an Entity is removed from the World, in order to
// remove it from this sytem as well.
func (*TerminalSystem) Remove(ecs.BasicEntity) {
	// TODO
}

// Update is ran every frame, with `dt` being the time in seconds since the
// last frame.
func (*TerminalSystem) Update(dt float32) {
	// TODO
}

// New is the initialisation of the System.
func (ts *TerminalSystem) New(w *ecs.World) {
	ts.world = w
	ts.pages = make(map[int]*page)

	ts.registerKeys()
	ts.addBackground(w)

	log.Println("TerminalSystem initialized")
}

func (ts *TerminalSystem) addBackground(w *ecs.World) {
	bkg1 := &TerminalViewer{BasicEntity: ecs.NewBasic()}
	bkg2 := &TerminalViewer{BasicEntity: ecs.NewBasic()}
	bkg3 := &TerminalViewer{BasicEntity: ecs.NewBasic()}

	bkg1.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 0, Y: 0},
		Width:    800,
		Height:   800,
	}

	tbkg1, err := common.LoadedSprite("textures/bkg_t1.jpg")
	if err != nil {
		log.Fatal("Unable to load texture: " + err.Error())
	}

	bkg1.RenderComponent = common.RenderComponent{
		Drawable: tbkg1,
		Scale:    engo.Point{X: 1, Y: 1},
	}

	bkg2.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 800, Y: 0},
		Width:    400,
		Height:   400,
	}

	tbkg2, err := common.LoadedSprite("textures/bkg_t2.jpg")
	if err != nil {
		log.Fatal("Unable to load texture: " + err.Error())
	}

	bkg2.RenderComponent = common.RenderComponent{
		Drawable: tbkg2,
		Scale:    engo.Point{X: 1, Y: 1},
	}

	bkg3.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 800, Y: 400},
		Width:    400,
		Height:   400,
	}

	tbkg3, err := common.LoadedSprite("textures/bkg_t3.jpg")
	if err != nil {
		log.Fatal("Unable to load texture: " + err.Error())
	}

	bkg3.RenderComponent = common.RenderComponent{
		Drawable: tbkg3,
		Scale:    engo.Point{X: 1, Y: 1},
	}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&bkg1.BasicEntity, &bkg1.RenderComponent, &bkg1.SpaceComponent)
			sys.Add(&bkg2.BasicEntity, &bkg2.RenderComponent, &bkg2.SpaceComponent)
			sys.Add(&bkg3.BasicEntity, &bkg3.RenderComponent, &bkg3.SpaceComponent)
		}
	}
}

func (ts *TerminalSystem) registerKeys() {
	input.RegisterKeys([]input.Key{
		input.Key{
			Name: "terminal-keys",
			Keys: []engo.Key{
				engo.KeyA,
				engo.KeyB,
				engo.KeyC,
				engo.KeyD,
				engo.KeyE,
				engo.KeyF,
				engo.KeyG,
				engo.KeyH,
				engo.KeyI,
				engo.KeyJ,
				engo.KeyK,
				engo.KeyL,
				engo.KeyM,
				engo.KeyN,
				engo.KeyO,
				engo.KeyP,
				engo.KeyQ,
				engo.KeyR,
				engo.KeyS,
				engo.KeyT,
				engo.KeyU,
				engo.KeyV,
				engo.KeyW,
				engo.KeyX,
				engo.KeyY,
				engo.KeyZ,

				engo.KeyZero,
				engo.KeyOne,
				engo.KeyTwo,
				engo.KeyThree,
				engo.KeyFour,
				engo.KeyFive,
				engo.KeySix,
				engo.KeySeven,
				engo.KeyEight,
				engo.KeyNine,

				engo.KeyBackspace,
				engo.KeyEnter,
				engo.KeySpace,
				engo.KeyTab,
				engo.KeyEscape,

				engo.KeyDash,
				engo.KeyGrave,
				engo.KeyApostrophe,
				engo.KeySemicolon,
				engo.KeyEquals,
				engo.KeyComma,
				engo.KeyPeriod,
				engo.KeySlash,
				engo.KeyBackslash,
				engo.KeyLeftBracket,
				engo.KeyRightBracket,
			},
			OnPress: ts.delegateKeyPress,
		},
	})
}

func (ts *TerminalSystem) delegateKeyPress(key engo.Key, mods *input.Modifiers) {
	log.Println(key, mods)

	if ts.pages[ts.page] == nil {
		ts.pages[ts.page] = &page{
			lines: make(map[int]*line),
			line:  0,
		}
	}

	if ts.pages[ts.page].lines[ts.pages[ts.page].line] == nil {
		ts.pages[ts.page].lines[ts.pages[ts.page].line] = newLine()
	}

	length := len(ts.pages[ts.page].lines[ts.pages[ts.page].line].text)
	switch key {
	case engo.KeyBackspace:
		if length > 0 {
			ts.pages[ts.page].lines[ts.pages[ts.page].line].text = ts.pages[ts.page].lines[ts.pages[ts.page].line].text[0 : length-1]
			ts.pages[ts.page].lines[ts.pages[ts.page].line].chars[length-1].Remove(ts.world)
			ts.pages[ts.page].lines[ts.pages[ts.page].line].chars = ts.pages[ts.page].lines[ts.pages[ts.page].line].chars[0 : length-1]
		}
	case engo.KeyEnter:
		ts.pages[ts.page].lines[ts.pages[ts.page].line].locked = true

		xoffset := ts.getXoffset()
		if xoffset > 710 {
			ts.pages[ts.page].line += int(math.Floor(float64(xoffset)/710)) + 1
		} else {
			ts.pages[ts.page].line++
		}

		yoffset := float32(ts.pages[ts.page].line * int(16))
		if yoffset > 704 {
			ts.pages[ts.page].pushScreenUp()
		}
	default:
		var symbol string

		// If the key is [a-z] apply shift rules.
		if key >= engo.KeyA && key <= engo.KeyZ {
			if mods.Shift {
				symbol = string(key)
			} else {
				symbol = string(key + 32)
			}
		} else {
			// Convert non [a-z] letters when shift is used
			if mods.Shift {
				// TODO
				//   - see above
				//   - will this be different for different keyboard layouts?
				//     - we can't assume everyone uses US QWERTY
				//	   - I should learn how other layouts work
				symbol = "*"
			} else {
				// Otherwise we just use the actual key here
				symbol = string(key)
			}
		}

		ts.pages[ts.page].lines[ts.pages[ts.page].line].text = append(ts.pages[ts.page].lines[ts.pages[ts.page].line].text, symbol)

		char := ui.NewText(symbol)

		push := false
		var xoffset, yoffset float32
		xoffset = ts.getXoffset()
		yoffset = float32(ts.pages[ts.page].lineOffset() * 16)

		if xoffset >= 710 {
			lines := int(math.Floor(float64(xoffset) / 710))

			xoffset = xoffset - float32(707*lines)
			yoffset += float32(16) * float32(lines)

			if yoffset > 704 {
				push = true
			}
		}

		char.X = 35 + xoffset
		char.Y = 35 + yoffset

		char.Insert(ts.world)
		ts.pages[ts.page].lines[ts.pages[ts.page].line].chars = append(ts.pages[ts.page].lines[ts.pages[ts.page].line].chars, char)

		if push {
			ts.pages[ts.page].pushScreenUp()
		}
	}
}

func (ts *TerminalSystem) getXoffset() float32 {
	if ts.pages[ts.page].lines[ts.pages[ts.page].line] == nil {
		return 0
	}

	return float32(len(ts.pages[ts.page].lines[ts.pages[ts.page].line].text)*int(16)) * .65
}
