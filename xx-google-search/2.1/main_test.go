package main

import (
	"math/rand"
	"time"
	"testing"
)

func TestGoogle(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()

	results := Google("golang")
	
	elapsed := time.Since(start)
	if len(results) != 3 {
		t.Fail()
	}
	println("elapsed:", elapsed)
}
