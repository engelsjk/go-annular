package goannular

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

type Annular struct {
	title           string
	width           float64
	height          float64
	maxRadialCenter float64
	maxArcLength    float64
	maxRadialLength float64
	maxN            int
	colors          *Colors
	canvas          *canvas.Canvas
	ctx             *canvas.Context
}

type AnnularOption func(*Annular)

func WithWidth(width float64) AnnularOption {
	return func(a *Annular) {
		a.width = width
	}
}

func WithHeight(height float64) AnnularOption {
	return func(a *Annular) {
		a.height = height
	}
}

func NewAnnular(opts ...AnnularOption) *Annular {

	const (
		defaultWidth           = 1000.0
		defaultHeight          = 1000.0
		defaultMaxRadialCenter = 0.10
		defaultMaxArcLength    = 0.05
		defaultMaxRadialLength = 0.05
		defaultMaxN            = 15000
	)

	a := &Annular{
		width:           defaultWidth,
		height:          defaultHeight,
		maxRadialCenter: defaultMaxRadialCenter,
		maxArcLength:    defaultMaxArcLength,
		maxRadialLength: defaultMaxRadialLength,
		maxN:            defaultMaxN,
	}

	for _, opt := range opts {
		opt(a)
	}

	seed := time.Now().Unix()
	a.title = strconv.FormatInt(seed, 10)

	rand.Seed(seed)

	a.colors = &Colors{palettes: palettes, numPalettes: len(palettes)}

	return a
}

func (a *Annular) Draw() {

	a.colors.SetRandomPalette()

	a.canvas = canvas.New(a.width, a.height)
	a.ctx = canvas.NewContext(a.canvas)

	a.ctx.SetFillColor(a.colors.RandomColorOrBlackRGBA())
	a.ctx.DrawPath(0, 0, canvas.Rectangle(a.canvas.W, a.canvas.H))

	// randomize parameters
	radialCenter := rand.Float64() * a.maxRadialCenter * float64(a.width) //px
	cx, cy := rand.Float64()*float64(a.width), rand.Float64()*float64(a.height)
	maxMaxArcLength := rand.Float64() * a.maxArcLength
	maxMaxRadialLength := rand.Float64() * a.maxRadialLength
	stype := rand.Intn(4)
	n := rand.Intn(a.maxN)

	// annuli
	for i := 0; i < n; i++ {

		arcStart := math.Mod(rand.Float64()*360.0/180.0*math.Pi, 2*math.Pi)
		radialStart := radialCenter + rand.Float64()*(math.Sqrt(2)*float64(a.width))

		radialLength := radialLength(
			radialStart,
			maxMaxRadialLength*float64(a.width),
			arcStart,
			stype,
		)

		arcLength := arcLength(
			arcStart,
			maxMaxArcLength*float64(a.width),
			radialStart,
			stype,
		)

		arcEnd := math.Mod(arcStart+arcLength/180.0*math.Pi, 2*math.Pi)
		radialEnd := radialStart + radialLength

		annulus := Annulus{
			x:     cx,
			y:     cy,
			start: arcStart,
			end:   arcEnd,
			inner: radialStart,
			outer: radialEnd,
		}

		svg := annulus.SVG()
		path := canvas.MustParseSVG(svg)
		a.ctx.SetFillColor(a.colors.RandomColorRGBA())
		a.ctx.DrawPath(0, 0, path)
	}

}

func (a *Annular) SVG(w io.Writer) error {
	return a.Render(w, "svg")
}

func (a *Annular) PNG(w io.Writer) error {
	return a.Render(w, "png")
}

func (a *Annular) Render(w io.Writer, format string) error {

	if a.canvas == nil {
		return fmt.Errorf("no canvas drawn yet")
	}

	switch format {
	case "png":
		cw := renderers.PNG()
		cw(w, a.canvas)
		return nil
	case "svg":
		cw := renderers.SVG()
		cw(w, a.canvas)
		return nil
	default:
		return fmt.Errorf("format not recognized")
	}
}
