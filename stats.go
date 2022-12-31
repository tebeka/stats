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

// ArgMin returns the index of minimal value in vec
func ArgMin[T Ordered](vec []T) (int, error) {
	if len(vec) == 0 {
		return 0, ErrEmpty
	}

	am, m := 0, vec[0]
	for i, v := range vec[1:] {
		if v < m {
			am, m = i+1, v
		}
	}

	return am, nil
}

// Min returns the minimal value in vec
func Min[T Ordered](vec []T) (T, error) {
	i, err := ArgMin(vec)
	if err != nil {
		var zero T
		return zero, err
	}

	return vec[i], nil
}

// ArgMax returns the index of the maximal value in vec
func ArgMax[T Ordered](vec []T) (int, error) {
	if len(vec) == 0 {
		return 0, ErrEmpty
	}

	am, m := 0, vec[0]
	for i, v := range vec[1:] {
		if v > m {
			am, m = i+1, v
		}
	}

	return am, nil

}

// Max returns the maximal value in vec
func Max[T Ordered](vec []T) (T, error) {
	i, err := ArgMax(vec)
	if err != nil {
		var zero T
		return zero, err
	}

	return vec[i], nil
}

// Sum return the sum of vec
func Sum[T Ordered](vec []T) T {
	var s T = 0
	for _, v := range vec {
		s += v
	}
	return s
}

// Prod return the product of vec
func Prod[T Ordered](vec []T) T {
	var p T = 1
	for _, v := range vec {
		p *= v
	}
	return p
}

// Mean return the arithmetic mean of vec
func Mean[T Ordered](vec []T) (float64, error) {
	if len(vec) == 0 {
		return 0, ErrEmpty
	}

	s := Sum(vec)
	m := float64(s) / float64(len(vec))
	return m, nil
}

// GeoMean return the geometric mean of vec
func GeoMean[T Ordered](vec []T) (float64, error) {
	if len(vec) == 0 {
		return 0, ErrEmpty
	}

	s := 0.0
	for _, v := range vec {
		s += math.Log(float64(v))
	}
	s /= float64(len(vec))
	return math.Exp(s), nil
}

// HarmonicMean return the harmonic mean of vec
func HarmonicMean[T Ordered](vec []T) (float64, error) {
	if len(vec) == 0 {
		return 0, ErrEmpty
	}

	s := 0.0
	for _, v := range vec {
		s += 1 / float64(v)
	}
	m := float64(len(vec)) / s
	return m, nil
}

// Median return the median of vec
func Median[T Ordered](vec []T) (float64, error) {
	if len(vec) == 0 {
		return 0, ErrEmpty
	}

	// Copy so we won't mutate vec
	vs := make([]T, len(vec))
	copy(vs, vec)

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

// Var returns the variance of vec
func Var[T Ordered](vec []T) (float64, error) {
	if len(vec) == 0 {
		return 0, ErrEmpty
	}

	m, err := Mean(vec)
	if err != nil {
		return 0, err
	}

	td := 0.0
	for _, v := range vec {
		d := m - float64(v)
		td += d * d
	}

	v := td / float64(len(vec))
	return v, nil
}

// Std returns the standard deviation of vec
func Std[T Ordered](vec []T) (float64, error) {
	if len(vec) == 0 {
		return 0, ErrEmpty
	}

	v, err := Var(vec)
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
		t += v * T1(v2[i])
	}

	return t, nil
}

// Mode returns the most common element in vec
func Mode[T comparable](vec []T) (T, error) {
	if len(vec) == 0 {
		var zero T
		return zero, ErrEmpty
	}

	freq := make(map[T]int)
	for _, v := range vec {
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

// Sample returns a sample of "k" elements from vec
func Sample[T any](vec []T, k int) ([]T, error) {
	if k > len(vec) {
		return nil, ErrSmall
	}

	// TODO: Make it more efficient
	idx := rand.Perm(len(vec))
	s := make([]T, 0, k)
	for _, i := range idx[:k] {
		s = append(s, vec[i])
	}

	return s, nil
}

// Shuffle shuffles vec in place
func Shuffle[T any](vec []T) {
	rand.Shuffle(len(vec), func(i, j int) {
		vec[i], vec[j] = vec[j], vec[i]
	})
}

// Magnitude returns the magnitude (norm) of vec
func Magnitude[T Ordered](vec []T) float64 {
	var total T = 0
	for _, v := range vec {
		total += v * v
	}

	return math.Sqrt(float64(total))
}

// CosineSim returns the cosine similarity between v1 and v2 in degrees
func CosineSim[T Ordered](v1, v2 []T) (float64, error) {
	if len(v1) != len(v2) {
		return 0, ErrSize
	}

	d, err := Dot(v1, v2)
	if err != nil {
		return 0, err
	}

	m1, m2 := Magnitude(v1), Magnitude(v2)
	return float64(d) / (m1 * m2), nil
}
