package main

import (
	"flag"
	"fmt"
	"gonum.org/v1/gonum/stat/distuv"
)

func getNormDistro(mu float64, sigma float64) {
	fmt.Println("f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,answer")
	e := distuv.Exponential{
		Rate: mu,
	}
	e2 := distuv.Exponential{
		Rate: mu * 0.5,
	}
	n := distuv.Normal{
		Mu:    mu,
		Sigma: sigma,
	}
	n2 := distuv.Normal{
		Mu:    mu * 0.5,
		Sigma: sigma + 3,
	}
	for i := 0; i < 2000; i++ {
		f1 := n.Rand()
		f2 := e.Rand()
		f3 := e2.Rand()
		f4 := n.Rand()
		f5 := n2.Rand()
		f6 := f1 + f2 - f4
		f7 := f1 + f3
		f8 := f3 + f4
		f9 := f4 * f5
		f10 := f1 + f2 + f3
		answer := f1 + f2 + f8 + n2.Rand()
		fmt.Printf("%.6f,%.6f,%.6f,%.6f,%.6f,%.6f,%.6f,%.6f,%.6f,%.6f,%.6f\n", f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, answer)
	}
}

func main() {
	m := flag.Float64("m", 1, "Mean")
	s := flag.Float64("s", 0.75, "Standard distribution")
	flag.Parse()
	getNormDistro(*m, *s)
}
