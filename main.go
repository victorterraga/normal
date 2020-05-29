package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func getNormDistro(mu float64, sigma float64) {
	fmt.Println("f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,answer")
	for i := 0; i < 2000; i++ {
		f1 := rand.NormFloat64()*float64(sigma) + float64(mu)
		f2 := rand.NormFloat64()*float64(sigma) + float64(mu)
		f3 := rand.NormFloat64()*float64(sigma) + float64(mu)
		f4 := rand.NormFloat64()*float64(sigma) + float64(mu)
		f5 := rand.NormFloat64()*float64(sigma) + float64(mu)
		f6 := f1 + f2 - f4
		f7 := f1 + 3*f3
		f8 := f3 * f4
		f9 := f4 * 2 * f5
		f10 := f1 + f2 + f3 + f4
		answer := f1 + f2 + 2*f5 + f8 + rand.NormFloat64()*float64(sigma) + float64(mu)
		fmt.Printf("%.6f,%.6f,%.6f,%.6f,%.6f,%.6f,%.6f,%.6f,%.6f,%.6f,%.6f\n", f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, answer)
	}
}

func main() {
	m := flag.Float64("m", 1, "Mean")
	s := flag.Float64("s", 0.75, "Standard distribution")
	flag.Parse()
	getNormDistro(*m, *s)
}
