package goannular

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
)

type Colors struct {
	palette            []string
	numColorsInPalette int
	palettes           [][]string
	numPalettes        int
}

func (c *Colors) Load(filename string) error {

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	var palettes [][]string

	byteValue, err := ioutil.ReadAll(f)
	err = json.Unmarshal(byteValue, &palettes)
	if err != nil {
		return err
	}

	c.palettes = palettes
	c.numPalettes = len(c.palettes)
	return nil
}

func (c *Colors) SetRandomPalette() {
	c.palette = c.palettes[rand.Intn(c.numPalettes)]
	c.numColorsInPalette = len(c.palette)
}

func (c *Colors) RandomColor() string {
	if c.palette == nil {
		c.SetRandomPalette()
	}
	return c.palette[rand.Intn(c.numColorsInPalette)]
}

func (c *Colors) RandomColorOrBlack() string {
	if c.palette == nil {
		c.SetRandomPalette()
	}
	palette := append(c.palette, "#000000")
	return palette[rand.Intn(c.numColorsInPalette+1)]
}
