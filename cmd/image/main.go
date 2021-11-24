package main

import (
	"fmt"
	"os"

	goannular "github.com/engelsjk/go-annular"
)

func main() {

	f, err := os.Create("annular.svg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	goannular.Run(f)

	fmt.Println("annular.svg saved")
}
