package wspr

import "testing"

func TestConvert(t *testing.T) {
	var flagtests = []struct {
		in int64
	}{
		{0},
		{1},
		{2},
		{3},
		{4},
		{5},
		{6},
		{7},
		{8},
		{9},
		{10},
		{11},
		{12},
		{13},
		{14},
		{15},
		{16},
		{17},
		{18},
		{19},
		{20},
	}
	for _, tt := range flagtests {
		Convert(tt.in)
	}

}
