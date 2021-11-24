package main

import (
	"fmt"
	"log"
	"os"

	goannular "github.com/engelsjk/go-annular"
)

func main() {

	f, err := os.Create("annular.svg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	annular := goannular.NewAnnular()

	annular.Draw()
	if err := annular.Render(f, "svg"); err != nil {
		log.Println(err.Error())
	}

	fmt.Println("annular.svg saved")
}
