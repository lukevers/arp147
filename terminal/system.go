package terminal

import (
	"log"
	"math"
	"strings"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/lukevers/arp147/input"
	"github.com/lukevers/arp147/navigator"
	"github.com/lukevers/arp147/ship"
	"github.com/lukevers/arp147/terminal/filesystem"
	"github.com/lukevers/arp147/ui"
)

// TerminalSystem is a scrollable, visual and text input-able system.
type TerminalSystem struct {
	Map *navigator.Map

	pages map[int]*page
	page  int

	world *ecs.World
	vfs   *filesystem.VirtualFS

	ship *ship.Ship
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
	ts.vfs = filesystem.New(ts.WriteError)
	ts.world = w
	ts.pages = make(map[int]*page)
	ts.pages[ts.page] = &page{
		lines:    make(map[int]*line),
		line:     0,
		readonly: true,
	}

	ts.registerKeys()
	ts.addBackground(w)

	result := make(chan bool)
	go ts.loginScript(&result)

	go (func() {
		<-result
		ts.pages[ts.page].lines[ts.pages[ts.page].line] = &line{}
		ts.pages[ts.page].lines[ts.pages[ts.page].line].prefix(ts.delegateKeyPress)
	})()

	engo.Mailbox.Listen("NewShipMessage", func(message engo.Message) {
		msg, ok := message.(ship.NewShipMessage)
		if !ok {
			return
		}

		ts.ship = msg.Ship
	})

	log.Println("TerminalSystem initialized")
}

func (ts *TerminalSystem) addBackground(w *ecs.World) {
	bkg := &TerminalViewer{BasicEntity: ecs.NewBasic()}

	bkg.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 0, Y: 0},
		Width:    800,
		Height:   800,
	}

	tbkg, err := common.LoadedSprite("textures/bkg_t1.jpg")
	if err != nil {
		log.Fatal("Unable to load texture: " + err.Error())
	}

	bkg.RenderComponent = common.RenderComponent{
		Drawable: tbkg,
		Scale:    engo.Point{X: 1, Y: 1},
	}

	bkg.RenderComponent.SetZIndex(0)

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&bkg.BasicEntity, &bkg.RenderComponent, &bkg.SpaceComponent)
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

				engo.KeyArrowUp,
				engo.KeyArrowDown,
			},
			OnPress: ts.delegateKeyPress,
		},
	})
}

func (ts *TerminalSystem) delegateKeyPress(key engo.Key, mods *input.Modifiers) {
	if ts.pages[ts.page] == nil {
		ts.pages[ts.page] = &page{
			lines: make(map[int]*line),
			line:  0,
		}
	}

	if ts.pages[ts.page].lines[ts.pages[ts.page].line] == nil {
		ts.pages[ts.page].lines[ts.pages[ts.page].line] = &line{}
		ts.pages[ts.page].lines[ts.pages[ts.page].line].prefix(ts.delegateKeyPress)
	}

	length := len(ts.pages[ts.page].lines[ts.pages[ts.page].line].text)
	prefixCount := ts.pages[ts.page].lines[ts.pages[ts.page].line].prefixCount
	switch key {
	case engo.KeyTab:
		// TODO: auto-completion?
	case engo.KeyBackspace:
		if ts.pages[ts.page].readonly && !mods.Output {
			break
		}

		if (length - prefixCount) > 0 {
			ts.pages[ts.page].lines[ts.pages[ts.page].line].text = ts.pages[ts.page].lines[ts.pages[ts.page].line].text[0 : length-1]
			ts.pages[ts.page].lines[ts.pages[ts.page].line].chars[length-1].Remove(ts.world)
			ts.pages[ts.page].lines[ts.pages[ts.page].line].chars = ts.pages[ts.page].lines[ts.pages[ts.page].line].chars[0 : length-1]
		}
	case engo.KeyArrowUp:
		// yoffset := float32(ts.pages[ts.page].line * int(16))
		// if yoffset > 704 && ts.pages[ts.page].enil > 0 {
		// 	ts.pages[ts.page].pushScreenDown()
		// }
	case engo.KeyArrowDown:
		// yoffset := float32(ts.pages[ts.page].line * 16)
		// if yoffset > 704 {
		// 	ts.pages[ts.page].pushScreenUp()
		// }
	case engo.KeyEscape:
		if ts.pages[ts.page].escapable {
			ts.pages[ts.page].hide()
			delete(ts.pages, ts.page)
			ts.page--
			ts.pages[ts.page].show()

			ts.pages[ts.page].lines[ts.pages[ts.page].line] = &line{}
			ts.pages[ts.page].lines[ts.pages[ts.page].line].prefix(ts.delegateKeyPress)
		}
	case engo.KeyEnter:
		if ts.pages[ts.page].readonly && !mods.Output {
			break
		}

		ts.pages[ts.page].lines[ts.pages[ts.page].line].locked = true

		xoffset := ts.getXoffset()
		if xoffset > 710 {
			ts.pages[ts.page].line += int(math.Floor(float64(xoffset)/710)) + 1
		} else {
			ts.pages[ts.page].line++
		}

		yoffset := float32(ts.pages[ts.page].line * 16)
		if yoffset > 704 {
			ts.pages[ts.page].pushScreenUp()
		}

		if !mods.Ignore {
			ts.pages[ts.page].lines[ts.pages[ts.page].line-1].evaluate(ts)
		}

		// Add a new line after everything
		ts.pages[ts.page].lines[ts.pages[ts.page].line] = &line{}
		if !mods.Ignore {
			ts.pages[ts.page].lines[ts.pages[ts.page].line].prefix(ts.delegateKeyPress)
		}
	default:
		if ts.pages[ts.page].readonly && !mods.Output {
			break
		}

		symbol := ""
		if mods.Line == nil {
			symbol = input.KeyToString(key, mods)
		} else {
			symbol = *mods.Line
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

func (ts *TerminalSystem) WriteLine(str string) {
	ts.pages[ts.page].lines[ts.pages[ts.page].line] = &line{}

	line := ""
	for _, char := range strings.Split(str, "") {
		if char == "\t" {
			char = "    "
		}

		line += char
	}

	ts.delegateKeyPress(engo.Key(-1), &input.Modifiers{Output: true, Line: &line})
	ts.delegateKeyPress(engo.KeyEnter, &input.Modifiers{Ignore: true, Output: true})
}

func (ts *TerminalSystem) WriteError(err error) {
	for _, line := range strings.Split(err.Error(), "\n") {
		ts.WriteLine(line)
	}
}
