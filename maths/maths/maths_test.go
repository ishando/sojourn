package maths

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Round(n float64, p int) float64 {
	pow := math.Pow10(p)
	n = math.Round(n * pow)
	return n/pow
}

func TestMinMax(t *testing.T) {
	require := require.New(t)

	testCases := map[string]struct{
		input interface{}
		qual  int
		min   interface{}
		max   interface{}
	}{
		"minmax 1 of 6": {
			input: []float64{8,3,5,2,9,1},
			qual:  1,
			min:   []float64{1},
			max:   []float64{9},
		},
		"minmax 3 of 6": {
			input: []float64{8,3,5,2,9,1},
			qual:  3,
			min:   []float64{1,2,3},
			max:   []float64{9,8,5},
		},
		"minmax 7 of 6": {
			input: []float64{8,3,5,2,9,1},
			qual:  7,
			min:   []float64{8,3,5,2,9,1},
			max:   []float64{8,3,5,2,9,1},
		},
		"minmax with repeats": {
			input: []float64{8,2,1,3,5,2,9,1,8,9},
			qual:  3,
			min:   []float64{1,1,2},
			max:   []float64{9,9,8},
		},
		"qual 0": {
			input: []float64{8,3,5,2,9,1},
			qual:  0,
			min:   []float64{1},
			max:   []float64{9},
		},
		"empty list": {
			input: []float64{},
			qual:  2,
			min:   []float64{},
			max:   []float64{},
		},
		"ints": {
			input: []int{2,4,5,1,6,9},
			qual:  2,
			min:   []int{1,2},
			max:   []int{9,6},
		},
		"int8s": {
			input: []int8{2,4,5,1,6,9},
			qual:  2,
			min:   []int8{1,2},
			max:   []int8{9,6},
		},
		"int16s": {
			input: []int16{2,4,5,1,6,9},
			qual:  2,
			min:   []int16{1,2},
			max:   []int16{9,6},
		},
		"int32s": {
			input: []int32{2,4,5,1,6,9},
			qual:  2,
			min:   []int32{1,2},
			max:   []int32{9,6},
		},
		"int64s": {
			input: []int64{2,4,5,1,6,9},
			qual:  2,
			min:   []int64{1,2},
			max:   []int64{9,6},
		},
		"float32s": {
			input: []float32{2,4,5,1,6,9},
			qual:  2,
			min:   []float32{1,2},
			max:   []float32{9,6},
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			switch inp := tc.input.(type) {
			case []int:
				min := Min(inp, tc.qual)
				require.Equal(tc.min.([]int), min, "should get the expected mins for int")
				max := Max(inp, tc.qual)
				require.Equal(tc.max.([]int), max, "should get the expected maxs for int")
			case []int8:
				min := Min(inp, tc.qual)
				require.Equal(tc.min.([]int8), min, "should get the expected mins for int8")
				max := Max(inp, tc.qual)
				require.Equal(tc.max.([]int8), max, "should get the expected maxs for int8")
			case []int16:
				min := Min(inp, tc.qual)
				require.Equal(tc.min.([]int16), min, "should get the expected mins for int16")
				max := Max(inp, tc.qual)
				require.Equal(tc.max.([]int16), max, "should get the expected maxs for int16")
			case []int32:
				min := Min(inp, tc.qual)
				require.Equal(tc.min.([]int32), min, "should get the expected mins for int32")
				max := Max(inp, tc.qual)
				require.Equal(tc.max.([]int32), max, "should get the expected maxs for int32")
			case []int64:
				min := Min(inp, tc.qual)
				require.Equal(tc.min.([]int64), min, "should get the expected mins for int64")
				max := Max(inp, tc.qual)
				require.Equal(tc.max.([]int64), max, "should get the expected maxs for int64")
			case []float32:
				min := Min(inp, tc.qual)
				require.Equal(tc.min.([]float32), min, "should get the expected mins for float32")
				max := Max(inp, tc.qual)
				require.Equal(tc.max.([]float32), max, "should get the expected maxs for float32")
			case []float64:
				min := Min(inp, tc.qual)
				require.Equal(tc.min.([]float64), min, "should get the expected mins for float64")
				max := Max(inp, tc.qual)
				require.Equal(tc.max.([]float64), max, "should get the expected maxs for float64")
			default:
				require.FailNow("Unhandled input type provided")
			}
		})
	}
}

func TestAvgMed(t *testing.T) {
	require := require.New(t)

	testCases := map[string]struct{
		input interface{}
		avg   float64
		med   float64
	}{
		"avgmed of 6": {
			input: []float64{8,3,7,2,9,1},
			avg:   float64(5),
			med:   float64(5),
		},
		"avgmed of 5": {
			input: []float64{8,3,6,2,7},
			avg:   float64(5.2),
			med:   float64(6),
		},
		"avgmed of long list": {
			input: []float64{8,2,3,6,2,7,4,6,2,2,5,6,7,3,32,4,65,7,5,34,2,34,45,65,7,76,57,456,45,4,1,6,6,5},
			avg:   float64(29.97),
			med:   float64(6),
		},
		"empty list": {
			input: []float64{},
			avg:   float64(0),
			med:   float64(0),
		},
		"ints": {
			input: []int{8,3,6,2,7},
			avg:   float64(5.2),
			med:   float64(6),
		},
		"int64s": {
			input: []int64{8,3,6,2,7},
			avg:   float64(5.2),
			med:   float64(6),
		},
		"float32s": {
			input: []float32{8,3,6,2,7},
			avg:   float64(5.2),
			med:   float64(6),
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			switch inp := tc.input.(type) {
			case []int:
				avg := Avg(inp)
				require.Equal(tc.avg, Round(avg,2), "should get the expected avg for ints")
				med := Median(inp)
				require.Equal(tc.med, med, "should get the expected median for ints")
			case []int64:
				avg := Avg(inp)
				require.Equal(tc.avg, Round(avg,2), "should get the expected avg for int64s")
				med := Median(inp)
				require.Equal(tc.med, med, "should get the expected median for int64s")
			case []float32:
				avg := Avg(inp)
				require.Equal(tc.avg, Round(avg,2), "should get the expected avg for float32s")
				med := Median(inp)
				require.Equal(tc.med, med, "should get the expected median for float32s")
			case []float64:
				avg := Avg(inp)
				require.Equal(tc.avg, Round(avg,2), "should get the expected avg for float64s")
				med := Median(inp)
				require.Equal(tc.med, med, "should get the expected median for float64s")
			default:
				require.FailNow("Unhandled input type provided")
			}
		})
	}
}

