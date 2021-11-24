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

	arcLength := a.end - a.start

	arc := math.Mod(arcLength, 2*math.Pi)

	largeArc := 0
	if arc > math.Pi {
		largeArc = 1
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
		strconv.Itoa(largeArc),
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

func (a Annulus) SVG() string {

	var d strings.Builder

	d.WriteString(a.moveTo())
	d.WriteString(a.ellipticalArcCurve("outer"))
	d.WriteString(a.lineTo())
	d.WriteString(a.ellipticalArcCurve("inner"))
	d.WriteString("z")

	return d.String()
}

func radialLength(radialStart, maxRadialLength, arcStart float64, stype int) float64 {

	adjRadiusLimits := []float64{250.0, 500.0}
	var radialLength float64

	switch stype {
	case 0:
		radialLength = rand.Float64() * maxRadialLength //px
	case 1:
		switch {
		case radialStart < adjRadiusLimits[0]:
			radialLength = rand.Float64() * maxRadialLength / 3.0 //px
		case radialStart >= adjRadiusLimits[0] && radialStart < adjRadiusLimits[1]:
			radialLength = rand.Float64() * maxRadialLength / 2.0 //px
		case radialStart >= adjRadiusLimits[1]:
			radialLength = rand.Float64() * maxRadialLength / 1.0 //px
		}
	case 2:
		radialLength = rand.Float64() * maxRadialLength
	case 3:
		maxRadialLength = rand.Float64() * maxRadialLength
		radialLength = rand.Float64() * maxRadialLength
	}

	return radialLength
}

func arcLength(arcStart, maxArcLength, radialStart float64, stype int) float64 {

	adjRadiusLimits := []float64{250.0, 500.0}
	var arcLength float64

	switch stype {
	case 0:
		arcLength = rand.Float64() * 1.0 //px
	case 1:
		switch {
		case radialStart < adjRadiusLimits[0]:
			arcLength = rand.Float64() * maxArcLength / 3.0 //px
		case radialStart >= adjRadiusLimits[0] && radialStart < adjRadiusLimits[1]:
			arcLength = rand.Float64() * maxArcLength / 2.0 //px
		case radialStart >= adjRadiusLimits[1]:
			arcLength = rand.Float64() * maxArcLength / 1.0 //px
		}
	case 2:
		arcLength = rand.Float64() * maxArcLength
	case 3:
		maxArcLength = rand.Float64() * maxArcLength
		arcLength = rand.Float64() * maxArcLength
	}

	return arcLength
}
