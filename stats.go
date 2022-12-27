package vec

import (
	"errors"
	"math"
	"math/rand"
	"sort"
)

var (
	ErrEmpty = errors.New("empty slice")
	ErrSize  = errors.New("different size")
	ErrSmall = errors.New("to small")
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// ArgMin returns the index of minimal value in values
func ArgMin[T Ordered](values []T) (int, error) {
	if len(values) == 0 {
		return 0, ErrEmpty
	}

	am, m := 0, values[0]
	for i, v := range values[1:] {
		if v < m {
			am, m = i+1, v
		}
	}

	return am, nil
}

// Min returns the minimal value in values
func Min[T Ordered](values []T) (T, error) {
	i, err := ArgMin(values)
	if err != nil {
		var zero T
		return zero, err
	}

	return values[i], nil
}

// ArgMax returns the index of the maximal value in values
func ArgMax[T Ordered](values []T) (int, error) {
	if len(values) == 0 {
		return 0, ErrEmpty
	}

	am, m := 0, values[0]
	for i, v := range values[1:] {
		if v > m {
			am, m = i+1, v
		}
	}

	return am, nil

}

// Max returns the maximal value in values
func Max[T Ordered](values []T) (T, error) {
	i, err := ArgMax(values)
	if err != nil {
		var zero T
		return zero, err
	}

	return values[i], nil
}

// Sum return the sum of values
func Sum[T Ordered](values []T) T {
	var s T = 0
	for _, v := range values {
		s += v
	}
	return s
}

// Prod return the product of values
func Prod[T Ordered](values []T) T {
	var p T = 1
	for _, v := range values {
		p *= v
	}
	return p
}

// Mean return the arithmetic mean of values
func Mean[T Ordered](values []T) (float64, error) {
	if len(values) == 0 {
		return 0, ErrEmpty
	}

	s := Sum(values)
	m := float64(s) / float64(len(values))
	return m, nil
}

// GeoMean return the geometric mean of values
func GeoMean[T Ordered](values []T) (float64, error) {
	if len(values) == 0 {
		return 0, ErrEmpty
	}

	s := 0.0
	for _, v := range values {
		s += math.Log(float64(v))
	}
	s /= float64(len(values))
	return math.Exp(s), nil
}

// HarmonicMean return the harmonic mean of values
func HarmonicMean[T Ordered](values []T) (float64, error) {
	if len(values) == 0 {
		return 0, ErrEmpty
	}

	s := 0.0
	for _, v := range values {
		s += 1 / float64(v)
	}
	m := float64(len(values)) / s
	return m, nil
}

// Median return the median of values
func Median[T Ordered](values []T) (float64, error) {
	if len(values) == 0 {
		return 0, ErrEmpty
	}

	// Copy so we won't mutate values
	vs := make([]T, len(values))
	copy(vs, values)

	sort.Slice(vs, func(i, j int) bool {
		return vs[i] < vs[j]
	})

	i := len(vs) / 2
	if len(vs)%2 == 1 {
		return float64(vs[i]), nil
	}

	m := (float64(vs[i-1]) + float64(vs[i])) / 2
	return m, nil
}

// Var returns the variance of values
func Var[T Ordered](values []T) (float64, error) {
	if len(values) == 0 {
		return 0, ErrEmpty
	}

	m, err := Mean(values)
	if err != nil {
		return 0, err
	}

	td := 0.0
	for _, v := range values {
		d := m - float64(v)
		td += d * d
	}

	v := td / float64(len(values))
	return v, nil
}

// Std returns the standard deviation of values
func Std[T Ordered](values []T) (float64, error) {
	if len(values) == 0 {
		return 0, ErrEmpty
	}

	v, err := Var(values)
	if err != nil {
		return 0, err
	}

	return math.Sqrt(v), nil
}

// Dot returns the dot product of v1 and v2
func Dot[T1 Ordered, T2 Ordered](v1 []T1, v2 []T2) (T1, error) {
	if len(v1) != len(v2) {
		var zero T1
		return zero, ErrSize
	}

	var t T1 = 0
	for i, v := range v1 {
		t += v * T1(v1[i])
	}

	return t, nil
}

// Mode returns the most common element in values
func Mode[T comparable](values []T) (T, error) {
	if len(values) == 0 {
		var zero T
		return zero, ErrEmpty
	}

	freq := make(map[T]int)
	for _, v := range values {
		freq[v]++
	}

	var mode T
	count := 0
	for v, c := range freq {
		if c > count {
			mode, count = v, c
		}
	}

	return mode, nil
}

// Sample returns a sample of "k" elements from values
func Sample[T any](values []T, k int) ([]T, error) {
	if k > len(values) {
		return nil, ErrSmall
	}

	// TODO: Make it more efficient
	idx := rand.Perm(len(values))
	s := make([]T, 0, k)
	for _, i := range idx[:k] {
		s = append(s, values[i])
	}

	return s, nil
}

// Shuffle shuffles values in place
func Shuffle[T any](values []T) {
	rand.Shuffle(len(values), func(i, j int) {
		values[i], values[j] = values[j], values[i]
	})
}
