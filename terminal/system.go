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

const (
	// FontSize is the size of the font used in the terminal
	FontSize float32 = 16
)

// System is a scrollable, visual and text input-able system.
type System struct {
	Map *navigator.Map

	pages map[int]*page
	page  int

	world *ecs.World
	vfs   *filesystem.VirtualFS

	ship *ship.Ship

	needsDraw []string

	boundKeys map[engo.Key]func() bool
}

// Viewer is the background entity for the terminal.
type Viewer struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// Remove is called whenever an Entity is removed from the World, in order to
// remove it from this sytem as well.
func (*System) Remove(ecs.BasicEntity) {
	// TODO
}

// Update is ran every frame, with `dt` being the time in seconds since the
// last frame.
func (ts *System) Update(dt float32) {
	for i := 0; i < 5; i++ {
		if len(ts.needsDraw) < 1 {
			break
		}

		char := ""
		char, ts.needsDraw = ts.needsDraw[0], ts.needsDraw[1:]

		switch char {
		case "\n":
			ts.delegateKeyPress(engo.KeyEnter, &input.Modifiers{Ignore: true})
		default:
			ts.delegateKeyPress(input.StringToKey(char))
		}
	}
}

// New is the initialisation of the System.
func (ts *System) New(w *ecs.World) {
	ts.vfs = filesystem.New(ts.WriteError)
	ts.world = w
	ts.pages = make(map[int]*page)
	ts.pages[ts.page] = &page{
		lines:    make(map[int]*line),
		line:     0,
		readonly: true,
	}

	ts.registerKeys()
	ts.boundKeys = make(map[engo.Key]func() bool)

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

	log.Println("System initialized")
}

func (ts *System) addBackground(w *ecs.World) {
	bkg := &Viewer{BasicEntity: ecs.NewBasic()}

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

func (ts *System) registerKeys() {
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
				engo.KeyArrowLeft,
				engo.KeyArrowRight,
			},
			OnPress: ts.delegateKeyPress,
		},
	})
}

func (ts *System) bind(key engo.Key, cb func() bool) {
	ts.boundKeys[key] = cb
}

func (ts *System) unbind(key engo.Key) {
	delete(ts.boundKeys, key)
}

