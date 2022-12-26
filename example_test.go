package stats

import "fmt"

func ExampleSats() {
	v := []int{3, 1, 4, 2}
	m, err := Min(v)
	fmt.Printf("min : %v (%v)\n", m, err)
	m, err = Max(v)
	fmt.Printf("max : %v (%v)\n", m, err)
	fmt.Println("sum :", Sum(v))
	fmt.Println("prod:", Prod(v))
	f, err := Mean(v)
	fmt.Printf("mean: %v (%v)\n", f, err)
	f, err = Median(v)
	fmt.Printf("med : %v (%v)\n", f, err)
	f, err = Var(v)
	fmt.Printf("var : %v (%v)\n", f, err)
	f, err = Std(v)
	fmt.Printf("std : %v (%v)\n", f, err)
	m, err = Dot(v, v)
	fmt.Printf("dot : %v (%v)\n", m, err)
	i, err := ArgMin(v)
	fmt.Printf("amin: %v (%v)\n", i, err)
	i, err = ArgMax(v)
	fmt.Printf("amax: %v (%v)\n", i, err)
	mod, err := Mode([]rune{'h', 'e', 'l', 'l', 'o'})
	fmt.Printf("mode: %c (%v)\n", mod, err)

	// Output:
	// min : 1 (<nil>)
	// max : 4 (<nil>)
	// sum : 10
	// prod: 24
	// mean: 2.5 (<nil>)
	// med : 2.5 (<nil>)
	// var : 1.25 (<nil>)
	// std : 1.118033988749895 (<nil>)
	// dot : 30 (<nil>)
	// amin: 1 (<nil>)
	// amax: 2 (<nil>)
	// mode: l (<nil>)
}
