package datastore

import (
	"log"
	"strconv"
)

func utilStrToInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func utilStrToFloat(s string) float64 {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return val
}