func (ts *System) delegateKeyPress(key engo.Key, mods *input.Modifiers) {
	if ts.pages[ts.page] == nil {
		ts.pages[ts.page] = &page{
			lines: make(map[int]*line),
			line:  0,
		}
	}

	if ts.pages[ts.page].readonly {
		if ts.pages[ts.page].cursor != nil {
			ts.pages[ts.page].cursor.Remove(ts.world)
			ts.pages[ts.page].cursor = nil
		}
	}

	if ts.pages[ts.page].lines[ts.pages[ts.page].line] == nil {
		ts.pages[ts.page].lines[ts.pages[ts.page].line] = &line{}

		if !ts.pages[ts.page].editable {
			ts.pages[ts.page].lines[ts.pages[ts.page].line].prefix(ts.delegateKeyPress)
		}
	}

	// Check to see if a script set a binding on a key.
	if cb, ok := ts.boundKeys[key]; ok {
		// If so, run the callback function. If it results in a false value
		// then we do not want to continue. If it results in a true value then
		// we continue with the normal mappings.
		if !cb() {
			return
		}
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

		if (length-prefixCount) > 0 && (length-prefixCount) > ts.pages[ts.page].cpoint {
			remove := (length - ts.pages[ts.page].cpoint - 1)
			ts.pages[ts.page].lines[ts.pages[ts.page].line].chars[remove].Remove(ts.world)

			ts.pages[ts.page].lines[ts.pages[ts.page].line].text = append(
				ts.pages[ts.page].lines[ts.pages[ts.page].line].text[:remove],
				ts.pages[ts.page].lines[ts.pages[ts.page].line].text[remove+1:]...,
			)

			ts.pages[ts.page].lines[ts.pages[ts.page].line].chars = append(
				ts.pages[ts.page].lines[ts.pages[ts.page].line].chars[:remove],
				ts.pages[ts.page].lines[ts.pages[ts.page].line].chars[remove+1:]...,
			)

			// Redraw entire line
			if remove != length-1 {
				for _, char := range ts.pages[ts.page].lines[ts.pages[ts.page].line].chars {
					char.Remove(ts.world)
				}

				ts.pages[ts.page].lines[ts.pages[ts.page].line].chars = nil
				line := ts.pages[ts.page].lines[ts.pages[ts.page].line].text
				ts.pages[ts.page].lines[ts.pages[ts.page].line].text = []string{}

				for _, char := range line {
					ts.delegateKeyPress(input.StringToKey(char, &input.Modifiers{Redraw: true}))
				}
			}
		}
	case engo.KeyArrowLeft:
		min := len(ts.pages[ts.page].lines[ts.pages[ts.page].line].chars) - ts.pages[ts.page].lines[ts.pages[ts.page].line].prefixCount
		if min > ts.pages[ts.page].cpoint {
			ts.pages[ts.page].cpoint++
		}
	case engo.KeyArrowRight:
		if ts.pages[ts.page].cpoint > 0 {
			ts.pages[ts.page].cpoint--
		}
	case engo.KeyArrowUp:
		if ts.pages[ts.page].editable {
			if ts.pages[ts.page].lines[ts.pages[ts.page].line-1] == nil {
				break
			}

			if ts.pages[ts.page].cpoint > 0 {
				ts.pages[ts.page].cpoint = 0
			}

			ts.pages[ts.page].line--
			ts.pages[ts.page].lines[ts.pages[ts.page].line].locked = false

			if ts.pages[ts.page].enil > 0 && ts.pages[ts.page].line < ts.pages[ts.page].enil {
				ts.pages[ts.page].pushScreenDown()
			}
		} else {
			if len(ts.pages[ts.page].commands) <= ts.pages[ts.page].cmdindex {
				break
			}

			ts.pages[ts.page].cmdindex++
			cmd := ts.pages[ts.page].commands[len(ts.pages[ts.page].commands)-(ts.pages[ts.page].cmdindex)]

			if (length - prefixCount) > 0 {
				for i := length - 1; i >= prefixCount; i-- {
					ts.pages[ts.page].lines[ts.pages[ts.page].line].text = ts.pages[ts.page].lines[ts.pages[ts.page].line].text[0:i]
					ts.pages[ts.page].lines[ts.pages[ts.page].line].chars[i].Remove(ts.world)
					ts.pages[ts.page].lines[ts.pages[ts.page].line].chars = ts.pages[ts.page].lines[ts.pages[ts.page].line].chars[0:i]
				}
			}

			for _, char := range cmd {
				ts.delegateKeyPress(input.StringToKey(string(char)))
			}
		}
	case engo.KeyArrowDown:
		if ts.pages[ts.page].editable {
			if ts.pages[ts.page].lines[ts.pages[ts.page].line+1] == nil {
				break
			}

			if ts.pages[ts.page].cpoint > 0 {
				ts.pages[ts.page].cpoint = 0
			}

			ts.pages[ts.page].line++
			ts.pages[ts.page].lines[ts.pages[ts.page].line].locked = false

			yoffset := float32(ts.pages[ts.page].line-ts.pages[ts.page].enil) * FontSize
			if yoffset > 704 {
				ts.pages[ts.page].pushScreenUp()
			}
		} else {
			if ts.pages[ts.page].cmdindex < 1 {
				break
			}

			ts.pages[ts.page].cmdindex--

			var cmd string
			if ts.pages[ts.page].cmdindex > 1 {
				cmd = ts.pages[ts.page].commands[len(ts.pages[ts.page].commands)-(ts.pages[ts.page].cmdindex)]
			}

			if (length - prefixCount) > 0 {
				for i := length - 1; i >= prefixCount; i-- {
					ts.pages[ts.page].lines[ts.pages[ts.page].line].text = ts.pages[ts.page].lines[ts.pages[ts.page].line].text[0:i]
					ts.pages[ts.page].lines[ts.pages[ts.page].line].chars[i].Remove(ts.world)
					ts.pages[ts.page].lines[ts.pages[ts.page].line].chars = ts.pages[ts.page].lines[ts.pages[ts.page].line].chars[0:i]
				}
			}

			for _, char := range cmd {
				ts.delegateKeyPress(input.StringToKey(string(char)))
			}
		}
	case engo.KeyEscape:
		if ts.pages[ts.page].escapable {
			if !ts.pages[ts.page].editable {
				ts.pages[ts.page].hide()
				delete(ts.pages, ts.page)
				ts.page--
				ts.pages[ts.page].show()

				ts.pages[ts.page].lines[ts.pages[ts.page].line] = &line{}
				ts.pages[ts.page].lines[ts.pages[ts.page].line].prefix(ts.delegateKeyPress)
			}
		}
	case engo.KeyEnter:
		ts.pages[ts.page].cmdindex = 0
		ts.pages[ts.page].cpoint = 0

		if ts.pages[ts.page].readonly && !mods.Output {
			break
		}

		if !ts.pages[ts.page].editable {
			ts.pages[ts.page].lines[ts.pages[ts.page].line].locked = true
		}

		xoffset := ts.getXoffset()
		if xoffset > 710 {
			ts.pages[ts.page].line += int(math.Floor(float64(xoffset)/710)) + 1
		} else {
			ts.pages[ts.page].line++
		}

		lastyoffset := float32(0)
		yoffset := float32(ts.pages[ts.page].line) * FontSize
		if ts.pages[ts.page].line < len(ts.pages[ts.page].lines) {
			lastyoffset = float32(len(ts.pages[ts.page].lines)-1) * FontSize
		}

		if yoffset > 704 || lastyoffset > 704 {
			ts.pages[ts.page].pushScreenUp()
		}

		if !mods.Ignore && !ts.pages[ts.page].editable {
			ts.pages[ts.page].commands = append(ts.pages[ts.page].commands, ts.pages[ts.page].lines[ts.pages[ts.page].line-1].String())
			ts.pages[ts.page].lines[ts.pages[ts.page].line-1].evaluate(ts)
		}

		if ts.pages[ts.page].line < len(ts.pages[ts.page].lines) {
			for i := len(ts.pages[ts.page].lines); i > ts.pages[ts.page].line; i-- {
				// fmt.Println(len(ts.pages[ts.page].lines), i, ts.pages[ts.page].line)

				ts.pages[ts.page].lines[i] = ts.pages[ts.page].lines[i-1]
				for _, char := range ts.pages[ts.page].lines[i].chars {
					char.SetY(char.Y + FontSize).Render()
				}
			}
		}

		// Add a new line
		ts.pages[ts.page].lines[ts.pages[ts.page].line] = &line{}
		if !mods.Ignore && !ts.pages[ts.page].editable {
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

		redraw := false
		if ts.pages[ts.page].cpoint > 0 && !mods.Redraw {
			redraw = true
			pos := len(ts.pages[ts.page].lines[ts.pages[ts.page].line].text) - ts.pages[ts.page].cpoint
			ts.pages[ts.page].lines[ts.pages[ts.page].line].text = append(
				ts.pages[ts.page].lines[ts.pages[ts.page].line].text[:pos],
				append([]string{symbol}, ts.pages[ts.page].lines[ts.pages[ts.page].line].text[pos:]...)...,
			)
		} else {
			ts.pages[ts.page].lines[ts.pages[ts.page].line].text = append(ts.pages[ts.page].lines[ts.pages[ts.page].line].text, symbol)
		}

		if !redraw {
			char := ui.NewText(symbol)

			push := false
			var xoffset, yoffset float32
			xoffset = ts.getXoffset()
			yoffset = float32(ts.pages[ts.page].lineOffset()) * FontSize

			if xoffset >= 710 {
				lines := int(math.Floor(float64(xoffset) / 710))

				xoffset = xoffset - float32(707*lines)
				yoffset += FontSize * float32(lines)

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
		} else {
			for _, char := range ts.pages[ts.page].lines[ts.pages[ts.page].line].chars {
				char.Remove(ts.world)
			}

			ts.pages[ts.page].lines[ts.pages[ts.page].line].chars = nil
			line := ts.pages[ts.page].lines[ts.pages[ts.page].line].text
			ts.pages[ts.page].lines[ts.pages[ts.page].line].text = []string{}

			for _, char := range line {
				ts.delegateKeyPress(input.StringToKey(char, &input.Modifiers{Redraw: true}))
			}
		}
	}

	if !ts.pages[ts.page].readonly {
		if ts.pages[ts.page].cursor == nil {
			ts.pages[ts.page].cursor = ui.NewText("_")
			ts.pages[ts.page].cursor.Insert(ts.world)
		}

		var xoffset, yoffset float32
		xoffset = ts.getXoffset()
		yoffset = float32(ts.pages[ts.page].lineOffset()) * FontSize

		if xoffset >= 710 {
			lines := int(math.Floor(float64(xoffset) / 710))

			xoffset = xoffset - float32(707*lines)
			yoffset += float32(FontSize) * float32(lines)
		}

		ts.pages[ts.page].cursor.SetX(xoffset + (35 + (FontSize * .65) - (float32(ts.pages[ts.page].cpoint)*FontSize)*.65))
		ts.pages[ts.page].cursor.SetY(yoffset + 38)
	}

}

func (ts *System) getXoffset() float32 {
	if ts.pages[ts.page].lines[ts.pages[ts.page].line] == nil {
		return 0
	}

	return float32(len(ts.pages[ts.page].lines[ts.pages[ts.page].line].text)*int(FontSize)) * .65
}

// WriteLine takes a string, builds characters, and writes it to the terminal.
func (ts *System) WriteLine(str string) {
	ts.pages[ts.page].lines[ts.pages[ts.page].line] = &line{}

	line := ""
	for _, char := range strings.Split(str, "") {
		if char == "\t" {
			char = "    "
		}

		line += char

		if ts.pages[ts.page].editable {
			ts.needsDraw = append(ts.needsDraw, char)
		}
	}

	if !ts.pages[ts.page].editable {
		ts.delegateKeyPress(engo.Key(-1), &input.Modifiers{Output: true, Line: &line})
		ts.delegateKeyPress(engo.KeyEnter, &input.Modifiers{Ignore: true, Output: true})
	} else {
		ts.needsDraw = append(ts.needsDraw, "\n")
	}
}

// WriteError takes an error and uses WriteLine to print the error. This
// supports multi-line errors.
func (ts *System) WriteError(err error) {
	for _, line := range strings.Split(err.Error(), "\n") {
		ts.WriteLine(line)
	}
}
