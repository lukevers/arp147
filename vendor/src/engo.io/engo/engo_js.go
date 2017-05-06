//+build netgo

package engo

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"engo.io/gl"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
	"honnef.co/go/js/xhr"
)

var (
	// Gl is the current OpenGL context
	Gl *gl.Context

	devicePixelRatio float64
)

func init() {
	rafPolyfill()
}

//var canvas *js.Object
var document = dom.GetWindow().Document().(dom.HTMLDocument)

// CreateWindow creates a window with the specified parameters
func CreateWindow(title string, width, height int, fullscreen bool, msaa int) {
	canvas := document.CreateElement("canvas").(*dom.HTMLCanvasElement)

	devicePixelRatio = js.Global.Get("devicePixelRatio").Float()
	canvas.Width = int(float64(width)*devicePixelRatio + 0.5)   // Nearest non-negative int.
	canvas.Height = int(float64(height)*devicePixelRatio + 0.5) // Nearest non-negative int.
	canvas.Style().SetProperty("width", fmt.Sprintf("%vpx", width), "")
	canvas.Style().SetProperty("height", fmt.Sprintf("%vpx", height), "")
	log.Println("devicePixelRatio", devicePixelRatio)

	if document.Body() == nil {
		js.Global.Get("document").Set("body", js.Global.Get("document").Call("createElement", "body"))
		log.Println("Creating body, since it doesn't exist.")
	}

	document.Body().Style().SetProperty("margin", "0", "")
	document.Body().AppendChild(canvas)

	document.SetTitle(title)

	var err error

	Gl, err = gl.NewContext(canvas.Underlying(), nil) // TODO: we can add arguments here
	if err != nil {
		log.Println("Could not create context:", err)
		return
	}
	fmt.Println("Gl.Viewport(0, 0,", width, ",", height, ")")
	Gl.GetExtension("OES_texture_float")

	// DEBUG: Add framebuffer information div.
	if false {
		//canvas.Height -= 30
		text := document.CreateElement("div")
		textContent := fmt.Sprintf("%v %v (%v %v %v) @%v", dom.GetWindow().InnerWidth(), canvas.Width, float64(width)*devicePixelRatio, GameWidth(), CanvasWidth(), devicePixelRatio)
		text.SetTextContent(textContent)
		document.Body().AppendChild(text)
	}
	gameWidth = float32(width)
	gameHeight = float32(height)
	windowWidth = WindowWidth()
	windowHeight = WindowHeight()

	w := dom.GetWindow()
	w.AddEventListener("keypress", false, func(ev dom.Event) {
		// TODO: Not sure what to do here, come back
		//ke := ev.(*dom.KeyboardEvent)
		//responser.Type(rune(keyStates[Key(ke.KeyCode)]))
	})
	w.AddEventListener("keydown", false, func(ev dom.Event) {
		ke := ev.(*dom.KeyboardEvent)
		Input.keys.Set(Key(ke.KeyCode), true)
	})

	w.AddEventListener("keyup", false, func(ev dom.Event) {
		ke := ev.(*dom.KeyboardEvent)
		Input.keys.Set(Key(ke.KeyCode), false)
	})

	w.AddEventListener("mousemove", false, func(ev dom.Event) {
		mm := ev.(*dom.MouseEvent)
		Input.Mouse.X = float32(float64(mm.ClientX) * devicePixelRatio)
		Input.Mouse.Y = float32(float64(mm.ClientY) * devicePixelRatio)
		//Mouse.Action = MOVE
	})

	w.AddEventListener("mousedown", false, func(ev dom.Event) {
		mm := ev.(*dom.MouseEvent)
		Input.Mouse.X = float32(float64(mm.ClientX) * devicePixelRatio)
		Input.Mouse.Y = float32(float64(mm.ClientY) * devicePixelRatio)
		Input.Mouse.Action = Press
	})

	w.AddEventListener("mouseup", false, func(ev dom.Event) {
		mm := ev.(*dom.MouseEvent)
		Input.Mouse.X = float32(float64(mm.ClientX) * devicePixelRatio)
		Input.Mouse.Y = float32(float64(mm.ClientY) * devicePixelRatio)
		Input.Mouse.Action = Release
	})
}

func DestroyWindow() {}

// CursorPos returns the current cursor position
func CursorPos() (x, y float32) {
	return Input.Mouse.X, Input.Mouse.Y
}

// SetTitle changes the title of the page to the given string
func SetTitle(title string) {
	document.SetTitle(title)
}

// WindowSize returns the width and height of the current window
func WindowSize() (w, h int) {
	w = int(WindowWidth())
	h = int(WindowHeight())
	return
}

// WindowWidth returns the current window width
func WindowWidth() float32 {
	return float32(dom.GetWindow().InnerWidth())
}

// WindowHeight returns the current window height
func WindowHeight() float32 {
	return float32(dom.GetWindow().InnerHeight())
}

