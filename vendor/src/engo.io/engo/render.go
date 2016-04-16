package engo

import (
	"fmt"
	"image/color"
	"math"
	"sort"
	"strings"

	"engo.io/ecs"
	"engo.io/webgl"
)

const (
	RenderSystemPriority = -1000
)

type renderChangeMessage struct{}

func (renderChangeMessage) Type() string {
	return "renderChangeMessage"
}

type Drawable interface {
	Texture() *webgl.Texture
	Width() float32
	Height() float32
	View() (float32, float32, float32, float32)
}

type RenderComponent struct {
	scale        Point
	Label        string
	Transparency float32
	Color        color.Color
	shader       Shader
	zIndex       float32

	drawable      Drawable
	buffer        *webgl.Buffer
	bufferContent []float32
}

func NewRenderComponent(d Drawable, scale Point, label string) *RenderComponent {
	rc := &RenderComponent{
		Label:        label,
		Transparency: 1,
		Color:        color.White,

		scale: scale,
	}
	rc.SetDrawable(d)

	return rc
}

func (r *RenderComponent) SetDrawable(d Drawable) {
	r.drawable = d
	r.preloadTexture()
}

func (r *RenderComponent) Drawable() Drawable {
	return r.drawable
}

func (r *RenderComponent) SetScale(scale Point) {
	r.scale = scale
	r.preloadTexture()
}

func (r *RenderComponent) Scale() Point {
	return r.scale
}

func (r *RenderComponent) SetShader(s Shader) {
	r.shader = s
	Mailbox.Dispatch(&renderChangeMessage{})
}

func (r *RenderComponent) SetZIndex(index float32) {
	r.zIndex = index
	Mailbox.Dispatch(&renderChangeMessage{})
}

func (*RenderComponent) Type() string {
	return "RenderComponent"
}

// Init is called to initialize the RenderElement
func (ren *RenderComponent) preloadTexture() {
	if ren.drawable == nil || headless {
		return
	}

	ren.bufferContent = ren.generateBufferContent()

	ren.buffer = Gl.CreateBuffer()
	Gl.BindBuffer(Gl.ARRAY_BUFFER, ren.buffer)
	Gl.BufferData(Gl.ARRAY_BUFFER, ren.bufferContent, Gl.STATIC_DRAW)

	// TODO: ask why this doesn't work
	// ren.bufferContent = make([]float32, 0)
}

// generateBufferContent computes information about the 4 vertices needed to draw the texture, which should
// be stored in the buffer
func (ren *RenderComponent) generateBufferContent() []float32 {
	scaleX := ren.scale.X
	scaleY := ren.scale.Y
	rotation := float32(0.0)
	transparency := float32(1.0)
	c := ren.Color

	fx := float32(0)
	fy := float32(0)
	fx2 := ren.drawable.Width()
	fy2 := ren.drawable.Height()

	if scaleX != 1 || scaleY != 1 {
		//fx *= scaleX
		//fy *= scaleY
		fx2 *= scaleX
		fy2 *= scaleY
	}

	p1x := fx
	p1y := fy
	p2x := fx
	p2y := fy2
	p3x := fx2
	p3y := fy2
	p4x := fx2
	p4y := fy

	var x1 float32
	var y1 float32
	var x2 float32
	var y2 float32
	var x3 float32
	var y3 float32
	var x4 float32
	var y4 float32

	if rotation != 0 {
		rot := float64(rotation * (math.Pi / 180.0))

		cos := float32(math.Cos(rot))
		sin := float32(math.Sin(rot))

		x1 = cos*p1x - sin*p1y
		y1 = sin*p1x + cos*p1y

		x2 = cos*p2x - sin*p2y
		y2 = sin*p2x + cos*p2y

		x3 = cos*p3x - sin*p3y
		y3 = sin*p3x + cos*p3y

		x4 = x1 + (x3 - x2)
		y4 = y3 - (y2 - y1)
	} else {
		x1 = p1x
		y1 = p1y

		x2 = p2x
		y2 = p2y

		x3 = p3x
		y3 = p3y

		x4 = p4x
		y4 = p4y
	}

	colorR, colorG, colorB, _ := c.RGBA()

	red := colorR
	green := colorG << 8
	blue := colorB << 16
	alpha := uint32(transparency*255.0) << 24

	tint := math.Float32frombits((alpha | blue | green | red) & 0xfeffffff)

	u, v, u2, v2 := ren.drawable.View()

	return []float32{x1, y1, u, v, tint, x4, y4, u2, v, tint, x3, y3, u2, v2, tint, x2, y2, u, v2, tint}
}

