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
	lines map[int]*line
	line  int

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
	ts.lines = make(map[int]*line)
	ts.line = 0

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

	if ts.lines[ts.line] == nil {
		ts.lines[ts.line] = &line{}
	}

	length := len(ts.lines[ts.line].text)
	switch key {
	case engo.KeyBackspace:
		if length > 0 {
			ts.lines[ts.line].text = ts.lines[ts.line].text[0 : length-1]
			ts.lines[ts.line].chars[length-1].Remove(ts.world)
			ts.lines[ts.line].chars = ts.lines[ts.line].chars[0 : length-1]
		}
	case engo.KeyEnter:
		ts.lines[ts.line].locked = true

		xoffset := float64(len(ts.lines[ts.line].text)*16) * .65
		if xoffset > 710 {
			ts.line += int(math.Floor(xoffset/710)) + 1
		} else {
			ts.line++
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

		ts.lines[ts.line].text = append(ts.lines[ts.line].text, symbol)

		char := ui.NewText(symbol)

		var xoffset, yoffset float32
		xoffset = float32(len(ts.lines[ts.line].text)*int(char.Font.Size)) * .65
		yoffset = float32(ts.line * int(char.Font.Size))

		if xoffset >= 710 {
			xoffset = xoffset - 710
			yoffset += float32(char.Font.Size)
		}

		char.X = 35 + xoffset
		char.Y = 35 + yoffset

		char.Insert(ts.world)
		ts.lines[ts.line].chars = append(ts.lines[ts.line].chars, char)
	}
}
