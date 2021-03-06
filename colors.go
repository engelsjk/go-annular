package goannular

import (
	"errors"
	"fmt"
	"image/color"
	"math/rand"
)

type Colors struct {
	palette            []string
	numColorsInPalette int
	palettes           [][]string
	numPalettes        int
}

func (c *Colors) SetRandomPalette() {
	c.palette = c.palettes[rand.Intn(c.numPalettes)]
	c.numColorsInPalette = len(c.palette)
}

func (c *Colors) RandomColorHex() string {
	if c.palette == nil {
		c.SetRandomPalette()
	}
	return c.palette[rand.Intn(c.numColorsInPalette)]
}

func (c *Colors) RandomColorRGBA() color.RGBA {
	h := c.RandomColorHex()
	rgba, err := parseHexColorFast(h)
	if err != nil {
		fmt.Println(err.Error())
	}
	return rgba
}

func (c *Colors) RandomColorOrBlackHex() string {
	if c.palette == nil {
		c.SetRandomPalette()
	}
	palette := append(c.palette, "#000000")
	return palette[rand.Intn(c.numColorsInPalette+1)]
}

func (c *Colors) RandomColorOrBlackRGBA() color.RGBA {
	h := c.RandomColorOrBlackHex()
	rgba, err := parseHexColorFast(h)
	if err != nil {
		fmt.Println(err.Error())
	}
	return rgba
}

// https://stackoverflow.com/questions/54197913/parse-hex-string-to-image-color
func parseHexColorFast(s string) (c color.RGBA, err error) {

	var errInvalidFormat = errors.New("invalid format")

	c.A = 0xff

	if s[0] != '#' {
		return c, errInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = errInvalidFormat
	}
	return
}
