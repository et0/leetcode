package l_e_0168_string

import "testing"

type data struct {
	columnNumber int
	result       string
}

func TestWrapper(t *testing.T) {
	tests := []data{
		{1, "A"},
		{2147483647, "FXSHRXW"},
		{52, "AZ"},
		{701, "ZY"},
		{28, "AB"},
	}

	for _, v := range tests {
		result := Wrapper(v.columnNumber)
		if v.result != result {
			t.Error("Expected ", v.result, ", got ", result)
		}
	}
}
