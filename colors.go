package goannular

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Colors struct {
	Palettes [][]string `json:"palettes"`
}

func loadColorPalettes(filename string) (*Colors, error) {

	var colors Colors
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	byteValue, _ := ioutil.ReadAll(f)
	err = json.Unmarshal(byteValue, &colors)
	if err != nil {
		return nil, err
	}
	return &colors, nil
}
