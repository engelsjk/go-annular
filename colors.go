package goannular

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Colors struct {
	Palettes [][]string `json:"palettes"`
}

func loadColorPalettes(filename string) Colors {

	var colors Colors
	jsonFile, err := os.Open(filename)
	if verbose {
		if err != nil {
			fmt.Println(err)
		}
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &colors)
	if verbose {
		if err != nil {
			fmt.Println(err)
		}
	}
	return (colors)
}
