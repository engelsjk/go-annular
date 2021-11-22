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
	width              = 1000
	height             = 1000
	radial_center_base = 50.0
	arc_length_base    = 50.0
	radial_length_base = 50.0
	n_base             = 15000
	palettesFilename   = "palettes.json"
	verbose            = false
)

func Run(w io.Writer) {

	s := svg.New(w)

	seed := time.Now().Unix()
	rand.Seed(seed) // initialize global pseudo random generator

	title := strconv.FormatInt(seed, 10)
	s.Title(title)

	colors := loadColorPalettes(palettesFilename)
	palette := colors.Palettes[rand.Intn(len(colors.Palettes))]

	s.Start(width, height)
	s.Rect(0, 0, width, height, s.RGB(0, 0, 0))

	radial_center := rand.Float64() * radial_center_base //px
	cx, cy := rand.Float64()*float64(width), rand.Float64()*float64(height)

	stype := rand.Intn(4)

	n := rand.Intn(n_base)

	for i := 0; i < n; i++ {

		arc_start := math.Mod(rand.Float64()*360.0/180.0*math.Pi, 2*math.Pi)
		radial_start := radial_center + rand.Float64()*(math.Sqrt(2)*float64(width))

		radial_length := radialLength(radial_start, radial_length_base, arc_start, stype) //px
		arc_length := arcLength(arc_start, arc_length_base, radial_start, stype)

		arc_end := math.Mod(arc_start+arc_length/180.0*math.Pi, 2*math.Pi)
		radial_end := radial_start + radial_length

		annulus := Annulus{x: cx, y: cy, start: arc_start, end: arc_end, inner: radial_start, outer: radial_end}

		path := annulus.path()
		fill := fmt.Sprintf("fill:%s", palette[rand.Intn(len(palette))])

		s.Path(path, fill)
	}

	s.End()
}
