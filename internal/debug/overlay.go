// Package debug provides a lightweight, asset free on screen overlay for
// inspecting the game state at runtime.
//
// The overlay knows nothing about the game itself. A scene feeds it labelled
// values every frame and the overlay takes care of layout, the background panel
// and the visibility toggle. That keeps this package free of imports from
// internal/game, internal/player or internal/world, so it can never cause an
// import cycle.
//
// Typical usage inside a Draw method:
//
//	if o.Visible() {
//	    o.Reset()
//	    o.Section("PLAYER")
//	    o.Vec2("Pos", p.Pos)
//	    o.Draw(screen)
//	}
package debug

import (
	"bytes"
	"image/color"
	"strconv"

	"github.com/Driemtax/Byteborn/pkg/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Metrics of the bitmap font used by ebitenutil.DebugPrintAt. They are hard
// coded constants inside ebitenutil.drawDebugText, so we mirror them here to be
// able to size the background panel correctly.
//
// Note that this font only covers U+0000 to U+00FF, so stick to plain ASCII.
const (
	charWidth  = 6
	lineHeight = 16

	// DebugPrintAt translates the text by one extra pixel on the x axis.
	// We compensate for it so the panel padding stays symmetric.
	debugPrintXOffset = 1
)

const (
	panelPadX = 4
	panelPadY = 2

	// Labels are padded to this column so that all values line up.
	labelWidth = 9
)

var panelColor = color.RGBA{R: 0, G: 0, B: 0, A: 170}

// Overlay collects labelled values and renders them as a text panel in the top
// left corner of the screen.
//
// The zero value is not ready to use, call NewOverlay instead.
type Overlay struct {
	visible bool

	// buf holds the fully rendered text of the current frame. bytes.Buffer
	// keeps its backing array when Reset is called, so after the first few
	// frames building the text does not allocate anymore. The only remaining
	// allocation per frame is the buf.String() call in Draw.
	buf bytes.Buffer

	// scratch is reused by the strconv.Append based helpers. Using them instead
	// of fmt.Sprintf avoids reflection and keeps the overlay cheap enough to
	// leave running permanently.
	scratch []byte

	lines      int // number of lines written so far
	curLineLen int // length of the line currently being written, in characters
	maxLineLen int // longest completed line, in characters
}

func NewOverlay() *Overlay {
	o := &Overlay{
		scratch: make([]byte, 0, 32),
	}
	o.buf.Grow(512)
	return o
}

// Toggle flips the visibility of the overlay.
func (o *Overlay) Toggle() { o.visible = !o.visible }

// Visible reports whether the overlay is currently shown. Guard the whole value
// collection with this so a hidden overlay costs nothing at all.
func (o *Overlay) Visible() bool { return o.visible }

// SetVisible forces a specific visibility, e.g. to show the overlay on startup.
func (o *Overlay) SetVisible(visible bool) { o.visible = visible }

// Reset drops all collected lines. Call it once before feeding a new frame.
func (o *Overlay) Reset() {
	o.buf.Reset()
	o.lines = 0
	o.curLineLen = 0
	o.maxLineLen = 0
}

// --- value collection ---

// Section starts a new titled group, separated by a blank line.
func (o *Overlay) Section(title string) {
	if o.lines > 0 {
		o.startLine() // blank separator line
	}
	o.startLine()
	o.writeString("-- ")
	o.writeString(title)
	o.writeString(" --")
}

// Text adds a label with a plain string value.
func (o *Overlay) Text(label, value string) {
	o.startLine()
	o.writeLabel(label)
	o.writeString(value)
}

// Int adds a label with an integer value.
func (o *Overlay) Int(label string, value int) {
	o.startLine()
	o.writeLabel(label)
	o.writeInt(int64(value))
}

// Int64 adds a label with a 64 bit integer value, e.g. ebiten.Tick().
func (o *Overlay) Int64(label string, value int64) {
	o.startLine()
	o.writeLabel(label)
	o.writeInt(value)
}

// Float adds a label with a floating point value rounded to decimals places.
func (o *Overlay) Float(label string, value float64, decimals int) {
	o.startLine()
	o.writeLabel(label)
	o.writeFloat(value, decimals)
}

// Bool adds a label with a boolean value.
func (o *Overlay) Bool(label string, value bool) {
	o.startLine()
	o.writeLabel(label)
	if value {
		o.writeString("yes")
	} else {
		o.writeString("no")
	}
}

// Vec2 adds a label with both components of a vector.
func (o *Overlay) Vec2(label string, value types.Vec2) {
	o.startLine()
	o.writeLabel(label)
	o.writeFloat(value.X, 1)
	o.writeString(" | ")
	o.writeFloat(value.Y, 1)
}

// --- rendering ---

func (o *Overlay) Draw(screen *ebiten.Image) {
	if !o.visible || o.lines == 0 {
		return
	}

	// The line currently in progress was never committed by startLine, so take
	// it into account here.
	cols := o.maxLineLen
	if o.curLineLen > cols {
		cols = o.curLineLen
	}

	panelW := float32(cols*charWidth + 2*panelPadX)
	panelH := float32(o.lines*lineHeight + 2*panelPadY)
	vector.FillRect(screen, 0, 0, panelW, panelH, panelColor, false)

	ebitenutil.DebugPrintAt(screen, o.buf.String(), panelPadX-debugPrintXOffset, panelPadY)
}

// --- internal helpers ---

// startLine commits the width of the previous line and begins a new one.
func (o *Overlay) startLine() {
	if o.lines > 0 {
		o.buf.WriteByte('\n')
	}
	if o.curLineLen > o.maxLineLen {
		o.maxLineLen = o.curLineLen
	}
	o.curLineLen = 0
	o.lines++
}

// writeLabel writes "label:" and pads it out to labelWidth so that the values
// of every row start in the same column. Because a label is always written at
// the very beginning of a line, curLineLen is exactly the current column.
func (o *Overlay) writeLabel(label string) {
	o.writeString(label)
	o.writeByte(':')
	o.writeByte(' ')
	for o.curLineLen < labelWidth {
		o.writeByte(' ')
	}
}

func (o *Overlay) writeString(s string) {
	o.buf.WriteString(s)
	o.curLineLen += len(s)
}

func (o *Overlay) writeByte(b byte) {
	o.buf.WriteByte(b)
	o.curLineLen++
}

func (o *Overlay) writeInt(v int64) {
	o.scratch = strconv.AppendInt(o.scratch[:0], v, 10)
	o.buf.Write(o.scratch)
	o.curLineLen += len(o.scratch)
}

func (o *Overlay) writeFloat(v float64, decimals int) {
	o.scratch = strconv.AppendFloat(o.scratch[:0], v, 'f', decimals, 64)
	o.buf.Write(o.scratch)
	o.curLineLen += len(o.scratch)
}
