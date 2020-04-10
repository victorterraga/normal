package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
)

func getNormDistro(x int, mu float64, sigma float64) {
	for i := 0; i < x; i++ {
		fmt.Printf("%.2f\n", rand.NormFloat64()*float64(sigma)+float64(mu))
	}
}

func main() {
	n := flag.Int("n", 0, "Number of numbers")
	m := flag.Float64("m", 0, "Mean")
	s := flag.Float64("s", 0, "Standard distribution")
	flag.Parse()

	if *n == 0 {
		flag.Usage()
		os.Exit(1)
	}
	getNormDistro(*n, *m, *s)
}
