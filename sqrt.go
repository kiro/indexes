package indexes

import("math")

type SqrtIndex struct {
	arr  []int 
	sqrt []int
	size int
	sqrtSize int	
	f func(a, b int) int
}

func NewSqrtIndex(size int, f func(a, b int) int) *SqrtIndex {	
	sqrtSize := int(math.Ceil(math.Sqrt(float64(size))))

	return &SqrtIndex{
		arr: make([]int, size),
		sqrt: make([]int, sqrtSize),
		size: size,
		sqrtSize: sqrtSize,
		f: f,
	}
}

func Min(a, b int) int {
	if a < b {
		return a 
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Sum(a, b int) int {
	return a + b
}

func (index *SqrtIndex) AddSqrt(at int, value int) {
	index.arr[at] = value

	sqrtIndex := at / index.sqrtSize

	start := sqrtIndex * index.sqrtSize 
	end := Min(start + index.sqrtSize, index.size)

	sqrtValue := index.arr[start]
	for i := start + 1; i < end; i++ {
		sqrtValue = index.f(sqrtValue, index.arr[i])
	}

	index.sqrt[sqrtIndex] = sqrtValue
}

func (index *SqrtIndex) Add(at int, value int) {	
	previousValue := index.arr[at]
	index.arr[at] = value
	index.sqrt[at / index.sqrtSize] += value - previousValue
}

func (index *SqrtIndex) Get(from int, to int) int {
	if from == to {
		return 0
	}

	result := index.arr[from]

	for i := from + 1; i < to; i++ {
		sqrtIndex := i / index.sqrtSize

		if sqrtIndex != from / index.sqrtSize && sqrtIndex != to / index.sqrtSize {
			result = index.f(result, index.sqrt[sqrtIndex])
			i += index.sqrtSize - 1;
		} else {
			result = index.f(result, index.arr[i])
		}	
	}

	return result
}

