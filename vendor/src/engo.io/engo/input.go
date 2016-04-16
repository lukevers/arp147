package engo

type Action int
type Key int
type Modifier int

var (
	Mouse mouse

	MOVE    = Action(0)
	PRESS   = Action(1)
	RELEASE = Action(2)
	NEUTRAL = Action(99)
	SHIFT   = Modifier(0x0001)
	CONTROL = Modifier(0x0002)
	ALT     = Modifier(0x0004)
	SUPER   = Modifier(0x0008)
)

// MouseButton corresponds to a mouse button.
type MouseButton int

// Mouse buttons
const (
	MouseButton1      MouseButton = 0 // left button
	MouseButton2      MouseButton = 1 // right button
	MouseButton3      MouseButton = 2 // middle button
	MouseButton4      MouseButton = 3
	MouseButton5      MouseButton = 4
	MouseButton6      MouseButton = 5
	MouseButton7      MouseButton = 6
	MouseButton8      MouseButton = 7
	MouseButtonLast   MouseButton = 7
	MouseButtonLeft   MouseButton = 0 // equivalent for MouseButton1
	MouseButtonRight  MouseButton = 1 // equivalent for MouseButton2
	MouseButtonMiddle MouseButton = 2 // equivalent for MouseButton3
)

// those are default values for engo_js defined here because some of them are shared
// with engo_glfw.
// engo_glfw redefines the variables it needs to other values during init() so
var (
	Dash         = Key(189)
	Apostrophe   = Key(222)
	Semicolon    = Key(186)
	Equals       = Key(187)
	Comma        = Key(188)
	Period       = Key(190)
	Slash        = Key(191)
	Backslash    = Key(220)
	Backspace    = Key(8)
	Tab          = Key(9)
	CapsLock     = Key(20)
	Space        = Key(32)
	Enter        = Key(13)
	Escape       = Key(27)
	Insert       = Key(45)
	PrintScreen  = Key(42)
	Delete       = Key(46)
	PageUp       = Key(33)
	PageDown     = Key(34)
	Home         = Key(36)
	End          = Key(35)
	Pause        = Key(19)
	ScrollLock   = Key(145)
	ArrowLeft    = Key(37)
	ArrowRight   = Key(39)
	ArrowDown    = Key(40)
	ArrowUp      = Key(38)
	LeftBracket  = Key(219)
	LeftShift    = Key(16)
	LeftControl  = Key(17)
	LeftSuper    = Key(73)
	LeftAlt      = Key(18)
	RightBracket = Key(221)
	RightShift   = Key(16)
	RightControl = Key(17)
	RightSuper   = Key(73)
	RightAlt     = Key(18)
	Zero         = Key(48)
	One          = Key(49)
	Two          = Key(50)
	Three        = Key(51)
	Four         = Key(52)
	Five         = Key(53)
	Six          = Key(54)
	Seven        = Key(55)
	Eight        = Key(56)
	Nine         = Key(57)
	F1           = Key(112)
	F2           = Key(113)
	F3           = Key(114)
	F4           = Key(115)
	F5           = Key(116)
	F6           = Key(117)
	F7           = Key(118)
	F8           = Key(119)
	F9           = Key(120)
	F10          = Key(121)
	F11          = Key(122)
	F12          = Key(123)
	A            = Key(65)
	B            = Key(66)
	C            = Key(67)
	D            = Key(68)
	E            = Key(69)
	F            = Key(70)
	G            = Key(71)
	H            = Key(72)
	I            = Key(73)
	J            = Key(74)
	K            = Key(75)
	L            = Key(76)
	M            = Key(77)
	N            = Key(78)
	O            = Key(79)
	P            = Key(80)
	Q            = Key(81)
	R            = Key(82)
	S            = Key(83)
	T            = Key(84)
	U            = Key(85)
	V            = Key(86)
	W            = Key(87)
	X            = Key(88)
	Y            = Key(89)
	Z            = Key(90)
	NumLock      = Key(144)
	NumMultiply  = Key(106)
	NumDivide    = Key(111)
	NumAdd       = Key(107)
	NumSubtract  = Key(109)
	NumZero      = Key(96)
	NumOne       = Key(97)
	NumTwo       = Key(98)
	NumThree     = Key(99)
	NumFour      = Key(100)
	NumFive      = Key(101)
	NumSix       = Key(102)
	NumSeven     = Key(103)
	NumEight     = Key(104)
	NumNine      = Key(105)
	NumDecimal   = Key(110)
	NumEnter     = Key(13)
)

type mouse struct {
	X, Y             float32
	ScrollX, ScrollY float32
	Action           Action
	Button           MouseButton
	Modifer          Modifier
}
