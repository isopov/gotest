package main

import (
	"image/jpeg"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("/home/isopov/tmp/original.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}

	for _, q := range []int{75, 90, 95, 99, 100} {
		w, err := os.Create("/home/isopov/tmp/encoded" + strconv.Itoa(q) + ".jpg")
		if err != nil {
			panic(err)
		}
		err = jpeg.Encode(w, m, &jpeg.Options{Quality: q})
		if err != nil {
			panic(err)
		}
		err = w.Close()
		if err != nil {
			panic(err)
		}
	}
}
