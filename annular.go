package goannular

// https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/d
// https://stackoverflow.com/questions/11479185/svg-donut-slice-as-path-element-annular-sector

import (
	"io"
	"math"
	"math/rand"
	"time"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

var (
	width            = 1000.0
	height           = 1000.0
	maxRadialCenter  = 0.10
	maxArcLength     = 0.05
	maxRadialLength  = 0.05
	maxN             = 15000
	palettesFilename = "palettes.json"
)

func Run(w io.Writer) {

	// seed
	seed := time.Now().Unix()
	rand.Seed(seed) // initialize global pseudo random generator

	// colors
	var colors Colors
	err := colors.Load(palettesFilename)
	if err != nil {
		panic(err)
	}
	colors.SetRandomPalette()

	// title
	// title := strconv.FormatInt(seed, 10)

	// init svg
	c := canvas.New(width, height)
	ctx := canvas.NewContext(c)

	ctx.SetFillColor(colors.RandomColorOrBlack())
	ctx.DrawPath(0, 0, canvas.Rectangle(c.W, c.H))

	// randomize parameters
	radialCenter := rand.Float64() * maxRadialCenter * float64(width) //px
	cx, cy := rand.Float64()*float64(width), rand.Float64()*float64(height)
	maxMaxArcLength := rand.Float64() * maxArcLength
	maxMaxRadialLength := rand.Float64() * maxRadialLength
	stype := rand.Intn(4)
	n := rand.Intn(maxN)

	// annuli
	for i := 0; i < n; i++ {

		arcStart := math.Mod(rand.Float64()*360.0/180.0*math.Pi, 2*math.Pi)
		radialStart := radialCenter + rand.Float64()*(math.Sqrt(2)*float64(width))

		radialLength := radialLength(radialStart, maxMaxRadialLength*float64(width), arcStart, stype) //px
		arcLength := arcLength(arcStart, maxMaxArcLength*float64(width), radialStart, stype)

		arcEnd := math.Mod(arcStart+arcLength/180.0*math.Pi, 2*math.Pi)
		radialEnd := radialStart + radialLength

		annulus := Annulus{x: cx, y: cy, start: arcStart, end: arcEnd, inner: radialStart, outer: radialEnd}

		svg := annulus.SVG()

		path := canvas.MustParseSVG(svg)
		ctx.SetFillColor(colors.RandomColor())
		ctx.DrawPath(0, 0, path)
	}

	cw := renderers.SVG()
	cw(w, c)
}
