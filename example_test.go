package stats_test

import (
	"fmt"
	"math/rand"

	"github.com/tebeka/stats"
)

func ExampleArgMin() {
	v := []int{3, 1, 4, 2}
	i, err := stats.ArgMin(v)
	fmt.Printf("argmin: %v (%v)\n", i, err)

	// Output:
	// argmin: 1 (<nil>)
}

func ExampleMin() {
	v := []int{3, 1, 4, 2}
	m, err := stats.Min(v)
	fmt.Printf("min: %v (%v)\n", m, err)

	// Output:
	// min: 1 (<nil>)
}

func ExampleArgMax() {
	v := []int{3, 1, 4, 2}
	i, err := stats.ArgMax(v)
	fmt.Printf("argmax: %v (%v)\n", i, err)

	// Output:
	// argmax: 2 (<nil>)
}

func ExampleMax() {
	v := []int{3, 1, 4, 2}
	m, err := stats.Max(v)
	fmt.Printf("max: %v (%v)\n", m, err)

	// Output:
	// max: 4 (<nil>)
}

func ExampleSum() {
	v := []int{3, 1, 4, 2}
	fmt.Println("sum:", stats.Sum(v))

	// Output:
	// sum: 10
}

func ExampleProd() {
	v := []int{3, 1, 4, 2}
	fmt.Println("prod:", stats.Prod(v))

	// Output:
	// prod: 24
}

func ExampleMean() {
	v := []int{3, 1, 4, 2}
	f, err := stats.Mean(v)
	fmt.Printf("mean: %v (%v)\n", f, err)

	// Output:
	// mean: 2.5 (<nil>)
}

func ExampleGeoMean() {
	v := []int{3, 1, 4, 2}
	f, err := stats.GeoMean(v)
	fmt.Printf("geo mean: %v (%v)\n", f, err)

	// Output:
	// geo mean: 2.213363839400643 (<nil>)
}

func ExampleHarmonicMean() {
	v := []int{3, 1, 4, 2}
	f, err := stats.HarmonicMean(v)
	fmt.Printf("harmonic mean: %v (%v)\n", f, err)

	// Output:
	// harmonic mean: 1.9200000000000004 (<nil>)
}

func ExampleMedian() {
	v := []int{3, 1, 2}
	f, err := stats.Median(v)
	fmt.Printf("median: %v (%v)\n", f, err)

	v = append(v, 4)
	f, err = stats.Median(v)
	fmt.Printf("median: %v (%v)\n", f, err)

	// Output:
	// median: 2 (<nil>)
	// median: 2.5 (<nil>)
}

func ExampleVar() {
	v := []int{3, 1, 4, 2}
	f, err := stats.Var(v)
	fmt.Printf("var: %v (%v)\n", f, err)

	// Output:
	// var: 1.25 (<nil>)
}

func ExampleStd() {
	v := []int{3, 1, 4, 2}
	f, err := stats.Std(v)
	fmt.Printf("std: %v (%v)\n", f, err)

	// Output:
	// std: 1.118033988749895 (<nil>)
}

func ExampleDot() {
	v := []int{3, 1, 4, 2}
	m, err := stats.Dot(v, v)
	fmt.Printf("dot : %v (%v)\n", m, err)

	// Output:
	// dot : 30 (<nil>)
}

func ExampleMode() {
	v := []rune{'h', 'e', 'l', 'l', 'o'}
	mod, err := stats.Mode(v)
	fmt.Printf("mode: %c (%v)\n", mod, err)

	// Output:
	// mode: l (<nil>)
}

func ExampleSample() {
	rand.Seed(353) // Deterministic result
	v := []string{"a", "b", "c", "d", "e", "f"}
	s, err := stats.Sample(v, 3)
	fmt.Printf("sample: %v (%v)\n", s, err)

	// Output:
	// sample: [f a c] (<nil>)
}

func ExampleShuffle() {
	rand.Seed(353) // Deterministic result
	v := []string{"a", "b", "c", "d", "e", "f"}
	stats.Shuffle(v)
	fmt.Println("shuffle:", v)

	// Output:
	// shuffle: [b a c d f e]
}
