package goannular

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
)

type Colors struct {
	palettes [][]string `json:"palettes"`
}

func (c Colors) Load(filename string) error {

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	byteValue, _ := ioutil.ReadAll(f)
	err = json.Unmarshal(byteValue, &c.palettes)
	if err != nil {
		return err
	}
	return nil
}

func (c Colors) RandomPalette() {
	c.palettes[rand.Intn(len(c.palettes))]
}
