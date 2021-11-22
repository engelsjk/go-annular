package goannular

// https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/d
// https://stackoverflow.com/questions/11479185/svg-donut-slice-as-path-element-annular-sector

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"strconv"
	"time"

	svg "github.com/ajstarks/svgo"
)

var (
	width            = 1000
	height           = 1000
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
	colors, err := loadColorPalettes(palettesFilename)
	if err != nil {
		panic(err)
	}
	palette := colors.Palettes[rand.Intn(len(colors.Palettes))]

	// title
	title := strconv.FormatInt(seed, 10)

	// init svg
	s := svg.New(w)
	s.Start(width, height)
	s.Title(title)
	s.Rect(0, 0, width, height, s.RGB(0, 0, 0))

	// randomize parameters
	radial_center := rand.Float64() * maxRadialCenter * float64(width) //px
	cx, cy := rand.Float64()*float64(width), rand.Float64()*float64(height)
	stype := rand.Intn(4)
	n := rand.Intn(maxN)

	// annuli
	for i := 0; i < n; i++ {

		arcStart := math.Mod(rand.Float64()*360.0/180.0*math.Pi, 2*math.Pi)
		radialStart := radial_center + rand.Float64()*(math.Sqrt(2)*float64(width))

		radialLength := radialLength(radialStart, maxRadialCenter*float64(width), arcStart, stype) //px
		arcLength := arcLength(arcStart, maxArcLength*float64(width), radialStart, stype)

		arc_end := math.Mod(arcStart+arcLength/180.0*math.Pi, 2*math.Pi)
		radial_end := radialStart + radialLength

		annulus := Annulus{x: cx, y: cy, start: arcStart, end: arc_end, inner: radialStart, outer: radial_end}

		path := annulus.path()
		fill := fmt.Sprintf("fill:%s", palette[rand.Intn(len(palette))])

		s.Path(path, fill)
	}

	// end
	s.End()
}
