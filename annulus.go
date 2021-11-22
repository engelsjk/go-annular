package goannular

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

type Annulus struct {
	x, y, start, end, inner, outer float64
}

func (a Annulus) moveTo() string {

	x := a.x + a.outer*math.Cos(a.start)
	y := a.y + a.outer*math.Sin(a.start)

	mt := fmt.Sprintf(
		"M%s,%s ",
		strconv.FormatFloat(x, 'f', -1, 64),
		strconv.FormatFloat(y, 'f', -1, 64),
	)
	return (mt)
}

func (a Annulus) ellipticalArcCurve(inorout string) string {

	var angle, sweepflag int
	var rx, ry, x, y float64

	arc_length := a.end - a.start

	arc := math.Mod(arc_length, 2*math.Pi)

	large_arc := 0
	if arc > math.Pi {
		large_arc = 1
	}

	switch inorout {
	case "outer":
		angle = 0
		sweepflag = 1
		rx = a.outer
		ry = a.outer
		x = a.x + a.outer*math.Cos(a.end)
		y = a.y + a.outer*math.Sin(a.end)
	case "inner":
		angle = 0
		sweepflag = 0
		rx = a.inner
		ry = a.inner
		x = a.x + a.inner*math.Cos(a.start)
		y = a.y + a.inner*math.Sin(a.start)
	}

	eac := fmt.Sprintf(
		"A%s,%s,%s,%s,%s,%s,%s ",
		strconv.FormatFloat(rx, 'f', -1, 64),
		strconv.FormatFloat(ry, 'f', -1, 64),
		strconv.Itoa(angle),
		strconv.Itoa(large_arc),
		strconv.Itoa(sweepflag),
		strconv.FormatFloat(x, 'f', -1, 64),
		strconv.FormatFloat(y, 'f', -1, 64),
	)
	return (eac)
}

func (a Annulus) lineTo() string {

	x := a.x + a.inner*math.Cos(a.end)
	y := a.y + a.inner*math.Sin(a.end)

	lt := fmt.Sprintf(
		"L%s,%s ",
		strconv.FormatFloat(x, 'f', -1, 64),
		strconv.FormatFloat(y, 'f', -1, 64),
	)
	return (lt)
}

func (a Annulus) path() string {

	var d strings.Builder

	move_to := a.moveTo()
	eac_outer := a.ellipticalArcCurve("outer")
	line_to := a.lineTo()
	eac_inner := a.ellipticalArcCurve("inner")
	close_path := "z"

	d.WriteString(move_to)
	d.WriteString(eac_outer)
	d.WriteString(line_to)
	d.WriteString(eac_inner)
	d.WriteString(close_path)

	return (d.String())
}

func radialLength(radial_start, radial_length_base, arc_start float64, stype int) float64 {

	adj_radius_limits := []float64{250.0, 500.0}
	var radial_length float64

	switch stype {
	case 0:
		radial_length = rand.Float64() * radial_length_base //px
	case 1:
		switch {
		case radial_start < adj_radius_limits[0]:
			radial_length = rand.Float64() * radial_length_base / 3.0 //px
		case radial_start >= adj_radius_limits[0] && radial_start < adj_radius_limits[1]:
			radial_length = rand.Float64() * radial_length_base / 2.0 //px
		case radial_start >= adj_radius_limits[1]:
			radial_length = rand.Float64() * radial_length_base / 1.0 //px
		}
	case 2:
		radial_length = rand.Float64() * radial_length_base
	case 3:
		radial_length_base = rand.Float64() * radial_length_base
		radial_length = rand.Float64() * radial_length_base
	}

	return (radial_length)
}

func arcLength(arc_start, arc_length_base, radial_start float64, stype int) float64 {

	adj_radius_limits := []float64{250.0, 500.0}
	var arc_length float64

	switch stype {
	case 0:
		arc_length = rand.Float64() * 1.0 //px
	case 1:
		switch {
		case radial_start < adj_radius_limits[0]:
			arc_length = rand.Float64() * arc_length_base / 3.0 //px
		case radial_start >= adj_radius_limits[0] && radial_start < adj_radius_limits[1]:
			arc_length = rand.Float64() * arc_length_base / 2.0 //px
		case radial_start >= adj_radius_limits[1]:
			arc_length = rand.Float64() * arc_length_base / 1.0 //px
		}
	case 2:
		arc_length = rand.Float64() * arc_length_base
	case 3:
		arc_length_base = rand.Float64() * arc_length_base
		arc_length = rand.Float64() * arc_length_base
	}

	return (arc_length)
}
