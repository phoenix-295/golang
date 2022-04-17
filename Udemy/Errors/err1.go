package main

import (
	"errors"
	"log"
)

func main() {
	_, err := sqr(-10.2)
	if err != nil {
		log.Fatalln(err)
	}

}

func sqr(f float64) (float64, error) {
	if f < 0 {
		// create manual error
		return 0, errors.New("Negative Int")
	}
	return 42, nil
}