// CanvasWidth returns the current canvas width
func CanvasWidth() float32 {
	flt, err := strconv.ParseFloat(document.Body().GetElementsByTagName("canvas")[0].GetAttribute("width"), 32)
	if err != nil {
		log.Println("[ERROR] [CanvasWidth]:", err)
	}
	return float32(flt)
}

// CanvasHeight returns the current canvas height
func CanvasHeight() float32 {
	flt, err := strconv.ParseFloat(document.Body().GetElementsByTagName("canvas")[0].GetAttribute("height"), 32)
	if err != nil {
		log.Println("[ERROR] [CanvasHeight]:", err)
	}
	return float32(flt)
}

func CanvasScale() float32 {
	return CanvasWidth()/WindowWidth()
}

func toPx(n int) string {
	return strconv.FormatInt(int64(n), 10) + "px"
}

func rafPolyfill() {
	window := js.Global
	vendors := []string{"ms", "moz", "webkit", "o"}
	if window.Get("requestAnimationFrame") == nil {
		for i := 0; i < len(vendors) && window.Get("requestAnimationFrame") == nil; i++ {
			vendor := vendors[i]
			window.Set("requestAnimationFrame", window.Get(vendor+"RequestAnimationFrame"))
			window.Set("cancelAnimationFrame", window.Get(vendor+"CancelAnimationFrame"))
			if window.Get("cancelAnimationFrame") == nil {
				window.Set("cancelAnimationFrame", window.Get(vendor+"CancelRequestAnimationFrame"))
			}
		}
	}

	lastTime := 0.0
	if window.Get("requestAnimationFrame") == nil {
		window.Set("requestAnimationFrame", func(callback func(float32)) int {
			currTime := js.Global.Get("Date").New().Call("getTime").Float()
			timeToCall := math.Max(0, 16-(currTime-lastTime))
			id := window.Call("setTimeout", func() { callback(float32(currTime + timeToCall)) }, timeToCall)
			lastTime = currTime + timeToCall
			return id.Int()
		})
	}

	if window.Get("cancelAnimationFrame") == nil {
		window.Set("cancelAnimationFrame", func(id int) {
			js.Global.Get("clearTimeout").Invoke(id)
		})
	}
}

// RunIteration runs one iteration per frame
func RunIteration() {
	Time.Tick()
	Input.update()
	currentWorld.Update(Time.Delta())
	// TODO: this may not work, and sky-rocket the FPS
	//  requestAnimationFrame(func(dt float32) {
	// 	currentWorld.Update(Time.Delta())
	// 	keysUpdate()
	// 	if !headless {
	// 		// TODO: does this require !headless?
	// 		Mouse.ScrollX, Mouse.ScrollY = 0, 0
	// 	}
	// 	Time.Tick()
	// })
}

func requestAnimationFrame(callback func(float32)) int {
	//return dom.GetWindow().RequestAnimationFrame(callback)
	return js.Global.Call("requestAnimationFrame", callback).Int()
}

func cancelAnimationFrame(id int) {
	dom.GetWindow().CancelAnimationFrame(id)
}

// RunPreparation is called automatically when calling Open. It should only be called once.
func RunPreparation() {
	Time = NewClock()

	dom.GetWindow().AddEventListener("onbeforeunload", false, func(e dom.Event) {
		dom.GetWindow().Alert("You're closing")
	})
}

func runLoop(defaultScene Scene, headless bool) {
	SetScene(defaultScene, false)
	RunPreparation()
	ticker := time.NewTicker(time.Duration(int(time.Second) / opts.FPSLimit))

	// Start tick, minimize the delta
	Time.Tick()

Outer:
	for {
		select {
		case <-ticker.C:
			if closeGame {
				break Outer
			}
			RunIteration()
		case <-resetLoopTicker:
			ticker.Stop()
			ticker = time.NewTicker(time.Duration(int(time.Second) / opts.FPSLimit))
		}
	}
	ticker.Stop()
}

func openFile(url string) (io.ReadCloser, error) {
	req := xhr.NewRequest("GET", url)

	req.ResponseType = xhr.ArrayBuffer

	if err := req.Send(""); err != nil {
		return nil, err
	}

	if req.Status != http.StatusOK {
		return nil, fmt.Errorf("unable to open resource (%s), expected HTTP status %d but got %d", url, http.StatusOK, req.Status)
	}

	buffer := bytes.NewBuffer(js.Global.Get("Uint8Array").New(req.Response).Interface().([]byte))

	return noCloseReadCloser{buffer}, nil
}

type noCloseReadCloser struct {
	r io.Reader
}

func (n noCloseReadCloser) Close() error { return nil }
func (n noCloseReadCloser) Read(p []byte) (int, error) {
	return n.r.Read(p)
}

// SetCursor changes the cursor
func SetCursor(c Cursor) {
	switch c {
	case CursorNone:
		document.Body().Style().Set("cursor", "default")
	case CursorHand:
		document.Body().Style().Set("cursor", "hand")
	}
}