type renderEntityList []*ecs.Entity

func (r renderEntityList) Len() int {
	return len(r)
}

func (r renderEntityList) Less(i, j int) bool {
	var (
		rc1 *RenderComponent
		rc2 *RenderComponent
		ok  bool
	)
	if rc1, ok = r[i].ComponentFast(rc1).(*RenderComponent); !ok {
		return false // those without render component go last
	}
	if rc2, ok = r[i].ComponentFast(rc1).(*RenderComponent); !ok {
		return true // those without render component go last
	}

	// Sort by shader-pointer if they have the same zIndex
	if rc1.zIndex == rc2.zIndex {
		// TODO: optimize this for performance
		return strings.Compare(fmt.Sprintf("%p", rc1.shader), fmt.Sprintf("%p", rc2.shader)) < 0
	}

	return rc1.zIndex < rc2.zIndex
}

func (r renderEntityList) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type RenderSystem struct {
	renders renderEntityList
	world   *ecs.World

	sortingNeeded bool
	currentShader Shader
}

func (rs *RenderSystem) New(w *ecs.World) {
	rs.world = w

	if !headless {
		initShaders(Width(), Height())
	}

	Mailbox.Listen("renderChangeMessage", func(Message) {
		rs.sortingNeeded = true
	})
}

func (rs *RenderSystem) AddEntity(e *ecs.Entity) {
	rs.renders = append(rs.renders, e)
	rs.sortingNeeded = true
}

func (rs *RenderSystem) RemoveEntity(e *ecs.Entity) {
	var removeIndex int = -1
	for index, entity := range rs.renders {
		if entity.ID() == e.ID() {
			removeIndex = index
			break
		}
	}
	if removeIndex >= 0 {
		rs.renders = append(rs.renders[:removeIndex], rs.renders[removeIndex+1:]...) // TODO: test for edge cases
		rs.sortingNeeded = true
	}
}

func (rs *RenderSystem) Update(dt float32) {
	if headless {
		return
	}

	if rs.sortingNeeded {
		sort.Sort(rs.renders)
		rs.sortingNeeded = false
	}

	Gl.Clear(Gl.COLOR_BUFFER_BIT)

	// TODO: it's linear for now, but that might very well be a bad idea
	for _, entity := range rs.renders {
		var (
			render *RenderComponent
			space  *SpaceComponent
			ok     bool
		)

		if render, ok = entity.ComponentFast(render).(*RenderComponent); !ok {
			continue // with other entities
		}

		if space, ok = entity.ComponentFast(space).(*SpaceComponent); !ok {
			continue // with other entities
		}

		// Retrieve a shader, may be the default one -- then use it if we aren't already using it
		shader := render.shader
		if shader == nil {
			shader = DefaultShader
		}

		// Change Shader if we have to
		if shader != rs.currentShader {
			if rs.currentShader != nil {
				rs.currentShader.Post()
			}
			shader.Pre()
			rs.currentShader = shader
		}

		rs.currentShader.Draw(render.drawable.Texture(), render.buffer, space.Position.X, space.Position.Y, 0) // TODO: add rotation
	}

	if rs.currentShader != nil {
		rs.currentShader.Post()
		rs.currentShader = nil
	}
}

func (*RenderSystem) Type() string {
	return "RenderSystem"
}

func (*RenderSystem) Priority() int {
	return RenderSystemPriority
}