func TestPercentile(t *testing.T) {
	require := require.New(t)

	testCases := map[string]struct{
		input interface{}
		pct   int
		ans   interface{}
		loop  bool
	}{
		"50th percentile of 6": {
			input: []float64{8,3,7,2,9,1},
			pct:   50,
			ans:   float64(3),
		},
		"25th percentile of 10": {
			input: []float64{8,3,7,2,9,1,5,4,6,0},
			pct:   25,
			ans:   float64(1),
		},
		"75th percentile of 10": {
			input: []float64{8,3,7,2,9,1,5,4,6,0},
			pct:   75,
			ans:   float64(6),
		},
		"100th percentile of 10": {
			input: []float64{8,3,7,2,9,1,5,4,6,0},
			pct:   100,
			ans:   float64(9),
		},
		"1st percentile of 10": {
			input: []float64{8,3,7,2,9,1,5,4,6,0},
			pct:   1,
			ans:   float64(0),
		},
		"90th percentile of list with repeats": {
			input: []float64{8,3,7,2,9,1,5,4,6,0,9,8,7,9,8,7,9,8,7,2},
			pct:   90,
			ans:   float64(9),
		},
		"70th percentile of list with repeats": {
			input: []float64{8,3,7,2,9,1,5,4,6,0,9,8,7,9,8,7,9,8,7,2},
			pct:   70,
			ans:   float64(8),
		},
		"50th percentile of list with repeats": {
			input: []float64{8,3,7,2,9,1,5,4,6,0,9,8,7,9,8,7,9,8,7,2},
			pct:   50,
			ans:   float64(7),
		},
		"30th percentile of list with repeats": {
			input: []float64{8,3,7,2,9,1,5,4,6,0,9,8,7,9,8,7,9,8,7,2},
			pct:   30,
			ans:   float64(4),
		},
		"ith percentile of list of 100": {
			input: []float64{61, 60, 32, 89, 62, 79, 80, 8, 98, 1, 76, 47, 94, 34, 14, 90, 25, 57, 40, 24,
			                 86, 44, 53, 100, 58, 77, 2, 13, 11, 74, 88, 83, 3, 10, 71, 22, 21, 51, 82, 18,
			                 73, 99, 67, 55, 46, 17, 4, 38, 78, 92, 85, 28, 20, 84, 27, 52, 45, 72, 31, 69,
			                 97, 12, 56, 64, 65, 81, 68, 93, 59, 75, 35, 66, 96, 19, 29, 37, 5, 54, 16, 39,
			                 43, 63, 30, 26, 15, 33, 95, 49, 23, 50, 91, 36, 7, 70, 6, 48, 41, 42, 87, 9,
		},
			pct:   30,
			ans:   float64(30),
			loop:  true,
		},
		"ith percentile of list of ints": {
			input: []int{61, 60, 32, 89, 62, 79, 80, 8, 98, 1, 76, 47, 94, 34, 14, 90, 25, 57, 40, 24,
			                 86, 44, 53, 100, 58, 77, 2, 13, 11, 74, 88, 83, 3, 10, 71, 22, 21, 51, 82, 18,
			                 73, 99, 67, 55, 46, 17, 4, 38, 78, 92, 85, 28, 20, 84, 27, 52, 45, 72, 31, 69,
			                 97, 12, 56, 64, 65, 81, 68, 93, 59, 75, 35, 66, 96, 19, 29, 37, 5, 54, 16, 39,
			                 43, 63, 30, 26, 15, 33, 95, 49, 23, 50, 91, 36, 7, 70, 6, 48, 41, 42, 87, 9,
		},
			pct:   30,
			ans:   30,
			loop:  true,
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			switch inp := tc.input.(type) {
			case []int:
				ans := Percentile(inp, tc.pct)
				require.Equal(tc.ans.(int), ans, "should get the expected percentile for ints")
				if tc.loop {
					for i := 1; i <= 100; i++ {
						require.Equal(i, Percentile(inp, i), "should get the expected percentile for ints")
					}
				}
			case []float64:
				ans := Percentile(inp, tc.pct)
				require.Equal(tc.ans.(float64), ans, "should get the expected percentile for floats")
				if tc.loop {
					for i := 1; i <= 100; i++ {
						require.Equal(float64(i), Percentile(inp, i), "should get the expected percentile for floats")
					}
				}
			default:
				require.FailNow("Unhandled input type provided")
			}
		})
	}
}
