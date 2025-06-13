package fsql

import "testing"

func Test_token_is(t *testing.T) {
	tests := []struct {
		tok      token
		ids      []uint8
		expected bool
	}{
		{newBase(add, pos{}), []uint8{add}, true},
		{newBase(add, pos{}), []uint8{sub}, false},
		{newBase(add, pos{}), []uint8{sub, add}, true},
		{newBase(add, pos{}), []uint8{mul, div, rem}, false},
		{newBase(add, pos{}), []uint8{}, false},
		{newVal(integer, pos{}, ""), []uint8{add}, false},
		{newVal(float, pos{}, ""), []uint8{sub}, false},
		{newVal(text, pos{}, ""), []uint8{sub, add}, false},
		{newVal(boolean, pos{}, ""), []uint8{mul, div, rem}, false},
		{newVal(integer, pos{}, ""), []uint8{integer, float}, true},
	}

	for ti, test := range tests {
		if test.expected != test.tok.is(test.ids...) {
			t.Errorf("%d: ожидал %t, получил %t",
				ti+1, test.expected, !test.expected)
		}
	}
}
